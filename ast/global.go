package ast

import "github.com/ontio/wast-parser/parser"

type Global struct {
	implModuleField
	Name    OptionId
	Exports InlineExport
	ValType GlobalValType
	Kind    GlobalKind
}

type GlobalKind interface {
	globalKind()
}

type implGobalKind struct{}

func (self implGobalKind) globalKind() {}

type GlobalKindImport struct {
	implGobalKind
	Module string
	Field  string
}

type GlobalKindInline struct {
	implGobalKind
	Expr Expression
}

func (self *Global) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("global")
	if err != nil {
		return err
	}
	self.Name.Parse(ps)

	err = self.Exports.Parse(ps)
	if err != nil {
		return err
	}

	token := ps.Peek2Token()
	if matchKeyword(token, "import") {
		var imp GlobalKindImport
		err = ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("import")
			if err != nil {
				panic(err)
			}
			imp.Module, err = ps.ExpectString()
			if err != nil {
				return err
			}
			imp.Field, err = ps.ExpectString()
			return err
		})
		if err != nil {
			return err
		}

		err = self.ValType.Parse(ps)
		if err != nil {
			return err
		}

		self.Kind = imp
	} else {
		err := self.ValType.Parse(ps)
		if err != nil {
			return err
		}
		var expr Expression
		err = expr.Parse(ps)
		if err != nil {
			return err
		}
		self.Kind = GlobalKindInline{Expr: expr}
	}

	return nil
}
