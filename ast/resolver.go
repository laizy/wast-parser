package ast

import (
	"errors"
	"sort"
)

func Resolve(module *Module) error {
	switch kind := module.Kind.(type) {
	case ModuleKindBinary:
		return nil
	case ModuleKindText:
		expander := Expander{}
		kind.Fields = expander.Process(kind.Fields, func(expander *Expander, field ModuleField) ModuleField {
			return expander.ExpandImport(field)
		})
		kind.Fields = expander.Process(kind.Fields, func(expander *Expander, field ModuleField) ModuleField {
			return expander.ExpandExport(field)
		})

		for i := 1; i < len(kind.Fields); i++ {
			if _, imp := kind.Fields[i].(Import); !imp {
				continue
			}
			switch kind.Fields[i-1].(type) {
			case Memory, Func, Table, Global:
				return errors.New("wrong import ordering")
			default:
				continue
			}
		}

		moveTypesFirst(kind.Fields)
		cur := 0
		typeExpander := NewTypeExpander()
		for cur < len(kind.Fields) {
			kind.Fields[cur] = typeExpander.Expand(kind.Fields[cur])
			var temp []ModuleField
			temp = append(temp, kind.Fields[:cur]...)
			temp = append(temp, typeExpander.toPrepend...)
			temp = append(temp, kind.Fields[cur:]...)
			kind.Fields = temp
			cur += len(typeExpander.toPrepend) + 1
		}

		moveImportFirst(kind.Fields)
		namesResolver := NewNameResolver()
		for i := 0; i < len(kind.Fields); i++ {
			namesResolver.Register(kind.Fields[i])
		}
		for i := 0; i < len(kind.Fields); i++ {
			var err error
			kind.Fields[i], err = namesResolver.Resolve(kind.Fields[i])
			if err != nil {
				return err
			}
		}

		module.Kind = kind
		return nil
	default:
		panic("unreachable")
	}
}

func moveImportFirst(fields []ModuleField) {
	sort.SliceStable(fields, func(i, j int) bool {
		_, iok := fields[i].(Import)
		_, jok := fields[j].(Import)
		if iok && !jok {
			return true
		}

		return false
	})
}

func moveTypesFirst(fields []ModuleField) {
	sort.SliceStable(fields, func(i, j int) bool {
		_, iok := fields[i].(Type)
		_, jok := fields[j].(Type)
		if iok && !jok {
			return true
		}

		return false
	})
}
