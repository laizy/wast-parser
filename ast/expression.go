package ast

import (
	"fmt"
	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Expression struct {
	Instrs []Instruction
}

func (self *Expression) Parse(ps *parser.ParserBuffer) error {
	var instrs instructions
	err := instrs.parseFoldedInstrs(ps)
	if err != nil {
		return err
	}
	self.Instrs = instrs.Instrs

	return nil
}

func (self *Expression) String() string {
	str := ""
	for _, inst := range self.Instrs {
		str += inst.String() + " "
	}

	return str
}

type Instruction interface {
	parseInstrBody(ps *parser.ParserBuffer) error
	String() string
}

type instructions struct {
	Instrs []Instruction
}

func (self *instructions) parseFoldedInstrs(ps *parser.ParserBuffer) error {
	for !ps.Empty() {
		err := self.parseOneInstr(ps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (self *instructions) parseOneInstr(ps *parser.ParserBuffer) error {
	if ps.PeekToken().Type() != lexer.LParenType {
		instr, err := parseInstr(ps)
		if err != nil {
			return err
		}
		self.Instrs = append(self.Instrs, instr)
		return nil
	}

	return ps.Parens(func(ps *parser.ParserBuffer) error {
		instr, err := parseInstr(ps)
		if err != nil {
			return err
		}
		switch val := instr.(type) {
		case *Block, *Loop: //loop
			self.Instrs = append(self.Instrs, val)

			err := self.parseFoldedInstrs(ps)
			if err != nil {
				return err
			}
			self.Instrs = append(self.Instrs, &End{Id: NoneOptionId()})
		case *If:
			if ps.PeekToken().Type() != lexer.LParenType {
				return fmt.Errorf("expected (")
			}
			if !matchKeyword(ps.Peek2Token(), "then") {
				err := self.parseOneInstr(ps)
				if err != nil {
					return err
				}
			}
			self.Instrs = append(self.Instrs, val)
			if ps.PeekToken().Type() != lexer.LParenType {
				return fmt.Errorf("expected `(`")
			}
			if matchKeyword(ps.Peek2Token(), "then") {
				err = ps.Parens(func(ps *parser.ParserBuffer) error {
					_ = ps.ExpectKeywordMatch("then")
					return self.parseFoldedInstrs(ps)
				})
				if err != nil {
					return err
				}
			} else {
				err := self.parseOneInstr(ps)
				if err != nil {
					return err
				}
			}
			if ps.PeekToken().Type() == lexer.LParenType {
				before := len(self.Instrs)
				self.Instrs = append(self.Instrs, &Else{})
				if matchKeyword(ps.Peek2Token(), "else") {
					err = ps.Parens(func(ps *parser.ParserBuffer) error {
						_ = ps.ExpectKeywordMatch("else")
						return self.parseFoldedInstrs(ps)
					})
					if err != nil {
						return err
					}
					if before+1 == len(self.Instrs) {
						self.Instrs = self.Instrs[:before]
					}
				} else {
					err := self.parseOneInstr(ps)
					if err != nil {
						return err
					}
				}
			}
			self.Instrs = append(self.Instrs, &End{})
		default:
			err := self.parseFoldedInstrs(ps)
			if err != nil {
				return err
			}

			self.Instrs = append(self.Instrs, val)
		}
		return nil
	})
}

type BrTableIndices struct {
	Labels  []Index
	Default Index
}

func (self *BrTableIndices) Parse(ps *parser.ParserBuffer) error {
	var index Index
	err := index.Parse(ps)
	if err != nil {
		return err
	}
	self.Labels = append(self.Labels, index)
	for !ps.Empty() {

		var index Index
		if ps.TryParse(&index) != nil {
			break
		}

		self.Labels = append(self.Labels, index)
	}
	self.Default = self.Labels[len(self.Labels)-1]
	self.Labels = self.Labels[:len(self.Labels)-1]
	return nil
}

type CallIndirectInner struct {
	Table Index
	Type  TypeUse
}

func (self *CallIndirectInner) Parse(ps *parser.ParserBuffer) error {
	var table OptionIndex
	table.Parse(ps)
	err := self.Type.ParseNoNames(ps)
	if err != nil {
		return err
	}
	// Turns out the official test suite at this time thinks table
	// identifiers comes first but wabt's test suites asserts differently
	// putting them second. Let's just handle both.
	if !table.IsSome() {
		table.Parse(ps)
	}
	self.Table = table.ToIndexOr(NewNumIndex(0))

	return nil
}

type SelectTypes struct {
	Types []ValType
}

func (self *SelectTypes) Parse(ps *parser.ParserBuffer) error {
	for matchKeyword(ps.Peek2Token(), "result") {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			_ = ps.ExpectKeywordMatch("result")
			for !ps.Empty() {
				var ty ValType
				err := ty.Parse(ps)
				if err != nil {
					return err
				}
				self.Types = append(self.Types, ty)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
