package resolve

import "github.com/ontio/wast-parser/ast"

const pageSize = 1<<16

type Expander struct {
	toAppend  []ast.ModuleField
	funcs uint32
	memories uint32
	tables uint32
	globals uint32
}

func (self *Expander)Process(fields []ast.ModuleField, fn func (expander *Expander, field ast.ModuleField) ast.ModuleField) []ast.ModuleField {
	cur := 0
	for cur < len(fields) {
		fields[cur] = fn(self, fields[cur])
		temp := fields[:cur]
		temp = append(temp, self.toAppend...)
		temp = fields[cur:]
		fields = temp
		cur += len(self.toAppend) + 1
		self.toAppend = nil
	}

	return fields
}

func (self *Expander)ExpandImport(item ast.ModuleField) ast.ModuleField {
	switch value := item.(type) {
	case ast.Func:
		switch kind := value.Kind.(type) {
		case ast.FuncKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					ast.Export{
						Name: name,
						Type: ast.ExportFunc,
						Index: ast.NewNumIndex(self.funcs),
					})
			}

			self.funcs += 1
			return ast.Import{
				Module:kind.Module,
				Field:kind.Name,
				Id:value.Name,
				Item: ast.ImportFunc{TypeUse:value.Type},
			}
		default:
			return item
		}
	case ast.Memory:
		switch kind := value.Kind.(type) {
		case *ast.MemoryKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					ast.Export{
						Name: name,
						Type: ast.ExportMemory,
						Index: ast.NewNumIndex(self.memories),
					})
			}

			self.memories += 1
			return ast.Import{
				Module:kind.Module,
				Field:kind.Name,
				Id:value.Name,
				Item: ast.ImportMemory{Mem:kind.Type},
			}
		default:
			return item
		}
	case ast.Table:
		switch kind := value.Kind.(type) {
		case *ast.TableKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					ast.Export{
						Name: name,
						Type: ast.ExportTable,
						Index: ast.NewNumIndex(self.tables),
					})
			}

			self.tables += 1
			return ast.Import{
				Module:kind.Module,
				Field:kind.Name,
				Id:value.Name,
				Item: ast.ImportTable{Table:kind.Type},
			}
		default:
			return item
		}
	case ast.Global:
		switch kind := value.Kind.(type) {
		case *ast.GlobalKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					ast.Export{
						Name: name,
						Type: ast.ExportGlobal,
						Index: ast.NewNumIndex(self.globals),
					})
			}

			self.globals += 1
			return ast.Import{
				Module:kind.Module,
				Field:kind.Field,
				Id:value.Name,
				Item: ast.ImportGlobal{Global:value.ValType},
			}
		default:
			return item
		}
	case ast.Import:
		switch value.Item.(type) {
		case ast.ImportFunc:
			self.funcs += 1
		case ast.ImportMemory:
			self.memories += 1
		case ast.ImportTable:
			self.tables += 1
		case ast.ImportGlobal:
			self.globals += 1
		}

		return item
	default:
		return item
	}
}

func (self *Expander)ExpandExport(item ast.ModuleField) ast.ModuleField {
	switch value := item.(type) {
	case ast.Func:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				ast.Export{
					Name:  name,
					Type:  ast.ExportFunc,
					Index: ast.NewNumIndex(self.funcs),
				})
		}

		self.funcs += 1
	case ast.Memory:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				ast.Export{
					Name:  name,
					Type:  ast.ExportMemory,
					Index: ast.NewNumIndex(self.memories),
				})
		}

		if inline, ok := value.Kind.(*ast.MemoryKindInline); ok {
			dataLen := 0
			for _, val := range inline.Val {
				dataLen += len(val)
			}
			pages := uint32(dataLen+pageSize-1) / pageSize
			value.Kind = &ast.MemoryKindNormal{
				Type: ast.MemoryType{
					Limits: ast.Limits{
						Min: pages,
						Max: pages,
					},
					Shared: false,
				},
			}
			self.toAppend = append(self.toAppend, ast.Data{
				Name: ast.NoneOptionId(),
				Kind: ast.DataKindActive{
					Memory: ast.NewNumIndex(self.memories),
					Offset: ast.Expression{
						Instrs: []ast.Instruction{&ast.I32Const{Val: 0}},
					},
				},
				Val: inline.Val,
			})
		}

		self.memories += 1
	case ast.Table:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				ast.Export{
					Name:  name,
					Type:  ast.ExportTable,
					Index: ast.NewNumIndex(self.tables),
				})
		}
		if inline, ok := value.Kind.(*ast.TableKindInline); ok {
			var length uint32
			switch payload := inline.Payload.(type) {
			case ast.ElemPayloadIndices:
				length = uint32(len(payload.Indices))
			case ast.ElemPayloadExprs:
				length = uint32(len(payload.Exprs))
			}

			value.Kind = &ast.TableKindNormal{
				Type: ast.TableType{
					Limits: ast.Limits{
						Min: length,
						Max: length,
					},
					Elem: inline.Elem,
				},
			}

			self.toAppend = append(self.toAppend, ast.Elem{
				Name: ast.NoneOptionId(),
				Kind: ast.ElemKindActive{
					Table: ast.NewNumIndex(self.tables),
					Offset: ast.Expression{
						Instrs: []ast.Instruction{&ast.I32Const{Val: 0}},
					},
				},
				Payload: inline.Payload,
			})
		}

		self.tables += 1
	case ast.Global:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				ast.Export{
					Name:  name,
					Type:  ast.ExportGlobal,
					Index: ast.NewNumIndex(self.globals),
				})
		}

		self.globals += 1
	default:
	}

	return item
}

