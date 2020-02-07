package ast

import (
	"errors"
)

const (
	NsData   = 0
	NsElem   = 1
	NsFunc   = 2
	NsGlobal = 3
	NsMemory = 4
	NsTable  = 5
	NsType   = 6
)

type NameResolver struct {
	ns    [7]Namespace
	types []FunctionType
}

func NewNameResolver() NameResolver {
	nr := NameResolver{}
	for i := 0; i < 7; i++ {
		nr.ns[i] = NewNamespace()
	}

	return nr
}

type Namespace struct {
	names map[Id]uint32
	count uint32
}

func (self *NameResolver) Register(item ModuleField) {
	switch value := item.(type) {
	case Import:
		switch value.Item.(type) {
		case ImportFunc:
			self.ns[NsFunc].register(value.Id)
		case ImportMemory:
			self.ns[NsMemory].register(value.Id)
		case ImportTable:
			self.ns[NsTable].register(value.Id)
		case ImportGlobal:
			self.ns[NsGlobal].register(value.Id)
		default:
			panic("unreachable")
		}
	case Global:
		self.ns[NsGlobal].register(value.Name)
	case Memory:
		self.ns[NsMemory].register(value.Name)
	case Func:
		self.ns[NsFunc].register(value.Name)
	case Table:
		self.ns[NsTable].register(value.Name)
	case Type:
		//todo
		self.ns[NsType].register(value.Name)
	case Elem:
		self.ns[NsElem].register(value.Name)
	case Data:
		self.ns[NsData].register(value.Name)
	case Export, StartField:

	default:
		panic("unreachable")
	}
}

func (self *NameResolver) Resolve(field ModuleField) (ModuleField, error) {
	switch value := field.(type) {
	case Import:
		if fn, ok := value.Item.(ImportFunc); ok {
			_, err := self.resolveTypeUse(&fn.TypeUse)
			value.Item = fn
			return value, err
		}
		return value, nil
	case Func:
		_, err := self.resolveTypeUse(&value.Type)
		if err != nil {
			return value, err
		}
		if inline, ok := value.Kind.(FuncKindInline); ok {
			exprResolver := NewExprResolver(self)
			for _, param := range value.Type.Type.Params {
				exprResolver.locals.register(param.Id)
			}
			for _, local := range inline.Locals {
				exprResolver.locals.register(local.Id)
			}

			err := exprResolver.Resolve(&inline.Expr)
			if err != nil {
				return nil, err
			}
			value.Kind = inline
		}

		return value, nil
	case Elem:
		if active, ok := value.Kind.(ElemKindActive); ok {
			err := self.resolveIdx(&active.Table, NsTable)
			if err != nil {
				return nil, err
			}
			err = self.resolveExpr(&active.Offset)
			if err != nil {
				return nil, err
			}

			value.Kind = active
		}
		switch payload := value.Payload.(type) {
		case ElemPayloadIndices:
			for i := 0; i < len(payload.Indices); i++ {
				err := self.resolveIdx(&payload.Indices[i], NsFunc)
				if err != nil {
					return nil, err
				}
			}
			value.Payload = payload
		case ElemPayloadExprs:
			for i := 0; i < len(payload.Exprs); i++ {
				if payload.Exprs[i].IsSome() {
					id := payload.Exprs[i].ToIndex()
					err := self.resolveIdx(&id, NsFunc)
					if err != nil {
						return nil, err
					}
					payload.Exprs[i] = NewOptionIndex(id)
				}
			}
			value.Payload = payload
		}

		return value, nil
	case Data:
		if active, ok := value.Kind.(DataKindActive); ok {
			err := self.resolveIdx(&active.Memory, NsMemory)
			if err != nil {
				return nil, err
			}
			err = self.resolveExpr(&active.Offset)
			if err != nil {
				return nil, err
			}
			value.Kind = active
		}

		return value, nil
	case StartField:
		err := self.resolveIdx(&value.Index, NsFunc)
		return value, err
	case Export:
		var err error
		var ns int
		switch value.Type {
		case ExportFunc:
			ns = NsFunc
		case ExportMemory:
			ns = NsMemory
		case ExportGlobal:
			ns = NsGlobal
		case ExportTable:
			ns = NsTable
		}
		err = self.resolveIdx(&value.Index, ns)
		return value, err
	case Global:
		if inline, ok := value.Kind.(GlobalKindInline); ok {
			err := self.resolveExpr(&inline.Expr)
			if err != nil {
				return nil, err
			}
			value.Kind = inline
		}

		return value, nil
	case Table, Memory, Type:
		return value, nil
	default:
		panic("unreachable")
	}
}

func (self *NameResolver) resolveIdx(idx *Index, ns int) error {
	_, err := self.ns[ns].resolve(idx)
	return err
}

func (self *NameResolver) resolveExpr(expr *Expression) error {
	exprResolver := NewExprResolver(self)
	return exprResolver.Resolve(expr)
}

