package ast

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
)

type Id struct {
	Name string
}

func (self *Id) Parse(ps *parser.ParserBuffer) error {
	id := ps.TryGetId()
	if len(id) == 0 {
		return errors.New("expect an identifier")
	}
	self.Name = id
	return nil
}

type OptionId struct {
	name Id
}

func NoneOptionId() OptionId {
	return OptionId{}
}

func (self *OptionId) IsSome() bool {
	return self.name.Name != ""
}

func (self OptionId) ToId() Id {
	if !self.IsSome() {
		panic("empty option id")
	}

	return self.name
}

func (self *OptionId) Parse(ps *parser.ParserBuffer) {
	_ = ps.TryParse(&self.name)
}

type Index struct {
	isnum bool
	Num   uint32
	Id    Id
}

func NewNumIndex(num uint32) Index {
	return Index{
		isnum: true,
		Num:   num,
	}
}

type OptionIndex struct {
	isSome bool
	index  Index
}

func (self *OptionIndex) Parse(ps *parser.ParserBuffer) {
	self.isSome = false
	if ps.TryParse(&self.index) == nil {
		self.isSome = true
	}
}

func (self *OptionIndex) IsSome() bool {
	return self.isSome
}

func NewOptionIndex(ind Index) OptionIndex {
	return OptionIndex{
		isSome: true,
		index:  ind,
	}
}

func NoneOptionIndex() OptionIndex {
	return OptionIndex{
		isSome: false,
	}
}

func (self OptionIndex) ToIndex() Index {
	if !self.isSome {
		panic("assert some index")
	}
	return self.index
}

func (self OptionIndex) ToIndexOr(ind Index) Index {
	if !self.isSome {
		return ind
	}
	return self.index
}

func (self *Index) Parse(ps *parser.ParserBuffer) error {
	id := ps.TryGetId()
	if len(id) != 0 {
		self.isnum = false
		self.Id = Id{Name: id}
		return nil
	}

	num, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.isnum = true
	self.Num = uint32(num)

	return nil
}

type Float32 struct {
	Bits uint32
}

func matchTokenType(token lexer.Token, ty lexer.TokenType) bool {
	return token != nil && token.Type() == ty
}

func string2f64(val lexer.Float) (uint64, error) {
	width := uint64(64)
	negOffset := width - 1
	expBits := uint64(11)
	expOffset := negOffset - expBits
	signifBits := width - 1 - expBits
	signifMask := uint64(1<<expOffset) - 1
	//bias := (1<<(expBits - 1)) - 1
	switch num := val.(type) {
	case lexer.Inf:
		exprBits := uint64((1 << expBits) - 1)
		negBits := uint64(0)
		if num.Neg {
			negBits = 1
		}
		return (negBits << negOffset) | (exprBits << expOffset), nil
	case lexer.Nan:
		exprBits := uint64((1 << expBits) - 1)
		negBits := uint64(0)
		if num.Neg {
			negBits = 1
		}
		signif := uint64(1 << (signifBits - 1))
		if num.SpecBit {
			signif = uint64(num.Val)
		}
		if signif&signifMask == 0 {
			return 0, errors.New("parse float64 error")
		}
		return (negBits << negOffset) | (exprBits << expOffset) | (signif & signifMask), nil
	case lexer.FloatVal:
		if !num.Hex {
			s := num.Integral
			if num.Decimal != "" {
				s += "." + num.Decimal
			}
			if num.Exponent != "" {
				s += "e" + num.Exponent
			}
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, err
			}
			// looks like the `*.wat` format considers infinite overflow to
			// be invalid.
			if math.IsInf(f, 0) {
				return 0, errors.New("parse float64 error, float infinite")
			}

			return math.Float64bits(f), nil
		}

		panic("todo: parse hex float not implemented yet")
	default:
		panic("unreachable")
	}
}

