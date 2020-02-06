package resolve

import "github.com/ontio/wast-parser/ast"

type TypeExpander struct {
	toPrepend []ast.ModuleField
	types     map[[2]string] uint32
	ntypes    uint32
}

func (self *TypeExpander)Expand(item ast.ModuleField) ast.ModuleField {
	switch val := item.(type) {
	case ast.Type:
		self.registerType(&val)
		return val
	case ast.Import:
		self.expandImport(&val)
		return val
	case ast.Func:
		self.expandFunc(&val)
		return val
	case ast.Global:
		self.expandGlobal(&val)
		return val
	case ast.Data:
		self.expandData(&val)
		return val
	case ast.Elem:
		self.expandElem(&val)
		return val
	default:
		return val
	}
}

func (self *TypeExpander)expandGlobal(global *ast.Global) {
	if inline, ok := global.Kind.(ast.GlobalKindInline); ok {
		self.expandExpression(&inline.Expr)
		global.Kind = inline
	}
}

func (self *TypeExpander)expandElem(elem *ast.Elem) {
	if inline, ok := elem.Kind.(ast.ElemKindActive); ok {
		self.expandExpression(&inline.Offset)
		elem.Kind = inline
	}
}

func (self *TypeExpander)expandData(data *ast.Data) {
	if inline, ok := data.Kind.(ast.DataKindActive); ok {
		self.expandExpression(&inline.Offset)
		data.Kind = inline
	}
}

func (self *TypeExpander)expandFunc(fn *ast.Func) {
	self.expandTypeUse(&fn.Type)
	if inline, ok :=  fn.Kind.(ast.FuncKindInline); ok {
		self.expandExpression(&inline.Expr)
		fn.Kind = inline
	}
}

func (self *TypeExpander)expandExpression(expr *ast.Expression) {
	for i:= 0; i < len(expr.Instrs); i++ {
		self.expandInstr(expr.Instrs[i])
	}
}

func (self *TypeExpander)expandInstr(instr ast.Instruction) {
	var blockType *ast.BlockType
	switch ins := instr.(type) {
	case *ast.Block:
		blockType = &ins.BlockType
	case *ast.If:
		blockType = &ins.BlockType
	case *ast.Loop:
		blockType = &ins.BlockType
	case *ast.CallIndirect:
		self.expandTypeUse(&ins.Impl.Type)
	case *ast.ReturnCallIndirect:
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

func TypeToKey(ty ast.FunctionType) [2]string {
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

func (self *TypeExpander)registerType(val *ast.Type) {
	key := TypeToKey(val.Func)
	if _, ok := self.types[key]; ok == false {
		self.types[key] = self.ntypes
	}

	self.ntypes += 1
}

func (self *TypeExpander) expandImport(val *ast.Import) {
	if fn, ok := val.Item.(ast.ImportFunc); ok {
		self.expandTypeUse(&fn.TypeUse)
		val.Item = fn
	}
}

func (self *TypeExpander) expandTypeUse(val *ast.TypeUse) {
	if val.Index.IsSome() {
		return
	}
	key := TypeToKey(val.Type)
	if v, ok := self.types[key]; ok {
		val.Index = ast.NewOptionIndex(ast.NewNumIndex(v))
	} else {
		val.Index = ast.NewOptionIndex(ast.NewNumIndex( self.prepend(key) ))
	}
}

func (self *TypeExpander)prepend(key [2]string) uint32 {
	params , results := decodeKey(key)
	var funcParams []ast.FuncParam
	for _, p := range params {
		funcParams = append(funcParams, ast.FuncParam{
			Id  : ast.NoneOptionId(),
			Val : p,
		})
	}

	self.toPrepend = append(self.toPrepend, ast.Type{
		Name:ast.NoneOptionId(),
		Func: ast.FunctionType{
			Params:funcParams,
			Results:results,
		},
	})
	self.types[key] = self.ntypes
	self.ntypes += 1

	return self.ntypes - 1
}

func decodeKey(key [2]string) (params []ast.ValType,  results []ast.ValType) {
	for _, by := range []byte(key[0]) {
		params = append(params, ast.ValTypeFromByte(by))
	}
	for _, by := range []byte(key[1]) {
		results = append(results, ast.ValTypeFromByte(by))
	}

	return
}