func (self *NameResolver) resolveTypeUse(ty *TypeUse) (uint32, error) {
	if !ty.Index.IsSome() {
		panic("must be some index")
	}
	idx := ty.Index.ToIndex()
	index, err := self.ns[NsType].resolve(&idx)
	if err != nil {
		return 0, err
	}
	ty.Index = NewOptionIndex(idx)
	if index >= uint32(len(self.types)) {
		return index, nil
	}
	expected := self.types[index]
	if len(ty.Type.Params) > 0 || len(ty.Type.Results) > 0 {
		//notEqual := len(expected.Params) != len(ty.Type.Params)
		//if notEqual
		//todo: check equal
	} else {
		ty.Type.Params = expected.Params
		ty.Type.Results = expected.Results
	}

	return index, nil
}

func NewNamespace() Namespace {
	return Namespace{
		names: make(map[Id]uint32),
	}
}
func (self *Namespace) register(name OptionId) {
	if name.IsSome() {
		self.names[name.ToId()] = self.count
	}
	self.count += 1
}

func (self *Namespace) resolve(idx *Index) (uint32, error) {
	if idx.Isnum {
		return idx.Num, nil
	}

	if n, ok := self.names[idx.Id]; ok {
		*idx = NewNumIndex(n)
		return n, nil
	}
	panic(errors.New("namespace can not resolve index"))

	return 0, errors.New("namespace can not resolve index")
}

type ExprResolver struct {
	resolver *NameResolver
	locals   Namespace
	labels   []OptionId
}

func NewExprResolver(resolver *NameResolver) ExprResolver {
	return ExprResolver{
		resolver: resolver,
		locals:   NewNamespace(),
	}
}

func (self *ExprResolver) Resolve(expr *Expression) error {
	for i := 0; i < len(expr.Instrs); i++ {
		err := self.resolveInstr(expr.Instrs[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *ExprResolver) resolveInstr(instr Instruction) error {
	handleBlockType := func(blockType *BlockType) error {
		self.labels = append(self.labels, blockType.Label)
		if blockType.Ty.Index.IsSome() {
			ind, err := self.resolver.resolveTypeUse(&blockType.Ty)
			if err != nil {
				return err
			}
			if ind >= uint32(len(self.resolver.types)) {
				return nil
			}
			ty := self.resolver.types[ind]
			if len(ty.Params) == 0 && len(ty.Results) <= 1 {
				blockType.Ty.Type.Params = nil
				blockType.Ty.Type.Results = ty.Results
				blockType.Ty.Index = NoneOptionIndex()
			}
		}
		return nil
	}

	switch inst := instr.(type) {
	//todo add TableInit, MemoryInit
	case *DataDrop:
		return self.resolver.resolveIdx(&inst.Index, NsData)
	case *ElemDrop:
		return self.resolver.resolveIdx(&inst.Index, NsElem)
	case *TableFill:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *TableSet:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *TableGet:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *TableSize:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *TableGrow:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *GlobalSet:
		return self.resolver.resolveIdx(&inst.Index, NsGlobal)
	case *GlobalGet:
		return self.resolver.resolveIdx(&inst.Index, NsGlobal)
	case *LocalSet:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *LocalGet:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *LocalTee:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *Call:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *RefFunc:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *ReturnCall:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *CallIndirect:
		err := self.resolver.resolveIdx(&inst.Impl.Table, NsTable)
		if err != nil {
			return err
		}
		_, err = self.resolver.resolveTypeUse(&inst.Impl.Type)
		return err
	case *ReturnCallIndirect:
		err := self.resolver.resolveIdx(&inst.Impl.Table, NsTable)
		if err != nil {
			return err
		}
		_, err = self.resolver.resolveTypeUse(&inst.Impl.Type)
		return err
	case *Block:
		return handleBlockType(&inst.BlockType)
	case *If:
		return handleBlockType(&inst.BlockType)
	case *Loop:
		return handleBlockType(&inst.BlockType)
	case *Else:
		if len(self.labels) == 0 {
			return nil
		}
		matching := self.labels[len(self.labels)-1]
		label := &inst.Id

		if !label.IsSome() || *label == matching {
			return nil
		}

		return errors.New("mismatching labels between block and end")
	case *End:
		if len(self.labels) == 0 {
			return nil
		}
		matching := self.labels[len(self.labels)-1]
		self.labels = self.labels[:len(self.labels)-1]
		label := &inst.Id

		if !label.IsSome() || *label == matching {
			return nil
		}

		return errors.New("mismatching labels between block and end")
	case *Br:
		return self.resolveLabel(&inst.Index)
	case *BrIf:
		return self.resolveLabel(&inst.Index)
	case *BrTable:
		for i := 0; i < len(inst.Indices.Labels); i++ {
			err := self.resolveLabel(&inst.Indices.Labels[i])
			if err != nil {
				return err
			}
		}

		return self.resolveLabel(&inst.Indices.Default)
	default:
		return nil
	}
}

func (self *ExprResolver) resolveLabel(label *Index) error {
	if label.Isnum {
		return nil
	}
	id := label.Id
	for i := len(self.labels) - 1; i >= 0; i-- {
		if self.labels[i].IsSome() && self.labels[i].ToId() == id {
			*label = NewNumIndex(uint32(len(self.labels) - i - 1))
			return nil
		}
	}

	return errors.New("failed to resolve label")
}
