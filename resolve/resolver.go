package resolve

import (
	"errors"
	"sort"

	"github.com/ontio/wast-parser/ast"
)

func Resolve(module *ast.Module) error {
	switch kind := module.Kind.(type) {
	case ast.ModuleKindBinary:
		return  nil
	case ast.ModuleKindText:
		expander := Expander{}
		kind.Fields = expander.Process(kind.Fields, func (expander *Expander, field ast.ModuleField) ast.ModuleField {
			return expander.ExpandImport(field)
		})
		kind.Fields = expander.Process(kind.Fields, func (expander *Expander, field ast.ModuleField) ast.ModuleField {
			return expander.ExpandExport(field)
		})

		for i := 1; i < len(kind.Fields); i++ {
			if _, imp := kind.Fields[i].(ast.Import); !imp {
				continue
			}
			switch kind.Fields[i-1].(type) {
			case ast.Memory, ast.Func, ast.Table, ast.Global:
				return errors.New("wrong import ordering")
			default:
				continue
			}
		}

		moveTypesFirst(kind.Fields)
		cur := 0
		typeExpander := TypeExpander{}
		for cur < len(kind.Fields) {
			kind.Fields[cur] = typeExpander.Expand(kind.Fields[cur])
			temp := kind.Fields[:cur]
			temp = append(temp, typeExpander.toPrepend...)
			temp = append(temp, kind.Fields[cur:]...)
			kind.Fields = temp
			cur += len(typeExpander.toPrepend) + 1
		}

		moveImportFirst(kind.Fields)
		namesResolver := NameResolver{}
		for i:=0; i < len(kind.Fields) ; i++ {
			namesResolver.Register(kind.Fields[i])
		}
		for i:=0; i < len(kind.Fields) ; i++ {
			var err error
			kind.Fields[i], err = namesResolver.Resolve(kind.Fields[i])
			if err != nil {
				return err
			}
		}

		return nil
	default:
		panic("unreachable")
	}
}

func moveImportFirst(fields []ast.ModuleField) {
	sort.SliceStable(fields, func (i,j int)bool {
		_, iok := fields[i].(ast.Import)
		_, jok := fields[j].(ast.Import)
		if iok && !jok {
			return true
		}

		return false
	})
}

func moveTypesFirst(fields []ast.ModuleField) {
	sort.SliceStable(fields, func (i,j int)bool {
		_, iok := fields[i].(ast.Type)
		_, jok := fields[j].(ast.Type)
		if iok && !jok {
			return true
		}

		return false
	})
}
