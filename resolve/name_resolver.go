package resolve

import (
	"errors"
	"github.com/ontio/wast-parser/ast"
)

const (
	NsData = 0
	NsElem = 1
	NsFunc = 2
	NsGlobal = 3
	NsMemory = 4
	NsTable = 5
	NsType = 6
)

type NameResolver struct {
 	ns [7]Namespace
 	types []ast.FunctionType
}

type Namespace struct {
	names map[ast.Id] uint32
	count uint32
}

func (self *NameResolver) Register(item ast.ModuleField) {
	switch value := item.(type) {
	case ast.Import:
		switch value.Item.(type) {
		case ast.ImportFunc:
			self.ns[NsFunc].register(value.Id)
		case ast.ImportMemory:
			self.ns[NsMemory].register(value.Id)
		case ast.ImportTable:
			self.ns[NsTable].register(value.Id)
		case ast.ImportGlobal:
			self.ns[NsGlobal].register(value.Id)
		default:
			panic("unreachable")
		}
	case ast.Global:
		self.ns[NsGlobal].register(value.Name)
	case ast.Memory:
		self.ns[NsMemory].register(value.Name)
	case ast.Func:
		self.ns[NsFunc].register(value.Name)
	case ast.Table:
		self.ns[NsTable].register(value.Name)
	case ast.Type:
		//todo
		self.ns[NsType].register(value.Name)
	case ast.Elem:
		self.ns[NsElem].register(value.Name)
	case ast.Data:
		self.ns[NsData].register(value.Name)
	case ast.Export, ast.StartField:

	default:
		panic("unreachable")
	}
}

