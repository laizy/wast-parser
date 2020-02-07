package ast

type TypeExpander struct {
	toPrepend []ModuleField
	types     map[[2]string] uint32
	ntypes    uint32
}

func NewTypeExpander() TypeExpander {
	return TypeExpander{
		types:make(map[[2]string]uint32),
	}
}

func (self *TypeExpander)Expand(item ModuleField) ModuleField {
	switch val := item.(type) {
	case Type:
		self.registerType(&val)
		return val
	case Import:
		self.expandImport(&val)
		return val
	case Func:
		self.expandFunc(&val)
		return val
	case Global:
		self.expandGlobal(&val)
		return val
	case Data:
		self.expandData(&val)
		return val
	case Elem:
		self.expandElem(&val)
		return val
	default:
		return val
	}
}

func (self *TypeExpander)expandGlobal(global *Global) {
	if inline, ok := global.Kind.(GlobalKindInline); ok {
		self.expandExpression(&inline.Expr)
		global.Kind = inline
	}
}

func (self *TypeExpander)expandElem(elem *Elem) {
	if inline, ok := elem.Kind.(ElemKindActive); ok {
		self.expandExpression(&inline.Offset)
		elem.Kind = inline
	}
}

func (self *TypeExpander)expandData(data *Data) {
	if inline, ok := data.Kind.(DataKindActive); ok {
		self.expandExpression(&inline.Offset)
		data.Kind = inline
	}
}

func (self *TypeExpander)expandFunc(fn *Func) {
	self.expandTypeUse(&fn.Type)
	if inline, ok :=  fn.Kind.(FuncKindInline); ok {
		self.expandExpression(&inline.Expr)
		fn.Kind = inline
	}
}

func (self *TypeExpander)expandExpression(expr *Expression) {
	for i:= 0; i < len(expr.Instrs); i++ {
		self.expandInstr(expr.Instrs[i])
	}
}

func (self *TypeExpander)expandInstr(instr Instruction) {
	var blockType *BlockType
	switch ins := instr.(type) {
	case *Block:
		blockType = &ins.BlockType
	case *If:
		blockType = &ins.BlockType
	case *Loop:
		blockType = &ins.BlockType
	case *CallIndirect:
		self.expandTypeUse(&ins.Impl.Type)
	case *ReturnCallIndirect:
		self.expandTypeUse(&ins.Impl.Type)
	default:
	}

	if blockType != nil {
		// Only actually expand `TypeUse` with an index which appends a
		// type if it looks like we need one. This way if the
		// multi-value proposal isn't enabled and/or used we won't
		// encode it.
		if len(blockType.Ty.Type.Params) == 0 && len(blockType.Ty.Type.Results) <= 1 {
			return
		}

		self.expandTypeUse(&blockType.Ty)
	}

}

func TypeToKey(ty FunctionType) [2]string {
	var params []byte
	var results []byte
	for _, v := range ty.Params {
		params = append(params, v.Val.Byte())
	}
	for _, v := range ty.Results {
		results = append(results, v.Byte())
	}

	return [2]string{string(params), string(results)}
}

func (self *TypeExpander)registerType(val *Type) {
	key := TypeToKey(val.Func)
	if _, ok := self.types[key]; ok == false {
		self.types[key] = self.ntypes
	}

	self.ntypes += 1
}

func (self *TypeExpander) expandImport(val *Import) {
	if fn, ok := val.Item.(ImportFunc); ok {
		self.expandTypeUse(&fn.TypeUse)
		val.Item = fn
	}
}

func (self *TypeExpander) expandTypeUse(val *TypeUse) {
	if val.Index.IsSome() {
		return
	}
	key := TypeToKey(val.Type)
	if v, ok := self.types[key]; ok {
		val.Index = NewOptionIndex(NewNumIndex(v))
	} else {
		val.Index = NewOptionIndex(NewNumIndex( self.prepend(key) ))
	}
}

func (self *TypeExpander)prepend(key [2]string) uint32 {
	params , results := decodeKey(key)
	var funcParams []FuncParam
	for _, p := range params {
		funcParams = append(funcParams, FuncParam{
			Id  : NoneOptionId(),
			Val : p,
		})
	}

	self.toPrepend = append(self.toPrepend, Type{
		Name: NoneOptionId(),
		Func: FunctionType{
			Params:funcParams,
			Results:results,
		},
	})
	self.types[key] = self.ntypes
	self.ntypes += 1

	return self.ntypes - 1
}

func decodeKey(key [2]string) (params []ValType,  results []ValType) {
	for _, by := range []byte(key[0]) {
		params = append(params, ValTypeFromByte(by))
	}
	for _, by := range []byte(key[1]) {
		results = append(results, ValTypeFromByte(by))
	}

	return
}

