package ast

import (
	"fmt"
	"github.com/ontio/wast-parser/parser"
)

type ExportType byte

const ExportFunc = ExportType(0)
const ExportTable = ExportType(1)
const ExportMemory = ExportType(2)
const ExportGlobal = ExportType(3)

type Export struct {
	implModuleField
	Name  string
	Type  ExportType
	Index Index
}

func (self *Export) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("export")
	if err != nil {
		return err
	}
	self.Name, err = ps.ExpectString()
	if err != nil {
		return err
	}
	return ps.Parens(func(ps *parser.ParserBuffer) error {
		kw, err := ps.ExpectKeyword()
		if err != nil {
			return err
		}

		switch kw {
		case "func":
			self.Type = ExportFunc
		case "table":
			self.Type = ExportTable
		case "memory":
			self.Type = ExportMemory
		case "global":
			self.Type = ExportGlobal
		default:
			return fmt.Errorf("parse export type error: unexpected keyword %s", kw)
		}

		return self.Index.Parse(ps)
	})
}

type InlineExport struct {
	Names []string
}

func (self *InlineExport) Parse(ps *parser.ParserBuffer) error {
	for {
		token := ps.Peek2Token()
		if !matchKeyword(token, "export") {
			break
		}
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("export")
			if err != nil {
				panic(err)
			}
			name, err := ps.ExpectString()
			if err != nil {
				return err
			}
			self.Names = append(self.Names, name)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
