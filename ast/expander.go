package ast

const pageSize = 1 << 16

type Expander struct {
	toAppend []ModuleField
	funcs    uint32
	memories uint32
	tables   uint32
	globals  uint32
}

func (self *Expander) Process(fields []ModuleField, fn func(expander *Expander, field ModuleField) ModuleField) []ModuleField {
	cur := 0
	for cur < len(fields) {
		fields[cur] = fn(self, fields[cur])
		var temp []ModuleField
		temp = append(temp, fields[:cur]...)
		temp = append(temp, self.toAppend...)
		temp = append(temp, fields[cur:]...)
		fields = temp
		cur += len(self.toAppend) + 1
		self.toAppend = nil
	}

	return fields
}

func (self *Expander) ExpandImport(item ModuleField) ModuleField {
	switch value := item.(type) {
	case Func:
		switch kind := value.Kind.(type) {
		case FuncKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					Export{
						Name:  name,
						Type:  ExportFunc,
						Index: NewNumIndex(self.funcs),
					})
			}

			self.funcs += 1
			return Import{
				Module: kind.Module,
				Field:  kind.Name,
				Id:     value.Name,
				Item:   ImportFunc{TypeUse: value.Type},
			}
		default:
			return item
		}
	case Memory:
		switch kind := value.Kind.(type) {
		case *MemoryKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					Export{
						Name:  name,
						Type:  ExportMemory,
						Index: NewNumIndex(self.memories),
					})
			}

			self.memories += 1
			return Import{
				Module: kind.Module,
				Field:  kind.Name,
				Id:     value.Name,
				Item:   ImportMemory{Mem: kind.Type},
			}
		default:
			return item
		}
	case Table:
		switch kind := value.Kind.(type) {
		case TableKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					Export{
						Name:  name,
						Type:  ExportTable,
						Index: NewNumIndex(self.tables),
					})
			}
			value.Exports.Names = nil

			self.tables += 1
			return Import{
				Module: kind.Module,
				Field:  kind.Name,
				Id:     value.Name,
				Item:   ImportTable{Table: kind.Type},
			}
		default:
			return item
		}
	case Global:
		switch kind := value.Kind.(type) {
		case GlobalKindImport:
			for _, name := range value.Exports.Names {
				self.toAppend = append(self.toAppend,
					Export{
						Name:  name,
						Type:  ExportGlobal,
						Index: NewNumIndex(self.globals),
					})
			}

			self.globals += 1
			return Import{
				Module: kind.Module,
				Field:  kind.Field,
				Id:     value.Name,
				Item:   ImportGlobal{Global: value.ValType},
			}
		default:
			return item
		}
	case Import:
		switch value.Item.(type) {
		case ImportFunc:
			self.funcs += 1
		case ImportMemory:
			self.memories += 1
		case ImportTable:
			self.tables += 1
		case ImportGlobal:
			self.globals += 1
		}

		return item
	default:
		return item
	}
}

func (self *Expander) ExpandExport(item ModuleField) ModuleField {
	switch value := item.(type) {
	case Func:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				Export{
					Name:  name,
					Type:  ExportFunc,
					Index: NewNumIndex(self.funcs),
				})
		}

		self.funcs += 1
		return value
	case Memory:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				Export{
					Name:  name,
					Type:  ExportMemory,
					Index: NewNumIndex(self.memories),
				})
		}

		if inline, ok := value.Kind.(*MemoryKindInline); ok {
			dataLen := 0
			for _, val := range inline.Val {
				dataLen += len(val)
			}
			pages := uint32(dataLen+pageSize-1) / pageSize
			value.Kind = &MemoryKindNormal{
				Type: MemoryType{
					Limits: Limits{
						Min: pages,
						Max: pages,
					},
					Shared: false,
				},
			}
			self.toAppend = append(self.toAppend, Data{
				Name: NoneOptionId(),
				Kind: DataKindActive{
					Memory: NewNumIndex(self.memories),
					Offset: Expression{
						Instrs: []Instruction{&I32Const{Val: 0}},
					},
				},
				Val: inline.Val,
			})
		}

		self.memories += 1
		return value
	case Table:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				Export{
					Name:  name,
					Type:  ExportTable,
					Index: NewNumIndex(self.tables),
				})
		}
		value.Exports.Names = nil
		if inline, ok := value.Kind.(TableKindInline); ok {
			var length uint32
			switch payload := inline.Payload.(type) {
			case ElemPayloadIndices:
				length = uint32(len(payload.Indices))
			case ElemPayloadExprs:
				length = uint32(len(payload.Exprs))
			}

			value.Kind = TableKindNormal{
				Type: TableType{
					Limits: Limits{
						Min: length,
						Max: length,
					},
					Elem: inline.Elem,
				},
			}

			self.toAppend = append(self.toAppend, Elem{
				Name: NoneOptionId(),
				Kind: ElemKindActive{
					Table: NewNumIndex(self.tables),
					Offset: Expression{
						Instrs: []Instruction{&I32Const{Val: 0}},
					},
				},
				Payload: inline.Payload,
			})
		}

		self.tables += 1
		return value
	case Global:
		for _, name := range value.Exports.Names {
			self.toAppend = append(self.toAppend,
				Export{
					Name:  name,
					Type:  ExportGlobal,
					Index: NewNumIndex(self.globals),
				})
		}
		value.Exports.Names = nil

		self.globals += 1
		return value
	default:
		return value
	}
}
