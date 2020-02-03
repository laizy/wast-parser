package ast

import "github.com/ontio/wast-parser/parser"

type NanPattern interface {
	ImplementNanPattern()
}

type CanonicalNan struct {
	val string
}

func (self *CanonicalNan) ImplementNanPattern() {}

func (self *CanonicalNan) parseInstrBody(ps *parser.ParserBuffer) error {
	return nil
}
func (self *CanonicalNan) String() string {
	return self.val
}

type ArithmeticNan struct {
	val string
}

func (self *ArithmeticNan) ImplementNanPattern() {}
func (self *ArithmeticNan) parseInstrBody(ps *parser.ParserBuffer) error {
	return nil
}

func (self *ArithmeticNan) String() string {
	return self.val
}

type Value struct {
	val string
}

func (self *Value) ImplementNanPattern() {}
func (self *Value) parseInstrBody(ps *parser.ParserBuffer) error {
	var err error
	self.val, err = ps.ExpectString()
	if err != nil {
		return err
	}
	return nil
}

func (self *Value) String() string {
	return self.val
}
