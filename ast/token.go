package ast

import (
	"errors"
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

func (self *OptionId) Parse(ps *parser.ParserBuffer) error {
	return ps.TryParse(&self.name)
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
	bits uint32
}

func (self *Float32) Parse(ps *parser.ParserBuffer) error {
	panic("todo")
}

type Float64 struct {
	bits uint64
}

func (self *Float64) Parse(ps *parser.ParserBuffer) error {
	panic("todo")
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