func (self *NameResolver) Resolve(field ast.ModuleField) (ast.ModuleField, error) {
	switch value := field.(type) {
	case ast.Import:
		if fn, ok := value.Item.(ast.ImportFunc); ok {
			_, err := self.resolveTypeUse(&fn.TypeUse)
			return value, err
		}
		return value, nil
	case ast.Func:
		_, err := self.resolveTypeUse(&value.Type)
		if err != nil {
			return value, err
		}
		if inline, ok := value.Kind.(ast.FuncKindInline); ok {
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
	case ast.Elem:
		if active, ok := value.Kind.(ast.ElemKindActive); ok {
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
		case ast.ElemPayloadIndices:
			for i:=0; i< len(payload.Indices) ; i++ {
				err := self.resolveIdx(&payload.Indices[i], NsFunc)
				if err != nil {
					return nil, err
				}
			}
			value.Payload = payload
		case ast.ElemPayloadExprs:
			for i:=0; i< len(payload.Exprs); i++ {
				if payload.Exprs[i].IsSome() {
					id := payload.Exprs[i].ToIndex()
					err := self.resolveIdx(&id, NsFunc)
					if err != nil {
						return nil, err
					}
					payload.Exprs[i] = ast.NewOptionIndex(id)
				}
			}
			value.Payload = payload
		}

		return value, nil
	case ast.Data:
		if active, ok := value.Kind.(ast.DataKindActive); ok {
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
	case ast.StartField:
		err := self.resolveIdx(&value.Index, NsFunc)
		return value, err
	case ast.Export:
		var err error
		var ns int
		switch value.Type {
		case ast.ExportFunc:
			ns = NsFunc
		case ast.ExportMemory:
			ns = NsMemory
		case ast.ExportGlobal:
			ns = NsGlobal
		case ast.ExportTable:
			ns = NsTable
		}
		err = self.resolveIdx(&value.Index, ns)
		return value, err
	case ast.Global:
		if inline, ok := value.Kind.(ast.GlobalKindInline); ok {
			err := self.resolveExpr(&inline.Expr)
			if err != nil {
				return nil, err
			}
			value.Kind = inline
		}

		return value, nil
	case ast.Table, ast.Memory, ast.Type:
		return value, nil
	default:
		panic("unreachable")
	}
}

func (self *NameResolver)resolveIdx(idx *ast.Index, ns int) error {
	_, err := self.ns[ns].resolve(idx)
	return err
}

func (self *NameResolver)resolveExpr(expr *ast.Expression) error {
	exprResolver := NewExprResolver(self)
	return exprResolver.Resolve(expr)
}


func (self *NameResolver)resolveTypeUse(ty *ast.TypeUse) (uint32, error) {
	if !ty.Index.IsSome() {
		panic("must be some index")
	}
	idx := ty.Index.ToIndex()
	index, err := self.ns[NsType].resolve(&idx)
	if err != nil {
		return 0, err
	}
	ty.Index = ast.NewOptionIndex(idx)
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

func (self *Namespace) register(name ast.OptionId) {
	if name.IsSome() {
		self.names[name.ToId()] = self.count
	}
	self.count += 1
}

func (self *Namespace)resolve(idx *ast.Index) (uint32, error) {
	if idx.Isnum {
		return idx.Num, nil
	}

	if n, ok := self.names[idx.Id]; ok {
		*idx = ast.NewNumIndex(n)
		return n, nil
	}

	return 0, errors.New("namespace can not resolve index")
}

type ExprResolver struct {
	resolver *NameResolver
	locals Namespace
	labels []ast.OptionId
}

func NewExprResolver(resolver *NameResolver) ExprResolver {
	return ExprResolver{
		resolver:resolver,
		locals:Namespace{},
	}
}

func (self *ExprResolver)Resolve(expr *ast.Expression) error {
	for i:=0; i< len(expr.Instrs); i++ {
		err := self.resolveInstr(expr.Instrs[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *ExprResolver)resolveInstr(instr ast.Instruction) error {
	handleBlockType := func (blockType *ast.BlockType) error {
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
			if len(ty.Params) == 0 && len(ty.Results) <= 1{
				blockType.Ty.Type.Params = nil
				blockType.Ty.Type.Results = ty.Results
				blockType.Ty.Index = ast.NoneOptionIndex()
			}
		}
		return nil
	}

	switch inst := instr.(type) {
	//todo add TableInit, MemoryInit
	case *ast.DataDrop:
		return self.resolver.resolveIdx(&inst.Index, NsData)
	case *ast.ElemDrop:
		return self.resolver.resolveIdx(&inst.Index, NsElem)
	case *ast.TableFill:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *ast.TableSet:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *ast.TableGet:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *ast.TableSize:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *ast.TableGrow:
		return self.resolver.resolveIdx(&inst.Index, NsTable)
	case *ast.GlobalSet:
		return self.resolver.resolveIdx(&inst.Index, NsGlobal)
	case *ast.GlobalGet:
		return self.resolver.resolveIdx(&inst.Index, NsGlobal)
	case *ast.LocalSet:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *ast.LocalGet:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *ast.LocalTee:
		_, err := self.locals.resolve(&inst.Index)
		return err
	case *ast.Call:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *ast.RefFunc:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *ast.ReturnCall:
		return self.resolver.resolveIdx(&inst.Index, NsFunc)
	case *ast.CallIndirect:
		err :=  self.resolver.resolveIdx(&inst.Impl.Table, NsTable)
		if err != nil {
			return err
		}
		_, err =  self.resolver.resolveTypeUse(&inst.Impl.Type)
		return err
	case *ast.ReturnCallIndirect:
		err :=  self.resolver.resolveIdx(&inst.Impl.Table, NsTable)
		if err != nil {
			return err
		}
		_, err =  self.resolver.resolveTypeUse(&inst.Impl.Type)
		return err
	case *ast.Block:
		return handleBlockType(&inst.BlockType)
	case *ast.If:
		return handleBlockType(&inst.BlockType)
	case *ast.Loop:
		return handleBlockType(&inst.BlockType)
	case *ast.Else:
		if len(self.labels) == 0 {
			return nil
		}
		matching := self.labels[len(self.labels) - 1]
		label := &inst.Id

		if !label.IsSome() || *label == matching {
			return nil
		}

		return errors.New("mismatching labels between block and end")
	case *ast.End:
		if len(self.labels) == 0 {
			return nil
		}
		matching := self.labels[len(self.labels) - 1]
		self.labels = self.labels[:len(self.labels) - 1]
		label := &inst.Id

		if !label.IsSome() || *label == matching {
			return nil
		}

		return errors.New("mismatching labels between block and end")
	case *ast.Br:
		return self.resolveLabel(&inst.Index)
	case *ast.BrIf:
		return self.resolveLabel(&inst.Index)
	case *ast.BrTable:
		for i:=0; i < len(inst.Indices.Labels); i++ {
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

func (self *ExprResolver)resolveLabel(label *ast.Index) error {
	if label.Isnum {
		return nil
	}
	id := label.Id
	for i := len(self.labels)-1; i >=0; i-- {
		if self.labels[i].IsSome() && self.labels[i].ToId() == id {
			*label = ast.NewNumIndex(uint32(len(self.labels) - i - 1))
			return nil
		}
	}

	return errors.New("failed to resolve label")
}








