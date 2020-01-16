package ast

import (
	"fmt"

	"github.com/ontio/wast-parser/parser"
)

type Import struct {
	implModuleField
	Module string
	Field  string
	Id     OptionId
	Item   ImportItem
}

type ImportItem interface {
	ImportType() string
}

type ImportFunc struct {
	TypeUse TypeUse
}

func (self ImportFunc) ImportType() string { return "func" }

type ImportGlobal struct {
	Global GlobalValType
}

func (self ImportGlobal) ImportType() string { return "global" }

type ImportMemory struct {
	Mem MemoryType
}

func (self ImportMemory) ImportType() string { return "memory" }

type ImportTable struct {
	Table TableType
}

func (self ImportTable) ImportType() string { return "table" }

func (self *Import) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("import")
	if err != nil {
		return err
	}

	self.Module, err = ps.ExpectString()
	if err != nil {
		return err
	}
	self.Field, err = ps.ExpectString()
	if err != nil {
		return err
	}

	err = ps.Parens(func(ps *parser.ParserBuffer) error {
		kw, err := ps.ExpectKeyword()
		if err != nil {
			return err
		}
		self.Id.Parse(ps)
		switch kw {
		case "func":
			var typeUse TypeUse
			err := typeUse.Parse(ps)
			if err != nil {
				return err
			}
			self.Item = ImportFunc{TypeUse: typeUse}
		case "table":
			var table TableType
			err := table.Parse(ps)
			if err != nil {
				return err
			}
			self.Item = ImportTable{Table: table}
		case "memory":
			var memory MemoryType
			err := memory.Parse(ps)
			if err != nil {
				return err
			}

			self.Item = ImportMemory{Mem: memory}
		case "global":
			var global GlobalValType
			err := global.Parse(ps)
			if err != nil {
				return err
			}
			self.Item = ImportGlobal{Global: global}
		default:
			return fmt.Errorf("parse import item error: unexpected keyword: %s", kw)
		}

		return nil
	})

	return err
}
