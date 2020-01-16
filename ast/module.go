package ast

import (
	"fmt"
	"github.com/ontio/wast-parser/parser"
)

type Module struct {
	implWastDirective
	implWastExecute
	implQuoteModule
	Name OptionId
	Kind ModuleKind
}

func (self *Module) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("module")
	if err != nil {
		return err
	}

	self.Name.Parse(ps)
	if matchKeyword(ps.PeekToken(), "binary") {
		err := ps.ExpectKeywordMatch("binary")
		if err != nil {
			panic(err)
		}
		var data [][]byte
		for !ps.Empty() {
			str, err := ps.ExpectString()
			if err != nil {
				return err
			}
			data = append(data, []byte(str))
		}

		self.Kind = ModuleKindBinary{Bins: data}
	} else {
		var fields []ModuleField
		for !ps.Empty() {
			err = ps.Parens(func(ps *parser.ParserBuffer) error {
				field, err := parseModuleField(ps)
				if err != nil {
					return err
				}
				fields = append(fields, field)
				return nil
			})
			if err != nil {
				return err
			}
		}

		self.Kind = ModuleKindText{Fields: fields}
	}

	return nil
}

type ModuleKind interface {
	moduleKind()
}

type implModuleKind struct{}

func (self implModuleKind) moduleKind() {}

type ModuleKindText struct {
	implModuleKind
	Fields []ModuleField
}

type ModuleKindBinary struct {
	implModuleKind
	Bins [][]byte
}

type ModuleField interface {
	moduleField()
}

type StartField struct {
	implModuleField
	Index Index
}

type implModuleField struct{}

func (self implModuleField) moduleField() {}

func parseModuleField(ps *parser.ParserBuffer) (ModuleField, error) {
	kw, err := ps.PeekKeyword()
	if err != nil {
		return nil, err
	}
	fmt.Println("kwkw:", kw)
	var field ModuleField
	switch kw {
	case "type":
		var ty Type
		err = ty.Parse(ps)
		field = ty
	case "import":
		var imp Import
		err = imp.Parse(ps)
		field = imp
	case "func":
		var fun Func
		err = fun.Parse(ps)
		field = fun
	case "table":
		var table Table
		err = table.Parse(ps)
		field = table
	case "memory":
		var mem Memory
		err = mem.Parse(ps)
		field = mem
	case "global":
		var global Global
		err = global.Parse(ps)
		field = global
	case "export":
		var val Export
		err = val.Parse(ps)
		field = val
	case "start":
		var val StartField
		_ = ps.ExpectKeywordMatch("start")
		err = val.Index.Parse(ps)
		field = val
	case "elem":
		var val Elem
		err = val.Parse(ps)
		field = val
	case "data":
		var val Data
		err = val.Parse(ps)
		field = val
	default:
		return nil, fmt.Errorf("invalid module field keyword: %s", kw)
	}

	if err != nil {
		return nil, err
	}

	return field, nil
}
