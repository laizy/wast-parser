package ast

import (
	"fmt"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Table struct {
	implModuleField
	Name    OptionId
	Exports InlineExport
	Kind    TableKind
}

type TableKind interface {
	tableKind()
}

type implTableKind struct{}

func (self implTableKind) tableKind() {}

func (self *Table) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("table")
	if err != nil {
		return err
	}
	self.Name.Parse(ps)
	err = self.Exports.Parse(ps)
	if err != nil {
		return err
	}

	// Afterwards figure out which style this is, either:
	//
	//  *   `elemtype (elem ...)`
	//  *   `(import "a" "b") limits`
	//  *   `limits`
	var elemType TableElemType
	if ps.TryParse(&elemType) == nil {
		return ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("elem")
			if err != nil {
				return err
			}
			if ps.PeekToken().Type() == lexer.LParenType {
				payload, err := parseElemPayloadExprs(ps, elemType)
				if err != nil {
					return err
				}

				self.Kind = TableKindInline{Elem: elemType, Payload: payload}
			} else {
				payload, err := parseElemPayloadIndices(ps)
				if err != nil {
					return err
				}

				self.Kind = TableKindInline{Elem: elemType, Payload: payload}
			}

			return nil
		})
	} else if ps.PeekToken().Type() == lexer.LParenType {
		var module, name string
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			err := ps.ExpectKeywordMatch("import")
			if err != nil {
				return err
			}
			module, err = ps.ExpectString()
			if err != nil {
				return err
			}
			name, err = ps.ExpectString()
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
		imp := TableKindImport{Module: module, Name: name}

		err = imp.Type.Parse(ps)
		if err != nil {
			return err
		}

		self.Kind = imp
		return nil
	} else if ps.PeekUint32() {
		var normal TableKindNormal
		err = normal.Type.Parse(ps)
		if err != nil {
			return err
		}
		self.Kind = normal
		return nil
	}
	return fmt.Errorf("table parse error")
}

type ElemPayload interface {
	implElemPayload()
	Encode(sink *ZeroCopySink)
}

func parseElemPayloadExprs(ps *parser.ParserBuffer, elemType TableElemType) (ElemPayload, error) {
	var exprs []OptionIndex
	for !ps.Empty() {
		var index OptionIndex
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			kw, err := ps.ExpectKeyword()
			if err != nil {
				return err
			}
			switch kw {
			case "ref_null":
				index = NoneOptionIndex()
			case "ref_func":
				var ind Index
				err := ind.Parse(ps)
				if err != nil {
					return err
				}
				index = NewOptionIndex(ind)
			default:
				return fmt.Errorf("parse elem payload error. invalid keyword: %s", kw)
			}

			return nil
		})
		if err != nil {
			return nil, err
		}

		exprs = append(exprs, index)
	}

	return ElemPayloadExprs{Exprs: exprs, Type: elemType}, nil
}

func parseElemPayloadIndices(ps *parser.ParserBuffer) (ElemPayload, error) {
	var indices []Index
	for !ps.Empty() {
		var index Index
		err := index.Parse(ps)
		if err != nil {
			return nil, err
		}

		indices = append(indices, index)
	}

	return ElemPayloadIndices{Indices: indices}, nil
}

func parseElemPayload(ps *parser.ParserBuffer) (ElemPayload, error) {
	var elemType TableElemType
	if ps.TryParse(&elemType) == nil {
		return parseElemPayloadExprs(ps, elemType)
	}

	return parseElemPayloadIndices(ps)
}

type ElemPayloadIndices struct {
	Indices []Index
}

func (self ElemPayloadIndices) implElemPayload() {}

type ElemPayloadExprs struct {
	Type  TableElemType
	Exprs []OptionIndex
}

func (self ElemPayloadExprs) implElemPayload() {}

type TableKindImport struct {
	implTableKind
	Module string
	Name   string
	Type   TableType
}

type TableKindNormal struct {
	implTableKind
	Type TableType
}

func (self *TableKindNormal) parseTableKindBody(ps *parser.ParserBuffer) error {
	return self.Type.Parse(ps)
}

type TableKindInline struct {
	implTableKind
	Elem    TableElemType
	Payload ElemPayload
}
