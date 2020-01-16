package ast

import (
	"fmt"
	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Memory struct {
	implModuleField
	Name    OptionId
	Exports InlineExport
	Kind    MemoryKind
}

func (self *Memory) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("memory")
	if err != nil {
		return err
	}

	self.Name.Parse(ps)
	err = self.Exports.Parse(ps)
	if err != nil {
		return err
	}

	// Afterwards figure out which style this is, either:
	//  *   `(data ...)`
	//  *   `(import "a" "b") limits`
	//  *   `limits`
	if ps.PeekToken().Type() == lexer.LParenType {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			kw, err := ps.ExpectKeyword()
			if err != nil {
				return err
			}
			switch kw {
			case "data":
				self.Kind = &MemoryKindInline{}
				self.Kind.parseMemoryKindBody(ps)
			case "import":
				module, err := ps.ExpectString()
				if err != nil {
					return err
				}
				name, err := ps.ExpectString()
				if err != nil {
					return err
				}
				self.Kind = &MemoryKindImport{
					Module: module,
					Name:   name,
				}
			default:
				return fmt.Errorf("invalid import memory kind: %s", kw)
			}
			return nil
		})
		if err != nil {
			return err
		}
		switch mki := self.Kind.(type) {
		case *MemoryKindImport:
			return mki.Type.Parse(ps)
		default:
			return nil
		}
	}
	_, err = ps.ExpectUint32()
	if err != nil {
		return err
	}
	ps.StepBack(1)

	self.Kind = &MemoryKindNormal{}
	return self.Kind.parseMemoryKindBody(ps)
}

type MemoryKind interface {
	parseMemoryKindBody(ps *parser.ParserBuffer) error
}

type MemoryKindImport struct {
	Module string
	Name   string
	Type   MemoryType
}

func (self *MemoryKindImport) parseMemoryKindBody(ps *parser.ParserBuffer) error {
	mod, err := ps.ExpectString()
	if err != nil {
		return err
	}
	self.Module = mod
	name, err := ps.ExpectString()
	if err != nil {
		return err
	}
	self.Name = name

	return self.Type.Parse(ps)
}

type MemoryKindNormal struct {
	Type MemoryType
}

func (self *MemoryKindNormal) parseMemoryKindBody(ps *parser.ParserBuffer) error {
	return self.Type.Parse(ps)
}

type MemoryKindInline struct {
	Val [][]byte
}

func (self *MemoryKindInline) parseMemoryKindBody(ps *parser.ParserBuffer) error {
	data := make([][]byte, 0)
	for !ps.Empty() {
		str, err := ps.ExpectString()
		if err != nil {
			return err
		}
		data = append(data, []byte(str))
	}
	self.Val = data
	return nil
}