func string2f32(val lexer.Float) (uint32, error) {
	width := uint32(32)
	negOffset := width - 1
	expBits := uint32(8)
	expOffset := negOffset - expBits
	signifBits := width - 1 - expBits
	signifMask := uint32(1<<expOffset) - 1
	//bias := (1<<(expBits - 1)) - 1
	switch num := val.(type) {
	case lexer.Inf:
		exprBits := uint32((1 << expBits) - 1)
		negBits := uint32(0)
		if num.Neg {
			negBits = 1
		}
		return (negBits << negOffset) | (exprBits << expOffset), nil
	case lexer.Nan:
		exprBits := uint32((1 << expBits) - 1)
		negBits := uint32(0)
		if num.Neg {
			negBits = 1
		}
		signif := uint32(1 << (signifBits - 1))
		if num.SpecBit {
			signif = uint32(num.Val)
		}
		if signif&signifMask == 0 {
			return 0, errors.New("parse float 32 error")
		}
		return (negBits << negOffset) | (exprBits << expOffset) | (signif & signifMask), nil
	case lexer.FloatVal:
		if !num.Hex {
			s := num.Integral
			if num.Decimal != "" {
				s += "." + num.Decimal
			}
			if num.Exponent != "" {
				s += "e" + num.Exponent
			}
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return 0, err
			}
			// looks like the `*.wat` format considers infinite overflow to
			// be invalid.
			if math.IsInf(f, 0) {
				return 0, errors.New("parse float32 error, float infinite")
			}

			return math.Float32bits(float32(f)), nil
		}

		panic("todo: parse hex float not implemented yet")
	default:
		panic("unreachable")
	}
}

func (self *Float32) Parse(ps *parser.ParserBuffer) error {
	token := ps.PeekToken()
	if matchTokenType(token, lexer.FloatType) {
		val, err := ps.Float()
		if err != nil {
			return err
		}
		self.Bits, err = string2f32(val)
		return err
	} else if matchTokenType(token, lexer.IntegerType) {
		num, err := ps.ExpectInteger()
		if err != nil {
			return err
		}
		self.Bits, err = string2f32(lexer.FloatVal{Hex: num.Hex, Integral: num.Val, Decimal: "", Exponent: ""})
		return err
	}

	return fmt.Errorf("parse float32 error. expect number type")
}

type Float64 struct {
	Bits uint64
}

func (self *Float64) Parse(ps *parser.ParserBuffer) error {
	token := ps.PeekToken()
	if matchTokenType(token, lexer.FloatType) {
		val, err := ps.Float()
		if err != nil {
			return err
		}
		self.Bits, err = string2f64(val)
		return err
	} else if matchTokenType(token, lexer.IntegerType) {
		num, err := ps.ExpectInteger()
		if err != nil {
			return err
		}
		self.Bits, err = string2f64(lexer.FloatVal{Hex: num.Hex, Integral: num.Val, Decimal: "", Exponent: ""})
		return err
	}

	return fmt.Errorf("parse float64 error. expect number type")
}

type BlockType struct {
	Label Id
	Ty    TypeUse
}

func (self *BlockType) Parse(ps *parser.ParserBuffer) error {
	var id Id
	err := id.Parse(ps)
	if err != nil {
		return err
	}
	self.Label = id
	ty := TypeUse{}
	err = ty.ParseNoNames(ps)
	if err != nil {
		return err
	}
	self.Ty = ty
	return nil
}

type MemArg struct {
	Align  uint32
	Offset uint32
}

func (self *MemArg) Parse(ps *parser.ParserBuffer, defaultAlign uint32) error {
	parseField := func(name string, ps *parser.ParserBuffer) (some bool, val uint32, err error) {
		kw, err := ps.ExpectKeyword()
		if err != nil {
			return false, 0, err
		}
		if strings.HasPrefix(kw, name) == false {
			return false, 0, nil
		}
		kw = kw[len(name):]
		if strings.HasPrefix(kw, "=") == false {
			return false, 0, nil
		}
		kw = kw[1:]
		base := 10
		if strings.HasPrefix(kw, "0x") {
			base = 16
			kw = kw[2:]
		}
		value, err := strconv.ParseUint(kw, base, 32)
		if err != nil {
			return false, 0, err
		}
		return true, uint32(value), nil
	}

	some, offset, err := parseField("offset", ps)
	if err != nil {
		return err
	}
	if !some {
		offset = 0
	}
	self.Offset = offset
	some, align, err := parseField("align", ps)
	if err != nil {
		self.Align = defaultAlign
		return nil
	}
	if some && !isTwoPower(align) {
		return fmt.Errorf("alignment must be a power of two, %d", align)
	} else {
		align = defaultAlign
	}

	self.Align = align
	return nil
}

func isTwoPower(num uint32) bool {
	return num&(num-1) == 0
}
