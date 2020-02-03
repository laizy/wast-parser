package ast

import (
	"fmt"

	"github.com/ontio/wast-parser/parser"
)

type Block struct {
	BlockType BlockType
}

func (self *Block) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.BlockType.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Block) String() string {
	return "block"
}

type If struct {
	BlockType BlockType
}

func (self *If) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.BlockType.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *If) String() string {
	return "if"
}

type Else struct {
	Id OptionId
}

func (self *Else) parseInstrBody(ps *parser.ParserBuffer) error {
	self.Id.Parse(ps)

	return nil
}

func (self *Else) String() string {
	return "else"
}

type Loop struct {
	BlockType BlockType
}

func (self *Loop) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.BlockType.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Loop) String() string {
	return "loop"
}

type End struct {
	Id OptionId
}

func (self *End) parseInstrBody(ps *parser.ParserBuffer) error {
	self.Id.Parse(ps)

	return nil
}

func (self *End) String() string {
	return "end"
}

type Unreachable struct {
}

func (self *Unreachable) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *Unreachable) String() string {
	return "unreachable"
}

type Nop struct {
}

func (self *Nop) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *Nop) String() string {
	return "nop"
}

type Br struct {
	Index Index
}

func (self *Br) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Br) String() string {
	return "br"
}

type BrIf struct {
	Index Index
}

func (self *BrIf) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *BrIf) String() string {
	return "br_if"
}

type BrTable struct {
	Indices BrTableIndices
}

func (self *BrTable) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Indices.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *BrTable) String() string {
	return "br_table"
}

type Return struct {
}

func (self *Return) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *Return) String() string {
	return "return"
}

type Call struct {
	Index Index
}

func (self *Call) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Call) String() string {
	return "call"
}

type CallIndirect struct {
	Impl CallIndirectInner
}

func (self *CallIndirect) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Impl.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *CallIndirect) String() string {
	return "call_indirect"
}

type ReturnCall struct {
	Index Index
}

func (self *ReturnCall) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *ReturnCall) String() string {
	return "return_call"
}

type ReturnCallIndirect struct {
	Impl CallIndirectInner
}

func (self *ReturnCallIndirect) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Impl.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *ReturnCallIndirect) String() string {
	return "return_call_indirect"
}

type Drop struct {
}

func (self *Drop) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *Drop) String() string {
	return "drop"
}

type Select struct {
	SelectTypes SelectTypes
}

func (self *Select) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.SelectTypes.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *Select) String() string {
	return "select"
}

type LocalGet struct {
	Index Index
}

func (self *LocalGet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *LocalGet) String() string {
	return "local.get"
}

type LocalSet struct {
	Index Index
}

func (self *LocalSet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *LocalSet) String() string {
	return "local.set"
}

type LocalTee struct {
	Index Index
}

func (self *LocalTee) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *LocalTee) String() string {
	return "local.tee"
}

type GlobalGet struct {
	Index Index
}

func (self *GlobalGet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *GlobalGet) String() string {
	return "global.get"
}

type GlobalSet struct {
	Index Index
}

func (self *GlobalSet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *GlobalSet) String() string {
	return "global.set"
}

type TableGet struct {
	Index Index
}

func (self *TableGet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableGet) String() string {
	return "table.get"
}

type TableSet struct {
	Index Index
}

func (self *TableSet) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableSet) String() string {
	return "table.set"
}

type I32Load struct {
	MemArg MemArg
}

func (self *I32Load) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Load) String() string {
	return "i32.load"
}

type I64Load struct {
	MemArg MemArg
}

func (self *I64Load) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load) String() string {
	return "i64.load"
}

type F32Load struct {
	MemArg MemArg
}

func (self *F32Load) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *F32Load) String() string {
	return "f32.load"
}

type F64Load struct {
	MemArg MemArg
}

func (self *F64Load) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *F64Load) String() string {
	return "f64.load"
}

type I32Load8s struct {
	MemArg MemArg
}

func (self *I32Load8s) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Load8s) String() string {
	return "i32.load8_s"
}

type I32Load8u struct {
	MemArg MemArg
}

func (self *I32Load8u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Load8u) String() string {
	return "i32.load8_u"
}

type I32Load16s struct {
	MemArg MemArg
}

func (self *I32Load16s) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Load16s) String() string {
	return "i32.load16_s"
}

type I32Load16u struct {
	MemArg MemArg
}

func (self *I32Load16u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Load16u) String() string {
	return "i32.load16_u"
}

type I64Load8s struct {
	MemArg MemArg
}

func (self *I64Load8s) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load8s) String() string {
	return "i64.load8_s"
}

type I64Load8u struct {
	MemArg MemArg
}

func (self *I64Load8u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load8u) String() string {
	return "i64.load8_u"
}

type I64Load16s struct {
	MemArg MemArg
}

func (self *I64Load16s) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load16s) String() string {
	return "i64.load16_s"
}

type I64Load16u struct {
	MemArg MemArg
}

func (self *I64Load16u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load16u) String() string {
	return "i64.load16_u"
}

type I64Load32s struct {
	MemArg MemArg
}

func (self *I64Load32s) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load32s) String() string {
	return "i64.load32_s"
}

type I64Load32u struct {
	MemArg MemArg
}

func (self *I64Load32u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Load32u) String() string {
	return "i64.load32_u"
}

type I32Store struct {
	MemArg MemArg
}

func (self *I32Store) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Store) String() string {
	return "i32.store"
}

type I64Store struct {
	MemArg MemArg
}

func (self *I64Store) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Store) String() string {
	return "i64.store"
}

type F32Store struct {
	MemArg MemArg
}

func (self *F32Store) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *F32Store) String() string {
	return "f32.store"
}

type F64Store struct {
	MemArg MemArg
}

func (self *F64Store) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *F64Store) String() string {
	return "f64.store"
}

type I32Store8 struct {
	MemArg MemArg
}

func (self *I32Store8) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Store8) String() string {
	return "i32.store8"
}

type I32Store16 struct {
	MemArg MemArg
}

func (self *I32Store16) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32Store16) String() string {
	return "i32.store16"
}

type I64Store8 struct {
	MemArg MemArg
}

func (self *I64Store8) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Store8) String() string {
	return "i64.store8"
}

type I64Store16 struct {
	MemArg MemArg
}

func (self *I64Store16) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Store16) String() string {
	return "i64.store16"
}

type I64Store32 struct {
	MemArg MemArg
}

func (self *I64Store32) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64Store32) String() string {
	return "i64.store32"
}

type MemorySize struct {
}

func (self *MemorySize) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *MemorySize) String() string {
	return "memory.size"
}

type MemoryGrow struct {
}

func (self *MemoryGrow) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *MemoryGrow) String() string {
	return "memory.grow"
}

type MemoryCopy struct {
}

func (self *MemoryCopy) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *MemoryCopy) String() string {
	return "memory.copy"
}

type MemoryFill struct {
}

func (self *MemoryFill) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *MemoryFill) String() string {
	return "memory.fill"
}

type DataDrop struct {
	Index Index
}

func (self *DataDrop) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *DataDrop) String() string {
	return "data.drop"
}

type ElemDrop struct {
	Index Index
}

func (self *ElemDrop) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *ElemDrop) String() string {
	return "elem.drop"
}

type TableCopy struct {
}

func (self *TableCopy) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *TableCopy) String() string {
	return "table.copy"
}

type TableFill struct {
	Index Index
}

func (self *TableFill) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableFill) String() string {
	return "table.fill"
}

type TableSize struct {
	Index Index
}

func (self *TableSize) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableSize) String() string {
	return "table.size"
}

type TableGrow struct {
	Index Index
}

func (self *TableGrow) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *TableGrow) String() string {
	return "table.grow"
}

type RefNull struct {
}

func (self *RefNull) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *RefNull) String() string {
	return "ref.null"
}

type RefIsNull struct {
}

func (self *RefIsNull) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *RefIsNull) String() string {
	return "ref.is_null"
}

type RefHost struct {
	Val uint32
}

func (self *RefHost) parseInstrBody(ps *parser.ParserBuffer) error {
	val, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.Val = val

	return nil
}

func (self *RefHost) String() string {
	return "ref.host"
}

type RefFunc struct {
	Index Index
}

func (self *RefFunc) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Index.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *RefFunc) String() string {
	return "ref.func"
}

type I32Const struct {
	Val uint32
}

func (self *I32Const) parseInstrBody(ps *parser.ParserBuffer) error {
	val, err := ps.ExpectUint32()
	if err != nil {
		return err
	}
	self.Val = val

	return nil
}

func (self *I32Const) String() string {
	return "i32.const"
}

type I64Const struct {
	Val int64
}

func (self *I64Const) parseInstrBody(ps *parser.ParserBuffer) error {
	val, err := ps.ExpectInt64()
	if err != nil {
		return err
	}
	self.Val = val

	return nil
}

func (self *I64Const) String() string {
	return "i64.const"
}

type F32Const struct {
	Val Float32
}

func (self *F32Const) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Val.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *F32Const) String() string {
	return "f32.const"
}

type F64Const struct {
	Val Float64
}

func (self *F64Const) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.Val.Parse(ps)
	if err != nil {
		return err
	}

	return nil
}

func (self *F64Const) String() string {
	return "f64.const"
}

type I32Clz struct {
}

func (self *I32Clz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Clz) String() string {
	return "i32.clz"
}

type I32Ctz struct {
}

func (self *I32Ctz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Ctz) String() string {
	return "i32.ctz"
}

type I32Pocnt struct {
}

func (self *I32Pocnt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Pocnt) String() string {
	return "i32.popcnt"
}

type I32Add struct {
}

func (self *I32Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Add) String() string {
	return "i32.add"
}

type I32Sub struct {
}

func (self *I32Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Sub) String() string {
	return "i32.sub"
}

type I32Mul struct {
}

func (self *I32Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Mul) String() string {
	return "i32.mul"
}

type I32DivS struct {
}

func (self *I32DivS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32DivS) String() string {
	return "i32.div_s"
}

type I32DivU struct {
}

func (self *I32DivU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32DivU) String() string {
	return "i32.div_u"
}

type I32RemS struct {
}

func (self *I32RemS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32RemS) String() string {
	return "i32.rem_s"
}

type I32RemU struct {
}

func (self *I32RemU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32RemU) String() string {
	return "i32.rem_u"
}

type I32And struct {
}

func (self *I32And) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32And) String() string {
	return "i32.and"
}

type I32Or struct {
}

func (self *I32Or) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Or) String() string {
	return "i32.or"
}

type I32Xor struct {
}

func (self *I32Xor) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Xor) String() string {
	return "i32.xor"
}

type I32Shl struct {
}

func (self *I32Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Shl) String() string {
	return "i32.shl"
}

type I32ShrS struct {
}

func (self *I32ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32ShrS) String() string {
	return "i32.shr_s"
}

type I32ShrU struct {
}

func (self *I32ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32ShrU) String() string {
	return "i32.shr_u"
}

type I32Rotl struct {
}

func (self *I32Rotl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Rotl) String() string {
	return "i32.rotl"
}

type I32Rotr struct {
}

func (self *I32Rotr) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Rotr) String() string {
	return "i32.rotr"
}

type I64Clz struct {
}

func (self *I64Clz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Clz) String() string {
	return "i64.clz"
}

type I64Ctz struct {
}

func (self *I64Ctz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Ctz) String() string {
	return "i64.ctz"
}

type I64Popcnt struct {
}

func (self *I64Popcnt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Popcnt) String() string {
	return "i64.popcnt"
}

type I64Add struct {
}

func (self *I64Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Add) String() string {
	return "i64.add"
}

type I64Sub struct {
}

func (self *I64Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Sub) String() string {
	return "i64.sub"
}

type I64Mul struct {
}

func (self *I64Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Mul) String() string {
	return "i64.mul"
}

type I64DivS struct {
}

func (self *I64DivS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64DivS) String() string {
	return "i64.div_s"
}

type I64DivU struct {
}

func (self *I64DivU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64DivU) String() string {
	return "i64.div_u"
}

type I64RemS struct {
}

func (self *I64RemS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64RemS) String() string {
	return "i64.rem_s"
}

type I64RemU struct {
}

func (self *I64RemU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64RemU) String() string {
	return "i64.rem_u"
}

type I64And struct {
}

func (self *I64And) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64And) String() string {
	return "i64.and"
}

type I64Or struct {
}

func (self *I64Or) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Or) String() string {
	return "i64.or"
}

type I64Xor struct {
}

func (self *I64Xor) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Xor) String() string {
	return "i64.xor"
}

type I64Shl struct {
}

func (self *I64Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Shl) String() string {
	return "i64.shl"
}

type I64ShrS struct {
}

func (self *I64ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64ShrS) String() string {
	return "i64.shr_s"
}

type I64ShrU struct {
}

func (self *I64ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64ShrU) String() string {
	return "i64.shr_u"
}

type I64Rotl struct {
}

func (self *I64Rotl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Rotl) String() string {
	return "i64.rotl"
}

type I64Rotr struct {
}

func (self *I64Rotr) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Rotr) String() string {
	return "i64.rotr"
}

type F32Abs struct {
}

func (self *F32Abs) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Abs) String() string {
	return "f32.abs"
}

type F32Neg struct {
}

func (self *F32Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Neg) String() string {
	return "f32.neg"
}

type F32Ceil struct {
}

func (self *F32Ceil) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Ceil) String() string {
	return "f32.ceil"
}

type F32Floor struct {
}

func (self *F32Floor) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Floor) String() string {
	return "f32.floor"
}

type F32Trunc struct {
}

func (self *F32Trunc) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Trunc) String() string {
	return "f32.trunc"
}

type F32Nearest struct {
}

func (self *F32Nearest) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Nearest) String() string {
	return "f32.nearest"
}

type F32Sqrt struct {
}

func (self *F32Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Sqrt) String() string {
	return "f32.sqrt"
}

type F32Add struct {
}

func (self *F32Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Add) String() string {
	return "f32.add"
}

type F32Sub struct {
}

func (self *F32Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Sub) String() string {
	return "f32.sub"
}

type F32Mul struct {
}

func (self *F32Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Mul) String() string {
	return "f32.mul"
}

type F32Div struct {
}

func (self *F32Div) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Div) String() string {
	return "f32.div"
}

type F32Min struct {
}

func (self *F32Min) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Min) String() string {
	return "f32.min"
}

type F32Max struct {
}

func (self *F32Max) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Max) String() string {
	return "f32.max"
}

type F32Copysign struct {
}

func (self *F32Copysign) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Copysign) String() string {
	return "f32.copysign"
}

type F64Abs struct {
}

func (self *F64Abs) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Abs) String() string {
	return "f64.abs"
}

type F64Neg struct {
}

func (self *F64Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Neg) String() string {
	return "f64.neg"
}

type F64Ceil struct {
}

func (self *F64Ceil) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Ceil) String() string {
	return "f64.ceil"
}

type F64Floor struct {
}

func (self *F64Floor) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Floor) String() string {
	return "f64.floor"
}

type F64Trunc struct {
}

func (self *F64Trunc) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Trunc) String() string {
	return "f64.trunc"
}

type F64Nearest struct {
}

func (self *F64Nearest) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Nearest) String() string {
	return "f64.nearest"
}

type F64Sqrt struct {
}

func (self *F64Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Sqrt) String() string {
	return "f64.sqrt"
}

type F64Add struct {
}

func (self *F64Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Add) String() string {
	return "f64.add"
}

type F64Sub struct {
}

func (self *F64Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Sub) String() string {
	return "f64.sub"
}

type F64Mul struct {
}

func (self *F64Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Mul) String() string {
	return "f64.mul"
}

type F64Div struct {
}

func (self *F64Div) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Div) String() string {
	return "f64.div"
}

type F64Min struct {
}

func (self *F64Min) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Min) String() string {
	return "f64.min"
}

type F64Max struct {
}

func (self *F64Max) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Max) String() string {
	return "f64.max"
}

type F64Copysign struct {
}

func (self *F64Copysign) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Copysign) String() string {
	return "f64.copysign"
}

type I32Eqz struct {
}

func (self *I32Eqz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Eqz) String() string {
	return "i32.eqz"
}

type I32Eq struct {
}

func (self *I32Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Eq) String() string {
	return "i32.eq"
}

type I32Ne struct {
}

func (self *I32Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Ne) String() string {
	return "i32.ne"
}

type I32LtS struct {
}

func (self *I32LtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32LtS) String() string {
	return "i32.lt_s"
}

type I32LtU struct {
}

func (self *I32LtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32LtU) String() string {
	return "i32.lt_u"
}

type I32GtS struct {
}

func (self *I32GtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32GtS) String() string {
	return "i32.gt_s"
}

type I32GtU struct {
}

func (self *I32GtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32GtU) String() string {
	return "i32.gt_u"
}

type I32LeS struct {
}

func (self *I32LeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32LeS) String() string {
	return "i32.le_s"
}

type I32LeU struct {
}

func (self *I32LeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32LeU) String() string {
	return "i32.le_u"
}

type I32GeS struct {
}

func (self *I32GeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32GeS) String() string {
	return "i32.ge_s"
}

type I32GeU struct {
}

func (self *I32GeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32GeU) String() string {
	return "i32.ge_u"
}

type I64Eqz struct {
}

func (self *I64Eqz) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Eqz) String() string {
	return "i64.eqz"
}

type I64Eq struct {
}

func (self *I64Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Eq) String() string {
	return "i64.eq"
}

type I64Ne struct {
}

func (self *I64Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Ne) String() string {
	return "i64.ne"
}

type I64LtS struct {
}

func (self *I64LtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64LtS) String() string {
	return "i64.lt_s"
}

type I64LtU struct {
}

func (self *I64LtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64LtU) String() string {
	return "i64.lt_u"
}

type I64GtS struct {
}

func (self *I64GtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64GtS) String() string {
	return "i64.gt_s"
}

type I64GtU struct {
}

func (self *I64GtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64GtU) String() string {
	return "i64.gt_u"
}

type I64LeS struct {
}

func (self *I64LeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64LeS) String() string {
	return "i64.le_s"
}

type I64LeU struct {
}

func (self *I64LeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64LeU) String() string {
	return "i64.le_u"
}

type I64GeS struct {
}

func (self *I64GeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64GeS) String() string {
	return "i64.ge_s"
}

type I64GeU struct {
}

func (self *I64GeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64GeU) String() string {
	return "i64.ge_u"
}

type F32Eq struct {
}

func (self *F32Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Eq) String() string {
	return "f32.eq"
}

type F32Ne struct {
}

func (self *F32Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Ne) String() string {
	return "f32.ne"
}

type F32Lt struct {
}

func (self *F32Lt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Lt) String() string {
	return "f32.lt"
}

type F32Gt struct {
}

func (self *F32Gt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Gt) String() string {
	return "f32.gt"
}

type F32Le struct {
}

func (self *F32Le) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Le) String() string {
	return "f32.le"
}

type F32Ge struct {
}

func (self *F32Ge) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32Ge) String() string {
	return "f32.ge"
}

type F64Eq struct {
}

func (self *F64Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Eq) String() string {
	return "f64.eq"
}

type F64Ne struct {
}

func (self *F64Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Ne) String() string {
	return "f64.ne"
}

type F64Lt struct {
}

func (self *F64Lt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Lt) String() string {
	return "f64.lt"
}

type F64Gt struct {
}

func (self *F64Gt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Gt) String() string {
	return "f64.gt"
}

type F64Le struct {
}

func (self *F64Le) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Le) String() string {
	return "f64.le"
}

type F64Ge struct {
}

func (self *F64Ge) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64Ge) String() string {
	return "f64.ge"
}

type I32WrapI64 struct {
}

func (self *I32WrapI64) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32WrapI64) String() string {
	return "i32.wrap_i64"
}

type I32TruncF32S struct {
}

func (self *I32TruncF32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncF32S) String() string {
	return "i32.trunc_f32_s"
}

type I32TruncF32U struct {
}

func (self *I32TruncF32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncF32U) String() string {
	return "i32.trunc_f32_u"
}

type I32TruncF64S struct {
}

func (self *I32TruncF64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncF64S) String() string {
	return "i32.trunc_f64_s"
}

type I32TruncF64U struct {
}

func (self *I32TruncF64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncF64U) String() string {
	return "i32.trunc_f64_u"
}

type I64ExtendI32S struct {
}

func (self *I64ExtendI32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64ExtendI32S) String() string {
	return "i64.extend_i32_s"
}

type I64ExtendI32U struct {
}

func (self *I64ExtendI32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64ExtendI32U) String() string {
	return "i64.extend_i32_u"
}

type I64TruncF32S struct {
}

func (self *I64TruncF32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncF32S) String() string {
	return "i64.trunc_f32_s"
}

type I64TruncF32U struct {
}

func (self *I64TruncF32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncF32U) String() string {
	return "i64.trunc_f32_u"
}

type I64TruncF64S struct {
}

func (self *I64TruncF64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncF64S) String() string {
	return "i64.trunc_f64_s"
}

type I64TruncF64U struct {
}

func (self *I64TruncF64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncF64U) String() string {
	return "i64.trunc_f64_u"
}

type F32ConvertI32S struct {
}

func (self *F32ConvertI32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32ConvertI32S) String() string {
	return "f32.convert_i32_s"
}

type F32ConvertI32U struct {
}

func (self *F32ConvertI32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32ConvertI32U) String() string {
	return "f32.convert_i32_u"
}

type F32ConvertI64S struct {
}

func (self *F32ConvertI64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32ConvertI64S) String() string {
	return "f32.convert_i64_s"
}

type F32ConvertI64U struct {
}

func (self *F32ConvertI64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32ConvertI64U) String() string {
	return "f32.convert_i64_u"
}

type F32DemoteF64 struct {
}

func (self *F32DemoteF64) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32DemoteF64) String() string {
	return "f32.demote_f64"
}

type F64ConvertI32S struct {
}

func (self *F64ConvertI32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64ConvertI32S) String() string {
	return "f64.convert_i32_s"
}

type F64ConvertI32U struct {
}

func (self *F64ConvertI32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64ConvertI32U) String() string {
	return "f64.convert_i32_u"
}

type F64ConvertI64S struct {
}

func (self *F64ConvertI64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64ConvertI64S) String() string {
	return "f64.convert_i64_s"
}

type F64ConvertI64U struct {
}

func (self *F64ConvertI64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64ConvertI64U) String() string {
	return "f64.convert_i64_u"
}

type F64PromoteF32 struct {
}

func (self *F64PromoteF32) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64PromoteF32) String() string {
	return "f64.promote_f32"
}

type I32ReinterpretF32 struct {
}

func (self *I32ReinterpretF32) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32ReinterpretF32) String() string {
	return "i32.reinterpret_f32"
}

type I64ReinterpretF64 struct {
}

func (self *I64ReinterpretF64) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64ReinterpretF64) String() string {
	return "i64.reinterpret_f64"
}

type F32ReinterpretI32 struct {
}

func (self *F32ReinterpretI32) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32ReinterpretI32) String() string {
	return "f32.reinterpret_i32"
}

type F64ReinterpretI64 struct {
}

func (self *F64ReinterpretI64) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64ReinterpretI64) String() string {
	return "f64.reinterpret_i64"
}

type I32TruncSatF32S struct {
}

func (self *I32TruncSatF32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncSatF32S) String() string {
	return "i32.trunc_sat_f32_s"
}

type I32TruncSatF32U struct {
}

func (self *I32TruncSatF32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncSatF32U) String() string {
	return "i32.trunc_sat_f32_u"
}

type I32TruncSatF64S struct {
}

func (self *I32TruncSatF64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncSatF64S) String() string {
	return "i32.trunc_sat_f64_s"
}

type I32TruncSatF64U struct {
}

func (self *I32TruncSatF64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32TruncSatF64U) String() string {
	return "i32.trunc_sat_f64_u"
}

type I64TruncSatF32S struct {
}

func (self *I64TruncSatF32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncSatF32S) String() string {
	return "i64.trunc_sat_f32_s"
}

type I64TruncSatF32U struct {
}

func (self *I64TruncSatF32U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncSatF32U) String() string {
	return "i64.trunc_sat_f32_u"
}

type I64TruncSatF64S struct {
}

func (self *I64TruncSatF64S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncSatF64S) String() string {
	return "i64.trunc_sat_f64_s"
}

type I64TruncSatF64U struct {
}

func (self *I64TruncSatF64U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64TruncSatF64U) String() string {
	return "i64.trunc_sat_f64_u"
}

type I32Extend8S struct {
}

func (self *I32Extend8S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Extend8S) String() string {
	return "i32.extend8_s"
}

type I32Extend16S struct {
}

func (self *I32Extend16S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32Extend16S) String() string {
	return "i32.extend16_s"
}

type I64Extend8S struct {
}

func (self *I64Extend8S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Extend8S) String() string {
	return "i64.extend8_s"
}

type I64Extend16S struct {
}

func (self *I64Extend16S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Extend16S) String() string {
	return "i64.extend16_s"
}

type I64Extend32S struct {
}

func (self *I64Extend32S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64Extend32S) String() string {
	return "i64.extend32_s"
}

type AtomicNotify struct {
	MemArg MemArg
}

func (self *AtomicNotify) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *AtomicNotify) String() string {
	return "atomic.notify"
}

type I32AtomicWait struct {
	MemArg MemArg
}

func (self *I32AtomicWait) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicWait) String() string {
	return "i32.atomic.wait"
}

type I64AtomicWait struct {
	MemArg MemArg
}

func (self *I64AtomicWait) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicWait) String() string {
	return "i64.atomic.wait"
}

type AtomicFence struct {
}

func (self *AtomicFence) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *AtomicFence) String() string {
	return "atomic.fence"
}

type I32AtomicLoad struct {
	MemArg MemArg
}

func (self *I32AtomicLoad) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicLoad) String() string {
	return "i32.atomic.load"
}

type I64AtomicLoad struct {
	MemArg MemArg
}

func (self *I64AtomicLoad) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicLoad) String() string {
	return "i64.atomic.load"
}

type I32AtomicLoad8u struct {
	MemArg MemArg
}

func (self *I32AtomicLoad8u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicLoad8u) String() string {
	return "i32.atomic.load8_u"
}

type I32AtomicLoad16u struct {
	MemArg MemArg
}

func (self *I32AtomicLoad16u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicLoad16u) String() string {
	return "i32.atomic.load16_u"
}

type I64AtomicLoad8u struct {
	MemArg MemArg
}

func (self *I64AtomicLoad8u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicLoad8u) String() string {
	return "i64.atomic.load8_u"
}

type I64AtomicLoad16u struct {
	MemArg MemArg
}

func (self *I64AtomicLoad16u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicLoad16u) String() string {
	return "i64.atomic.load16_u"
}

type I64AtomicLoad32u struct {
	MemArg MemArg
}

func (self *I64AtomicLoad32u) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicLoad32u) String() string {
	return "i64.atomic.load32_u"
}

type I32AtomicStore struct {
	MemArg MemArg
}

func (self *I32AtomicStore) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicStore) String() string {
	return "i32.atomic.store"
}

type I64AtomicStore struct {
	MemArg MemArg
}

func (self *I64AtomicStore) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicStore) String() string {
	return "i64.atomic.store"
}

type I32AtomicStore8 struct {
	MemArg MemArg
}

func (self *I32AtomicStore8) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicStore8) String() string {
	return "i32.atomic.store8"
}

type I32AtomicStore16 struct {
	MemArg MemArg
}

func (self *I32AtomicStore16) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicStore16) String() string {
	return "i32.atomic.store16"
}

type I64AtomicStore8 struct {
	MemArg MemArg
}

func (self *I64AtomicStore8) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicStore8) String() string {
	return "i64.atomic.store8"
}

type I64AtomicStore16 struct {
	MemArg MemArg
}

func (self *I64AtomicStore16) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicStore16) String() string {
	return "i64.atomic.store16"
}

type I64AtomicStore32 struct {
	MemArg MemArg
}

func (self *I64AtomicStore32) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicStore32) String() string {
	return "i64.atomic.store32"
}

type I32AtomicRmwAdd struct {
	MemArg MemArg
}

func (self *I32AtomicRmwAdd) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwAdd) String() string {
	return "i32.atomic.rmw.add"
}

type I64AtomicRmwAdd struct {
	MemArg MemArg
}

func (self *I64AtomicRmwAdd) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwAdd) String() string {
	return "i64.atomic.rmw.add"
}

type I32AtomicRmw8AddU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8AddU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8AddU) String() string {
	return "i32.atomic.rmw8.add_u"
}

type I32AtomicRmw16AddU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16AddU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16AddU) String() string {
	return "i32.atomic.rmw16.add_u"
}

type I64AtomicRmw8AddU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8AddU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8AddU) String() string {
	return "i64.atomic.rmw8.add_u"
}

type I64AtomicRmw16AddU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16AddU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16AddU) String() string {
	return "i64.atomic.rmw16.add_u"
}

type I64AtomicRmw32AddU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32AddU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32AddU) String() string {
	return "i64.atomic.rmw32.add_u"
}

type I32AtomicRmwSub struct {
	MemArg MemArg
}

func (self *I32AtomicRmwSub) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwSub) String() string {
	return "i32.atomic.rmw.sub"
}

type I64AtomicRmwSub struct {
	MemArg MemArg
}

func (self *I64AtomicRmwSub) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwSub) String() string {
	return "i64.atomic.rmw.sub"
}

type I32AtomicRmw8SubU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8SubU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8SubU) String() string {
	return "i32.atomic.rmw8.sub_u"
}

type I32AtomicRmw16SubU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16SubU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16SubU) String() string {
	return "i32.atomic.rmw16.sub_u"
}

type I64AtomicRmw8SubU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8SubU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8SubU) String() string {
	return "i64.atomic.rmw8.sub_u"
}

type I64AtomicRmw16SubU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16SubU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16SubU) String() string {
	return "i64.atomic.rmw16.sub_u"
}

type I64AtomicRmw32SubU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32SubU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32SubU) String() string {
	return "i64.atomic.rmw32.sub_u"
}

type I32AtomicRmwAnd struct {
	MemArg MemArg
}

func (self *I32AtomicRmwAnd) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwAnd) String() string {
	return "i32.atomic.rmw.and"
}

type I64AtomicRmwAnd struct {
	MemArg MemArg
}

func (self *I64AtomicRmwAnd) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwAnd) String() string {
	return "i64.atomic.rmw.and"
}

type I32AtomicRmw8AndU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8AndU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8AndU) String() string {
	return "i32.atomic.rmw8.and_u"
}

type I32AtomicRmw16AndU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16AndU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16AndU) String() string {
	return "i32.atomic.rmw16.and_u"
}

type I64AtomicRmw8AndU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8AndU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8AndU) String() string {
	return "i64.atomic.rmw8.and_u"
}

type I64AtomicRmw16AndU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16AndU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16AndU) String() string {
	return "i64.atomic.rmw16.and_u"
}

type I64AtomicRmw32AndU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32AndU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32AndU) String() string {
	return "i64.atomic.rmw32.and_u"
}

type I32AtomicRmwOr struct {
	MemArg MemArg
}

func (self *I32AtomicRmwOr) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwOr) String() string {
	return "i32.atomic.rmw.or"
}

type I64AtomicRmwOr struct {
	MemArg MemArg
}

func (self *I64AtomicRmwOr) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwOr) String() string {
	return "i64.atomic.rmw.or"
}

type I32AtomicRmw8OrU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8OrU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8OrU) String() string {
	return "i32.atomic.rmw8.or_u"
}

type I32AtomicRmw16OrU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16OrU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16OrU) String() string {
	return "i32.atomic.rmw16.or_u"
}

type I64AtomicRmw8OrU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8OrU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8OrU) String() string {
	return "i64.atomic.rmw8.or_u"
}

type I64AtomicRmw16OrU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16OrU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16OrU) String() string {
	return "i64.atomic.rmw16.or_u"
}

type I64AtomicRmw32OrU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32OrU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32OrU) String() string {
	return "i64.atomic.rmw32.or_u"
}

type I32AtomicRmwXor struct {
	MemArg MemArg
}

func (self *I32AtomicRmwXor) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwXor) String() string {
	return "i32.atomic.rmw.xor"
}

type I64AtomicRmwXor struct {
	MemArg MemArg
}

func (self *I64AtomicRmwXor) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwXor) String() string {
	return "i64.atomic.rmw.xor"
}

type I32AtomicRmw8XorU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8XorU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8XorU) String() string {
	return "i32.atomic.rmw8.xor_u"
}

type I32AtomicRmw16XorU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16XorU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16XorU) String() string {
	return "i32.atomic.rmw16.xor_u"
}

type I64AtomicRmw8XorU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8XorU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8XorU) String() string {
	return "i64.atomic.rmw8.xor_u"
}

type I64AtomicRmw16XorU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16XorU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16XorU) String() string {
	return "i64.atomic.rmw16.xor_u"
}

type I64AtomicRmw32XorU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32XorU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32XorU) String() string {
	return "i64.atomic.rmw32.xor_u"
}

type I32AtomicRmwXchg struct {
	MemArg MemArg
}

func (self *I32AtomicRmwXchg) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwXchg) String() string {
	return "i32.atomic.rmw.xchg"
}

type I64AtomicRmwXchg struct {
	MemArg MemArg
}

func (self *I64AtomicRmwXchg) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwXchg) String() string {
	return "i64.atomic.rmw.xchg"
}

type I32AtomicRmw8XchgU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8XchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8XchgU) String() string {
	return "i32.atomic.rmw8.xchg_u"
}

type I32AtomicRmw16XchgU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16XchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16XchgU) String() string {
	return "i32.atomic.rmw16.xchg_u"
}

type I64AtomicRmw8XchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8XchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8XchgU) String() string {
	return "i64.atomic.rmw8.xchg_u"
}

type I64AtomicRmw16XchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16XchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16XchgU) String() string {
	return "i64.atomic.rmw16.xchg_u"
}

type I64AtomicRmw32XchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32XchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32XchgU) String() string {
	return "i64.atomic.rmw32.xchg_u"
}

type I32AtomicRmwCmpxchg struct {
	MemArg MemArg
}

func (self *I32AtomicRmwCmpxchg) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmwCmpxchg) String() string {
	return "i32.atomic.rmw.cmpxchg"
}

type I64AtomicRmwCmpxchg struct {
	MemArg MemArg
}

func (self *I64AtomicRmwCmpxchg) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmwCmpxchg) String() string {
	return "i64.atomic.rmw.cmpxchg"
}

type I32AtomicRmw8CmpxchgU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw8CmpxchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw8CmpxchgU) String() string {
	return "i32.atomic.rmw8.cmpxchg_u"
}

type I32AtomicRmw16CmpxchgU struct {
	MemArg MemArg
}

func (self *I32AtomicRmw16CmpxchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32AtomicRmw16CmpxchgU) String() string {
	return "i32.atomic.rmw16.cmpxchg_u"
}

type I64AtomicRmw8CmpxchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw8CmpxchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw8CmpxchgU) String() string {
	return "i64.atomic.rmw8.cmpxchg_u"
}

type I64AtomicRmw16CmpxchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw16CmpxchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw16CmpxchgU) String() string {
	return "i64.atomic.rmw16.cmpxchg_u"
}

type I64AtomicRmw32CmpxchgU struct {
	MemArg MemArg
}

func (self *I64AtomicRmw32CmpxchgU) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64AtomicRmw32CmpxchgU) String() string {
	return "i64.atomic.rmw32.cmpxchg_u"
}

type V128Load struct {
	MemArg MemArg
}

func (self *V128Load) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 16)
	if err != nil {
		return err
	}

	return nil
}

func (self *V128Load) String() string {
	return "v128.load"
}

type V128Store struct {
	MemArg MemArg
}

func (self *V128Store) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 16)
	if err != nil {
		return err
	}

	return nil
}

func (self *V128Store) String() string {
	return "v128.store"
}

type I8x16Eq struct {
}

func (self *I8x16Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Eq) String() string {
	return "i8x16.eq"
}

type I8x16Ne struct {
}

func (self *I8x16Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Ne) String() string {
	return "i8x16.ne"
}

type I8x16LtS struct {
}

func (self *I8x16LtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16LtS) String() string {
	return "i8x16.lt_s"
}

type I8x16LtU struct {
}

func (self *I8x16LtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16LtU) String() string {
	return "i8x16.lt_u"
}

type I8x16GtS struct {
}

func (self *I8x16GtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16GtS) String() string {
	return "i8x16.gt_s"
}

type I8x16GtU struct {
}

func (self *I8x16GtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16GtU) String() string {
	return "i8x16.gt_u"
}

type I8x16LeS struct {
}

func (self *I8x16LeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16LeS) String() string {
	return "i8x16.le_s"
}

type I8x16LeU struct {
}

func (self *I8x16LeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16LeU) String() string {
	return "i8x16.le_u"
}

type I8x16GeS struct {
}

func (self *I8x16GeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16GeS) String() string {
	return "i8x16.ge_s"
}

type I8x16GeU struct {
}

func (self *I8x16GeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16GeU) String() string {
	return "i8x16.ge_u"
}

type I16x8Eq struct {
}

func (self *I16x8Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Eq) String() string {
	return "i16x8.eq"
}

type I16x8Ne struct {
}

func (self *I16x8Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Ne) String() string {
	return "i16x8.ne"
}

type I16x8LtS struct {
}

func (self *I16x8LtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8LtS) String() string {
	return "i16x8.lt_s"
}

type I16x8LtU struct {
}

func (self *I16x8LtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8LtU) String() string {
	return "i16x8.lt_u"
}

type I16x8GtS struct {
}

func (self *I16x8GtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8GtS) String() string {
	return "i16x8.gt_s"
}

type I16x8GtU struct {
}

func (self *I16x8GtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8GtU) String() string {
	return "i16x8.gt_u"
}

type I16x8LeS struct {
}

func (self *I16x8LeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8LeS) String() string {
	return "i16x8.le_s"
}

type I16x8LeU struct {
}

func (self *I16x8LeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8LeU) String() string {
	return "i16x8.le_u"
}

type I16x8GeS struct {
}

func (self *I16x8GeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8GeS) String() string {
	return "i16x8.ge_s"
}

type I16x8GeU struct {
}

func (self *I16x8GeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8GeU) String() string {
	return "i16x8.ge_u"
}

type I32x4Eq struct {
}

func (self *I32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Eq) String() string {
	return "i32x4.eq"
}

type I32x4Ne struct {
}

func (self *I32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Ne) String() string {
	return "i32x4.ne"
}

type I32x4LtS struct {
}

func (self *I32x4LtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4LtS) String() string {
	return "i32x4.lt_s"
}

type I32x4LtU struct {
}

func (self *I32x4LtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4LtU) String() string {
	return "i32x4.lt_u"
}

type I32x4GtS struct {
}

func (self *I32x4GtS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4GtS) String() string {
	return "i32x4.gt_s"
}

type I32x4GtU struct {
}

func (self *I32x4GtU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4GtU) String() string {
	return "i32x4.gt_u"
}

type I32x4LeS struct {
}

func (self *I32x4LeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4LeS) String() string {
	return "i32x4.le_s"
}

type I32x4LeU struct {
}

func (self *I32x4LeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4LeU) String() string {
	return "i32x4.le_u"
}

type I32x4GeS struct {
}

func (self *I32x4GeS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4GeS) String() string {
	return "i32x4.ge_s"
}

type I32x4GeU struct {
}

func (self *I32x4GeU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4GeU) String() string {
	return "i32x4.ge_u"
}

type F32x4Eq struct {
}

func (self *F32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Eq) String() string {
	return "f32x4.eq"
}

type F32x4Ne struct {
}

func (self *F32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Ne) String() string {
	return "f32x4.ne"
}

type F32x4Lt struct {
}

func (self *F32x4Lt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Lt) String() string {
	return "f32x4.lt"
}

type F32x4Gt struct {
}

func (self *F32x4Gt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Gt) String() string {
	return "f32x4.gt"
}

type F32x4Le struct {
}

func (self *F32x4Le) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Le) String() string {
	return "f32x4.le"
}

type F32x4Ge struct {
}

func (self *F32x4Ge) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Ge) String() string {
	return "f32x4.ge"
}

type F64x2Eq struct {
}

func (self *F64x2Eq) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Eq) String() string {
	return "f64x2.eq"
}

type F64x2Ne struct {
}

func (self *F64x2Ne) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Ne) String() string {
	return "f64x2.ne"
}

type F64x2Lt struct {
}

func (self *F64x2Lt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Lt) String() string {
	return "f64x2.lt"
}

type F64x2Gt struct {
}

func (self *F64x2Gt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Gt) String() string {
	return "f64x2.gt"
}

type F64x2Le struct {
}

func (self *F64x2Le) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Le) String() string {
	return "f64x2.le"
}

type F64x2Ge struct {
}

func (self *F64x2Ge) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Ge) String() string {
	return "f64x2.ge"
}

type V128Not struct {
}

func (self *V128Not) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128Not) String() string {
	return "v128.not"
}

type V128And struct {
}

func (self *V128And) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128And) String() string {
	return "v128.and"
}

type V128Or struct {
}

func (self *V128Or) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128Or) String() string {
	return "v128.or"
}

type V128Xor struct {
}

func (self *V128Xor) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128Xor) String() string {
	return "v128.xor"
}

type V128Bitselect struct {
}

func (self *V128Bitselect) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128Bitselect) String() string {
	return "v128.bitselect"
}

type I8x16Neg struct {
}

func (self *I8x16Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Neg) String() string {
	return "i8x16.neg"
}

type I8x16AnyTrue struct {
}

func (self *I8x16AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16AnyTrue) String() string {
	return "i8x16.any_true"
}

type I8x16AllTrue struct {
}

func (self *I8x16AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16AllTrue) String() string {
	return "i8x16.all_true"
}

type I8x16Shl struct {
}

func (self *I8x16Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Shl) String() string {
	return "i8x16.shl"
}

type I8x16ShrS struct {
}

func (self *I8x16ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16ShrS) String() string {
	return "i8x16.shr_s"
}

type I8x16ShrU struct {
}

func (self *I8x16ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16ShrU) String() string {
	return "i8x16.shr_u"
}

type I8x16Add struct {
}

func (self *I8x16Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Add) String() string {
	return "i8x16.add"
}

type I8x16AddSaturateS struct {
}

func (self *I8x16AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16AddSaturateS) String() string {
	return "i8x16.add_saturate_s"
}

type I8x16AddSaturateU struct {
}

func (self *I8x16AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16AddSaturateU) String() string {
	return "i8x16.add_saturate_u"
}

type I8x16Sub struct {
}

func (self *I8x16Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Sub) String() string {
	return "i8x16.sub"
}

type I8x16SubSaturateS struct {
}

func (self *I8x16SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16SubSaturateS) String() string {
	return "i8x16.sub_saturate_s"
}

type I8x16SubSaturateU struct {
}

func (self *I8x16SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16SubSaturateU) String() string {
	return "i8x16.sub_saturate_u"
}

type I8x16Mul struct {
}

func (self *I8x16Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16Mul) String() string {
	return "i8x16.mul"
}

type I16x8Neg struct {
}

func (self *I16x8Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Neg) String() string {
	return "i16x8.neg"
}

type I16x8AnyTrue struct {
}

func (self *I16x8AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8AnyTrue) String() string {
	return "i16x8.any_true"
}

type I16x8AllTrue struct {
}

func (self *I16x8AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8AllTrue) String() string {
	return "i16x8.all_true"
}

type I16x8Shl struct {
}

func (self *I16x8Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Shl) String() string {
	return "i16x8.shl"
}

type I16x8ShrS struct {
}

func (self *I16x8ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8ShrS) String() string {
	return "i16x8.shr_s"
}

type I16x8ShrU struct {
}

func (self *I16x8ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8ShrU) String() string {
	return "i16x8.shr_u"
}

type I16x8Add struct {
}

func (self *I16x8Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Add) String() string {
	return "i16x8.add"
}

type I16x8AddSaturateS struct {
}

func (self *I16x8AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8AddSaturateS) String() string {
	return "i16x8.add_saturate_s"
}

type I16x8AddSaturateU struct {
}

func (self *I16x8AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8AddSaturateU) String() string {
	return "i16x8.add_saturate_u"
}

type I16x8Sub struct {
}

func (self *I16x8Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Sub) String() string {
	return "i16x8.sub"
}

type I16x8SubSaturateS struct {
}

func (self *I16x8SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8SubSaturateS) String() string {
	return "i16x8.sub_saturate_s"
}

type I16x8SubSaturateU struct {
}

func (self *I16x8SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8SubSaturateU) String() string {
	return "i16x8.sub_saturate_u"
}

type I16x8Mul struct {
}

func (self *I16x8Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8Mul) String() string {
	return "i16x8.mul"
}

type I32x4Neg struct {
}

func (self *I32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Neg) String() string {
	return "i32x4.neg"
}

type I32x4AnyTrue struct {
}

func (self *I32x4AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4AnyTrue) String() string {
	return "i32x4.any_true"
}

type I32x4AllTrue struct {
}

func (self *I32x4AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4AllTrue) String() string {
	return "i32x4.all_true"
}

type I32x4Shl struct {
}

func (self *I32x4Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Shl) String() string {
	return "i32x4.shl"
}

type I32x4ShrS struct {
}

func (self *I32x4ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4ShrS) String() string {
	return "i32x4.shr_s"
}

type I32x4ShrU struct {
}

func (self *I32x4ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4ShrU) String() string {
	return "i32x4.shr_u"
}

type I32x4Add struct {
}

func (self *I32x4Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Add) String() string {
	return "i32x4.add"
}

type I32x4Sub struct {
}

func (self *I32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Sub) String() string {
	return "i32x4.sub"
}

type I32x4Mul struct {
}

func (self *I32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4Mul) String() string {
	return "i32x4.mul"
}

type I64x2Neg struct {
}

func (self *I64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2Neg) String() string {
	return "i64x2.neg"
}

type I64x2AnyTrue struct {
}

func (self *I64x2AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2AnyTrue) String() string {
	return "i64x2.any_true"
}

type I64x2AllTrue struct {
}

func (self *I64x2AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2AllTrue) String() string {
	return "i64x2.all_true"
}

type I64x2Shl struct {
}

func (self *I64x2Shl) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2Shl) String() string {
	return "i64x2.shl"
}

type I64x2ShrS struct {
}

func (self *I64x2ShrS) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2ShrS) String() string {
	return "i64x2.shr_s"
}

type I64x2ShrU struct {
}

func (self *I64x2ShrU) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2ShrU) String() string {
	return "i64x2.shr_u"
}

type I64x2Add struct {
}

func (self *I64x2Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2Add) String() string {
	return "i64x2.add"
}

type I64x2Sub struct {
}

func (self *I64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2Sub) String() string {
	return "i64x2.sub"
}

type I64x2Mul struct {
}

func (self *I64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2Mul) String() string {
	return "i64x2.mul"
}

type F32x4Abs struct {
}

func (self *F32x4Abs) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Abs) String() string {
	return "f32x4.abs"
}

type F32x4Neg struct {
}

func (self *F32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Neg) String() string {
	return "f32x4.neg"
}

type F32x4Sqrt struct {
}

func (self *F32x4Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Sqrt) String() string {
	return "f32x4.sqrt"
}

type F32x4Add struct {
}

func (self *F32x4Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Add) String() string {
	return "f32x4.add"
}

type F32x4Sub struct {
}

func (self *F32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Sub) String() string {
	return "f32x4.sub"
}

type F32x4Mul struct {
}

func (self *F32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Mul) String() string {
	return "f32x4.mul"
}

type F32x4Div struct {
}

func (self *F32x4Div) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Div) String() string {
	return "f32x4.div"
}

type F32x4Min struct {
}

func (self *F32x4Min) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Min) String() string {
	return "f32x4.min"
}

type F32x4Max struct {
}

func (self *F32x4Max) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4Max) String() string {
	return "f32x4.max"
}

type F64x2Abs struct {
}

func (self *F64x2Abs) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Abs) String() string {
	return "f64x2.abs"
}

type F64x2Neg struct {
}

func (self *F64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Neg) String() string {
	return "f64x2.neg"
}

type F64x2Sqrt struct {
}

func (self *F64x2Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Sqrt) String() string {
	return "f64x2.sqrt"
}

type F64x2Add struct {
}

func (self *F64x2Add) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Add) String() string {
	return "f64x2.add"
}

type F64x2Sub struct {
}

func (self *F64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Sub) String() string {
	return "f64x2.sub"
}

type F64x2Mul struct {
}

func (self *F64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Mul) String() string {
	return "f64x2.mul"
}

type F64x2Div struct {
}

func (self *F64x2Div) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Div) String() string {
	return "f64x2.div"
}

type F64x2Min struct {
}

func (self *F64x2Min) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Min) String() string {
	return "f64x2.min"
}

type F64x2Max struct {
}

func (self *F64x2Max) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2Max) String() string {
	return "f64x2.max"
}

type I32x4TruncSatF32x4S struct {
}

func (self *I32x4TruncSatF32x4S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4TruncSatF32x4S) String() string {
	return "i32x4.trunc_sat_f32x4_s"
}

type I32x4TruncSatF32x4U struct {
}

func (self *I32x4TruncSatF32x4U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4TruncSatF32x4U) String() string {
	return "i32x4.trunc_sat_f32x4_u"
}

type I64x2TruncSatF64x2S struct {
}

func (self *I64x2TruncSatF64x2S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2TruncSatF64x2S) String() string {
	return "i64x2.trunc_sat_f64x2_s"
}

type I64x2TruncSatF64x2U struct {
}

func (self *I64x2TruncSatF64x2U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I64x2TruncSatF64x2U) String() string {
	return "i64x2.trunc_sat_f64x2_u"
}

type F32x4ConvertI32x4S struct {
}

func (self *F32x4ConvertI32x4S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4ConvertI32x4S) String() string {
	return "f32x4.convert_i32x4_s"
}

type F32x4ConvertI32x4U struct {
}

func (self *F32x4ConvertI32x4U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F32x4ConvertI32x4U) String() string {
	return "f32x4.convert_i32x4_u"
}

type F64x2ConvertI64x2S struct {
}

func (self *F64x2ConvertI64x2S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2ConvertI64x2S) String() string {
	return "f64x2.convert_i64x2_s"
}

type F64x2ConvertI64x2U struct {
}

func (self *F64x2ConvertI64x2U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *F64x2ConvertI64x2U) String() string {
	return "f64x2.convert_i64x2_u"
}

type V8x16Swizzle struct {
}

func (self *V8x16Swizzle) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V8x16Swizzle) String() string {
	return "v8x16.swizzle"
}

type V8x16LoadSplat struct {
	MemArg MemArg
}

func (self *V8x16LoadSplat) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *V8x16LoadSplat) String() string {
	return "v8x16.load_splat"
}

type V16x8LoadSplat struct {
	MemArg MemArg
}

func (self *V16x8LoadSplat) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *V16x8LoadSplat) String() string {
	return "v16x8.load_splat"
}

type V32x4LoadSplat struct {
	MemArg MemArg
}

func (self *V32x4LoadSplat) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *V32x4LoadSplat) String() string {
	return "v32x4.load_splat"
}

type V64x2LoadSplat struct {
	MemArg MemArg
}

func (self *V64x2LoadSplat) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 8)
	if err != nil {
		return err
	}

	return nil
}

func (self *V64x2LoadSplat) String() string {
	return "v64x2.load_splat"
}

type I8x16NarrowI16x8S struct {
}

func (self *I8x16NarrowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16NarrowI16x8S) String() string {
	return "i8x16.narrow_i16x8_s"
}

type I8x16NarrowI16x8U struct {
}

func (self *I8x16NarrowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I8x16NarrowI16x8U) String() string {
	return "i8x16.narrow_i16x8_u"
}

type I16x8NarrowI32x4S struct {
}

func (self *I16x8NarrowI32x4S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8NarrowI32x4S) String() string {
	return "i16x8.narrow_i32x4_s"
}

type I16x8NarrowI32x4U struct {
}

func (self *I16x8NarrowI32x4U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8NarrowI32x4U) String() string {
	return "i16x8.narrow_i32x4_u"
}

type I16x8WidenLowI8x16S struct {
}

func (self *I16x8WidenLowI8x16S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8WidenLowI8x16S) String() string {
	return "i16x8.widen_low_i8x16_s"
}

type I16x8WidenHighI8x16S struct {
}

func (self *I16x8WidenHighI8x16S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8WidenHighI8x16S) String() string {
	return "i16x8.widen_high_i8x16_s"
}

type I16x8WidenLowI8x16U struct {
}

func (self *I16x8WidenLowI8x16U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8WidenLowI8x16U) String() string {
	return "i16x8.widen_low_i8x16_u"
}

type I16x8WidenHighI8x16u struct {
}

func (self *I16x8WidenHighI8x16u) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I16x8WidenHighI8x16u) String() string {
	return "i16x8.widen_high_i8x16_u"
}

type I32x4WidenLowI16x8S struct {
}

func (self *I32x4WidenLowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4WidenLowI16x8S) String() string {
	return "i32x4.widen_low_i16x8_s"
}

type I32x4WidenHighI16x8S struct {
}

func (self *I32x4WidenHighI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4WidenHighI16x8S) String() string {
	return "i32x4.widen_high_i16x8_s"
}

type I32x4WidenLowI16x8U struct {
}

func (self *I32x4WidenLowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4WidenLowI16x8U) String() string {
	return "i32x4.widen_low_i16x8_u"
}

type I32x4WidenHighI16x8u struct {
}

func (self *I32x4WidenHighI16x8u) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *I32x4WidenHighI16x8u) String() string {
	return "i32x4.widen_high_i16x8_u"
}

type I16x8Load8x8S struct {
	MemArg MemArg
}

func (self *I16x8Load8x8S) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I16x8Load8x8S) String() string {
	return "i16x8.load8x8_s"
}

type I16x8Load8x8U struct {
	MemArg MemArg
}

func (self *I16x8Load8x8U) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 1)
	if err != nil {
		return err
	}

	return nil
}

func (self *I16x8Load8x8U) String() string {
	return "i16x8.load8x8_u"
}

type I32x4Load16x4S struct {
	MemArg MemArg
}

func (self *I32x4Load16x4S) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32x4Load16x4S) String() string {
	return "i32x4.load16x4_s"
}

type I32x4Load16x4U struct {
	MemArg MemArg
}

func (self *I32x4Load16x4U) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 2)
	if err != nil {
		return err
	}

	return nil
}

func (self *I32x4Load16x4U) String() string {
	return "i32x4.load16x4_u"
}

type I64x2Load32x2S struct {
	MemArg MemArg
}

func (self *I64x2Load32x2S) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64x2Load32x2S) String() string {
	return "i64x2.load32x2_s"
}

type I64x2Load32x2U struct {
	MemArg MemArg
}

func (self *I64x2Load32x2U) parseInstrBody(ps *parser.ParserBuffer) error {
	err := self.MemArg.Parse(ps, 4)
	if err != nil {
		return err
	}

	return nil
}

func (self *I64x2Load32x2U) String() string {
	return "i64x2.load32x2_u"
}

type V128Andnot struct {
}

func (self *V128Andnot) parseInstrBody(ps *parser.ParserBuffer) error {

	return nil
}

func (self *V128Andnot) String() string {
	return "v128.andnot"
}

func parseInstr(ps *parser.ParserBuffer) (Instruction, error) {
	var inst Instruction
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	case "block":
		inst = &Block{}
	case "if":
		inst = &If{}
	case "else":
		inst = &Else{}
	case "loop":
		inst = &Loop{}
	case "end":
		inst = &End{}
	case "unreachable":
		inst = &Unreachable{}
	case "nop":
		inst = &Nop{}
	case "br":
		inst = &Br{}
	case "br_if":
		inst = &BrIf{}
	case "br_table":
		inst = &BrTable{}
	case "return":
		inst = &Return{}
	case "call":
		inst = &Call{}
	case "call_indirect":
		inst = &CallIndirect{}
	case "return_call":
		inst = &ReturnCall{}
	case "return_call_indirect":
		inst = &ReturnCallIndirect{}
	case "drop":
		inst = &Drop{}
	case "select":
		inst = &Select{}
	case "local.get", "get_local":
		inst = &LocalGet{}
	case "local.set", "set_local":
		inst = &LocalSet{}
	case "local.tee", "tee_local":
		inst = &LocalTee{}
	case "global.get", "get_global":
		inst = &GlobalGet{}
	case "global.set", "set_global":
		inst = &GlobalSet{}
	case "table.get":
		inst = &TableGet{}
	case "table.set":
		inst = &TableSet{}
	case "i32.load":
		inst = &I32Load{}
	case "i64.load":
		inst = &I64Load{}
	case "f32.load":
		inst = &F32Load{}
	case "f64.load":
		inst = &F64Load{}
	case "i32.load8_s":
		inst = &I32Load8s{}
	case "i32.load8_u":
		inst = &I32Load8u{}
	case "i32.load16_s":
		inst = &I32Load16s{}
	case "i32.load16_u":
		inst = &I32Load16u{}
	case "i64.load8_s":
		inst = &I64Load8s{}
	case "i64.load8_u":
		inst = &I64Load8u{}
	case "i64.load16_s":
		inst = &I64Load16s{}
	case "i64.load16_u":
		inst = &I64Load16u{}
	case "i64.load32_s":
		inst = &I64Load32s{}
	case "i64.load32_u":
		inst = &I64Load32u{}
	case "i32.store":
		inst = &I32Store{}
	case "i64.store":
		inst = &I64Store{}
	case "f32.store":
		inst = &F32Store{}
	case "f64.store":
		inst = &F64Store{}
	case "i32.store8":
		inst = &I32Store8{}
	case "i32.store16":
		inst = &I32Store16{}
	case "i64.store8":
		inst = &I64Store8{}
	case "i64.store16":
		inst = &I64Store16{}
	case "i64.store32":
		inst = &I64Store32{}
	case "memory.size", "current_memory":
		inst = &MemorySize{}
	case "memory.grow", "grow_memory":
		inst = &MemoryGrow{}
	case "memory.copy":
		inst = &MemoryCopy{}
	case "memory.fill":
		inst = &MemoryFill{}
	case "data.drop":
		inst = &DataDrop{}
	case "elem.drop":
		inst = &ElemDrop{}
	case "table.copy":
		inst = &TableCopy{}
	case "table.fill":
		inst = &TableFill{}
	case "table.size":
		inst = &TableSize{}
	case "table.grow":
		inst = &TableGrow{}
	case "ref.null":
		inst = &RefNull{}
	case "ref.is_null":
		inst = &RefIsNull{}
	case "ref.host":
		inst = &RefHost{}
	case "ref.func":
		inst = &RefFunc{}
	case "i32.const":
		inst = &I32Const{}
	case "i64.const":
		inst = &I64Const{}
	case "f32.const":
		inst = &F32Const{}
	case "f64.const":
		inst = &F64Const{}
	case "i32.clz":
		inst = &I32Clz{}
	case "i32.ctz":
		inst = &I32Ctz{}
	case "i32.popcnt":
		inst = &I32Pocnt{}
	case "i32.add":
		inst = &I32Add{}
	case "i32.sub":
		inst = &I32Sub{}
	case "i32.mul":
		inst = &I32Mul{}
	case "i32.div_s":
		inst = &I32DivS{}
	case "i32.div_u":
		inst = &I32DivU{}
	case "i32.rem_s":
		inst = &I32RemS{}
	case "i32.rem_u":
		inst = &I32RemU{}
	case "i32.and":
		inst = &I32And{}
	case "i32.or":
		inst = &I32Or{}
	case "i32.xor":
		inst = &I32Xor{}
	case "i32.shl":
		inst = &I32Shl{}
	case "i32.shr_s":
		inst = &I32ShrS{}
	case "i32.shr_u":
		inst = &I32ShrU{}
	case "i32.rotl":
		inst = &I32Rotl{}
	case "i32.rotr":
		inst = &I32Rotr{}
	case "i64.clz":
		inst = &I64Clz{}
	case "i64.ctz":
		inst = &I64Ctz{}
	case "i64.popcnt":
		inst = &I64Popcnt{}
	case "i64.add":
		inst = &I64Add{}
	case "i64.sub":
		inst = &I64Sub{}
	case "i64.mul":
		inst = &I64Mul{}
	case "i64.div_s":
		inst = &I64DivS{}
	case "i64.div_u":
		inst = &I64DivU{}
	case "i64.rem_s":
		inst = &I64RemS{}
	case "i64.rem_u":
		inst = &I64RemU{}
	case "i64.and":
		inst = &I64And{}
	case "i64.or":
		inst = &I64Or{}
	case "i64.xor":
		inst = &I64Xor{}
	case "i64.shl":
		inst = &I64Shl{}
	case "i64.shr_s":
		inst = &I64ShrS{}
	case "i64.shr_u":
		inst = &I64ShrU{}
	case "i64.rotl":
		inst = &I64Rotl{}
	case "i64.rotr":
		inst = &I64Rotr{}
	case "f32.abs":
		inst = &F32Abs{}
	case "f32.neg":
		inst = &F32Neg{}
	case "f32.ceil":
		inst = &F32Ceil{}
	case "f32.floor":
		inst = &F32Floor{}
	case "f32.trunc":
		inst = &F32Trunc{}
	case "f32.nearest":
		inst = &F32Nearest{}
	case "f32.sqrt":
		inst = &F32Sqrt{}
	case "f32.add":
		inst = &F32Add{}
	case "f32.sub":
		inst = &F32Sub{}
	case "f32.mul":
		inst = &F32Mul{}
	case "f32.div":
		inst = &F32Div{}
	case "f32.min":
		inst = &F32Min{}
	case "f32.max":
		inst = &F32Max{}
	case "f32.copysign":
		inst = &F32Copysign{}
	case "f64.abs":
		inst = &F64Abs{}
	case "f64.neg":
		inst = &F64Neg{}
	case "f64.ceil":
		inst = &F64Ceil{}
	case "f64.floor":
		inst = &F64Floor{}
	case "f64.trunc":
		inst = &F64Trunc{}
	case "f64.nearest":
		inst = &F64Nearest{}
	case "f64.sqrt":
		inst = &F64Sqrt{}
	case "f64.add":
		inst = &F64Add{}
	case "f64.sub":
		inst = &F64Sub{}
	case "f64.mul":
		inst = &F64Mul{}
	case "f64.div":
		inst = &F64Div{}
	case "f64.min":
		inst = &F64Min{}
	case "f64.max":
		inst = &F64Max{}
	case "f64.copysign":
		inst = &F64Copysign{}
	case "i32.eqz":
		inst = &I32Eqz{}
	case "i32.eq":
		inst = &I32Eq{}
	case "i32.ne":
		inst = &I32Ne{}
	case "i32.lt_s":
		inst = &I32LtS{}
	case "i32.lt_u":
		inst = &I32LtU{}
	case "i32.gt_s":
		inst = &I32GtS{}
	case "i32.gt_u":
		inst = &I32GtU{}
	case "i32.le_s":
		inst = &I32LeS{}
	case "i32.le_u":
		inst = &I32LeU{}
	case "i32.ge_s":
		inst = &I32GeS{}
	case "i32.ge_u":
		inst = &I32GeU{}
	case "i64.eqz":
		inst = &I64Eqz{}
	case "i64.eq":
		inst = &I64Eq{}
	case "i64.ne":
		inst = &I64Ne{}
	case "i64.lt_s":
		inst = &I64LtS{}
	case "i64.lt_u":
		inst = &I64LtU{}
	case "i64.gt_s":
		inst = &I64GtS{}
	case "i64.gt_u":
		inst = &I64GtU{}
	case "i64.le_s":
		inst = &I64LeS{}
	case "i64.le_u":
		inst = &I64LeU{}
	case "i64.ge_s":
		inst = &I64GeS{}
	case "i64.ge_u":
		inst = &I64GeU{}
	case "f32.eq":
		inst = &F32Eq{}
	case "f32.ne":
		inst = &F32Ne{}
	case "f32.lt":
		inst = &F32Lt{}
	case "f32.gt":
		inst = &F32Gt{}
	case "f32.le":
		inst = &F32Le{}
	case "f32.ge":
		inst = &F32Ge{}
	case "f64.eq":
		inst = &F64Eq{}
	case "f64.ne":
		inst = &F64Ne{}
	case "f64.lt":
		inst = &F64Lt{}
	case "f64.gt":
		inst = &F64Gt{}
	case "f64.le":
		inst = &F64Le{}
	case "f64.ge":
		inst = &F64Ge{}
	case "i32.wrap_i64", "i32.wrap/i64":
		inst = &I32WrapI64{}
	case "i32.trunc_f32_s", "i32.trunc_s/f32":
		inst = &I32TruncF32S{}
	case "i32.trunc_f32_u", "i32.trunc_u/f32":
		inst = &I32TruncF32U{}
	case "i32.trunc_f64_s", "i32.trunc_s/f64":
		inst = &I32TruncF64S{}
	case "i32.trunc_f64_u", "i32.trunc_u/f64":
		inst = &I32TruncF64U{}
	case "i64.extend_i32_s", "i64.extend_s/i32":
		inst = &I64ExtendI32S{}
	case "i64.extend_i32_u", "i64.extend_u/i32":
		inst = &I64ExtendI32U{}
	case "i64.trunc_f32_s", "i64.trunc_s/f32":
		inst = &I64TruncF32S{}
	case "i64.trunc_f32_u", "i64.trunc_u/f32":
		inst = &I64TruncF32U{}
	case "i64.trunc_f64_s", "i64.trunc_s/f64":
		inst = &I64TruncF64S{}
	case "i64.trunc_f64_u", "i64.trunc_u/f64":
		inst = &I64TruncF64U{}
	case "f32.convert_i32_s", "f32.convert_s/i32":
		inst = &F32ConvertI32S{}
	case "f32.convert_i32_u", "f32.convert_u/i32":
		inst = &F32ConvertI32U{}
	case "f32.convert_i64_s", "f32.convert_s/i64":
		inst = &F32ConvertI64S{}
	case "f32.convert_i64_u", "f32.convert_u/i64":
		inst = &F32ConvertI64U{}
	case "f32.demote_f64", "f32.demote/f64":
		inst = &F32DemoteF64{}
	case "f64.convert_i32_s", "f64.convert_s/i32":
		inst = &F64ConvertI32S{}
	case "f64.convert_i32_u", "f64.convert_u/i32":
		inst = &F64ConvertI32U{}
	case "f64.convert_i64_s", "f64.convert_s/i64":
		inst = &F64ConvertI64S{}
	case "f64.convert_i64_u", "f64.convert_u/i64":
		inst = &F64ConvertI64U{}
	case "f64.promote_f32", "f64.promote/f32":
		inst = &F64PromoteF32{}
	case "i32.reinterpret_f32", "i32.reinterpret/f32":
		inst = &I32ReinterpretF32{}
	case "i64.reinterpret_f64", "i64.reinterpret/f64":
		inst = &I64ReinterpretF64{}
	case "f32.reinterpret_i32", "f32.reinterpret/i32":
		inst = &F32ReinterpretI32{}
	case "f64.reinterpret_i64", "f64.reinterpret/i64":
		inst = &F64ReinterpretI64{}
	case "i32.trunc_sat_f32_s", "i32.trunc_s:sat/f32":
		inst = &I32TruncSatF32S{}
	case "i32.trunc_sat_f32_u", "i32.trunc_u:sat/f32":
		inst = &I32TruncSatF32U{}
	case "i32.trunc_sat_f64_s", "i32.trunc_s:sat/f64":
		inst = &I32TruncSatF64S{}
	case "i32.trunc_sat_f64_u", "i32.trunc_u:sat/f64":
		inst = &I32TruncSatF64U{}
	case "i64.trunc_sat_f32_s", "i64.trunc_s:sat/f32":
		inst = &I64TruncSatF32S{}
	case "i64.trunc_sat_f32_u", "i64.trunc_u:sat/f32":
		inst = &I64TruncSatF32U{}
	case "i64.trunc_sat_f64_s", "i64.trunc_s:sat/f64":
		inst = &I64TruncSatF64S{}
	case "i64.trunc_sat_f64_u", "i64.trunc_u:sat/f64":
		inst = &I64TruncSatF64U{}
	case "i32.extend8_s":
		inst = &I32Extend8S{}
	case "i32.extend16_s":
		inst = &I32Extend16S{}
	case "i64.extend8_s":
		inst = &I64Extend8S{}
	case "i64.extend16_s":
		inst = &I64Extend16S{}
	case "i64.extend32_s":
		inst = &I64Extend32S{}
	case "atomic.notify":
		inst = &AtomicNotify{}
	case "i32.atomic.wait":
		inst = &I32AtomicWait{}
	case "i64.atomic.wait":
		inst = &I64AtomicWait{}
	case "atomic.fence":
		inst = &AtomicFence{}
	case "i32.atomic.load":
		inst = &I32AtomicLoad{}
	case "i64.atomic.load":
		inst = &I64AtomicLoad{}
	case "i32.atomic.load8_u":
		inst = &I32AtomicLoad8u{}
	case "i32.atomic.load16_u":
		inst = &I32AtomicLoad16u{}
	case "i64.atomic.load8_u":
		inst = &I64AtomicLoad8u{}
	case "i64.atomic.load16_u":
		inst = &I64AtomicLoad16u{}
	case "i64.atomic.load32_u":
		inst = &I64AtomicLoad32u{}
	case "i32.atomic.store":
		inst = &I32AtomicStore{}
	case "i64.atomic.store":
		inst = &I64AtomicStore{}
	case "i32.atomic.store8":
		inst = &I32AtomicStore8{}
	case "i32.atomic.store16":
		inst = &I32AtomicStore16{}
	case "i64.atomic.store8":
		inst = &I64AtomicStore8{}
	case "i64.atomic.store16":
		inst = &I64AtomicStore16{}
	case "i64.atomic.store32":
		inst = &I64AtomicStore32{}
	case "i32.atomic.rmw.add":
		inst = &I32AtomicRmwAdd{}
	case "i64.atomic.rmw.add":
		inst = &I64AtomicRmwAdd{}
	case "i32.atomic.rmw8.add_u":
		inst = &I32AtomicRmw8AddU{}
	case "i32.atomic.rmw16.add_u":
		inst = &I32AtomicRmw16AddU{}
	case "i64.atomic.rmw8.add_u":
		inst = &I64AtomicRmw8AddU{}
	case "i64.atomic.rmw16.add_u":
		inst = &I64AtomicRmw16AddU{}
	case "i64.atomic.rmw32.add_u":
		inst = &I64AtomicRmw32AddU{}
	case "i32.atomic.rmw.sub":
		inst = &I32AtomicRmwSub{}
	case "i64.atomic.rmw.sub":
		inst = &I64AtomicRmwSub{}
	case "i32.atomic.rmw8.sub_u":
		inst = &I32AtomicRmw8SubU{}
	case "i32.atomic.rmw16.sub_u":
		inst = &I32AtomicRmw16SubU{}
	case "i64.atomic.rmw8.sub_u":
		inst = &I64AtomicRmw8SubU{}
	case "i64.atomic.rmw16.sub_u":
		inst = &I64AtomicRmw16SubU{}
	case "i64.atomic.rmw32.sub_u":
		inst = &I64AtomicRmw32SubU{}
	case "i32.atomic.rmw.and":
		inst = &I32AtomicRmwAnd{}
	case "i64.atomic.rmw.and":
		inst = &I64AtomicRmwAnd{}
	case "i32.atomic.rmw8.and_u":
		inst = &I32AtomicRmw8AndU{}
	case "i32.atomic.rmw16.and_u":
		inst = &I32AtomicRmw16AndU{}
	case "i64.atomic.rmw8.and_u":
		inst = &I64AtomicRmw8AndU{}
	case "i64.atomic.rmw16.and_u":
		inst = &I64AtomicRmw16AndU{}
	case "i64.atomic.rmw32.and_u":
		inst = &I64AtomicRmw32AndU{}
	case "i32.atomic.rmw.or":
		inst = &I32AtomicRmwOr{}
	case "i64.atomic.rmw.or":
		inst = &I64AtomicRmwOr{}
	case "i32.atomic.rmw8.or_u":
		inst = &I32AtomicRmw8OrU{}
	case "i32.atomic.rmw16.or_u":
		inst = &I32AtomicRmw16OrU{}
	case "i64.atomic.rmw8.or_u":
		inst = &I64AtomicRmw8OrU{}
	case "i64.atomic.rmw16.or_u":
		inst = &I64AtomicRmw16OrU{}
	case "i64.atomic.rmw32.or_u":
		inst = &I64AtomicRmw32OrU{}
	case "i32.atomic.rmw.xor":
		inst = &I32AtomicRmwXor{}
	case "i64.atomic.rmw.xor":
		inst = &I64AtomicRmwXor{}
	case "i32.atomic.rmw8.xor_u":
		inst = &I32AtomicRmw8XorU{}
	case "i32.atomic.rmw16.xor_u":
		inst = &I32AtomicRmw16XorU{}
	case "i64.atomic.rmw8.xor_u":
		inst = &I64AtomicRmw8XorU{}
	case "i64.atomic.rmw16.xor_u":
		inst = &I64AtomicRmw16XorU{}
	case "i64.atomic.rmw32.xor_u":
		inst = &I64AtomicRmw32XorU{}
	case "i32.atomic.rmw.xchg":
		inst = &I32AtomicRmwXchg{}
	case "i64.atomic.rmw.xchg":
		inst = &I64AtomicRmwXchg{}
	case "i32.atomic.rmw8.xchg_u":
		inst = &I32AtomicRmw8XchgU{}
	case "i32.atomic.rmw16.xchg_u":
		inst = &I32AtomicRmw16XchgU{}
	case "i64.atomic.rmw8.xchg_u":
		inst = &I64AtomicRmw8XchgU{}
	case "i64.atomic.rmw16.xchg_u":
		inst = &I64AtomicRmw16XchgU{}
	case "i64.atomic.rmw32.xchg_u":
		inst = &I64AtomicRmw32XchgU{}
	case "i32.atomic.rmw.cmpxchg":
		inst = &I32AtomicRmwCmpxchg{}
	case "i64.atomic.rmw.cmpxchg":
		inst = &I64AtomicRmwCmpxchg{}
	case "i32.atomic.rmw8.cmpxchg_u":
		inst = &I32AtomicRmw8CmpxchgU{}
	case "i32.atomic.rmw16.cmpxchg_u":
		inst = &I32AtomicRmw16CmpxchgU{}
	case "i64.atomic.rmw8.cmpxchg_u":
		inst = &I64AtomicRmw8CmpxchgU{}
	case "i64.atomic.rmw16.cmpxchg_u":
		inst = &I64AtomicRmw16CmpxchgU{}
	case "i64.atomic.rmw32.cmpxchg_u":
		inst = &I64AtomicRmw32CmpxchgU{}
	case "v128.load":
		inst = &V128Load{}
	case "v128.store":
		inst = &V128Store{}
	case "i8x16.eq":
		inst = &I8x16Eq{}
	case "i8x16.ne":
		inst = &I8x16Ne{}
	case "i8x16.lt_s":
		inst = &I8x16LtS{}
	case "i8x16.lt_u":
		inst = &I8x16LtU{}
	case "i8x16.gt_s":
		inst = &I8x16GtS{}
	case "i8x16.gt_u":
		inst = &I8x16GtU{}
	case "i8x16.le_s":
		inst = &I8x16LeS{}
	case "i8x16.le_u":
		inst = &I8x16LeU{}
	case "i8x16.ge_s":
		inst = &I8x16GeS{}
	case "i8x16.ge_u":
		inst = &I8x16GeU{}
	case "i16x8.eq":
		inst = &I16x8Eq{}
	case "i16x8.ne":
		inst = &I16x8Ne{}
	case "i16x8.lt_s":
		inst = &I16x8LtS{}
	case "i16x8.lt_u":
		inst = &I16x8LtU{}
	case "i16x8.gt_s":
		inst = &I16x8GtS{}
	case "i16x8.gt_u":
		inst = &I16x8GtU{}
	case "i16x8.le_s":
		inst = &I16x8LeS{}
	case "i16x8.le_u":
		inst = &I16x8LeU{}
	case "i16x8.ge_s":
		inst = &I16x8GeS{}
	case "i16x8.ge_u":
		inst = &I16x8GeU{}
	case "i32x4.eq":
		inst = &I32x4Eq{}
	case "i32x4.ne":
		inst = &I32x4Ne{}
	case "i32x4.lt_s":
		inst = &I32x4LtS{}
	case "i32x4.lt_u":
		inst = &I32x4LtU{}
	case "i32x4.gt_s":
		inst = &I32x4GtS{}
	case "i32x4.gt_u":
		inst = &I32x4GtU{}
	case "i32x4.le_s":
		inst = &I32x4LeS{}
	case "i32x4.le_u":
		inst = &I32x4LeU{}
	case "i32x4.ge_s":
		inst = &I32x4GeS{}
	case "i32x4.ge_u":
		inst = &I32x4GeU{}
	case "f32x4.eq":
		inst = &F32x4Eq{}
	case "f32x4.ne":
		inst = &F32x4Ne{}
	case "f32x4.lt":
		inst = &F32x4Lt{}
	case "f32x4.gt":
		inst = &F32x4Gt{}
	case "f32x4.le":
		inst = &F32x4Le{}
	case "f32x4.ge":
		inst = &F32x4Ge{}
	case "f64x2.eq":
		inst = &F64x2Eq{}
	case "f64x2.ne":
		inst = &F64x2Ne{}
	case "f64x2.lt":
		inst = &F64x2Lt{}
	case "f64x2.gt":
		inst = &F64x2Gt{}
	case "f64x2.le":
		inst = &F64x2Le{}
	case "f64x2.ge":
		inst = &F64x2Ge{}
	case "v128.not":
		inst = &V128Not{}
	case "v128.and":
		inst = &V128And{}
	case "v128.or":
		inst = &V128Or{}
	case "v128.xor":
		inst = &V128Xor{}
	case "v128.bitselect":
		inst = &V128Bitselect{}
	case "i8x16.neg":
		inst = &I8x16Neg{}
	case "i8x16.any_true":
		inst = &I8x16AnyTrue{}
	case "i8x16.all_true":
		inst = &I8x16AllTrue{}
	case "i8x16.shl":
		inst = &I8x16Shl{}
	case "i8x16.shr_s":
		inst = &I8x16ShrS{}
	case "i8x16.shr_u":
		inst = &I8x16ShrU{}
	case "i8x16.add":
		inst = &I8x16Add{}
	case "i8x16.add_saturate_s":
		inst = &I8x16AddSaturateS{}
	case "i8x16.add_saturate_u":
		inst = &I8x16AddSaturateU{}
	case "i8x16.sub":
		inst = &I8x16Sub{}
	case "i8x16.sub_saturate_s":
		inst = &I8x16SubSaturateS{}
	case "i8x16.sub_saturate_u":
		inst = &I8x16SubSaturateU{}
	case "i8x16.mul":
		inst = &I8x16Mul{}
	case "i16x8.neg":
		inst = &I16x8Neg{}
	case "i16x8.any_true":
		inst = &I16x8AnyTrue{}
	case "i16x8.all_true":
		inst = &I16x8AllTrue{}
	case "i16x8.shl":
		inst = &I16x8Shl{}
	case "i16x8.shr_s":
		inst = &I16x8ShrS{}
	case "i16x8.shr_u":
		inst = &I16x8ShrU{}
	case "i16x8.add":
		inst = &I16x8Add{}
	case "i16x8.add_saturate_s":
		inst = &I16x8AddSaturateS{}
	case "i16x8.add_saturate_u":
		inst = &I16x8AddSaturateU{}
	case "i16x8.sub":
		inst = &I16x8Sub{}
	case "i16x8.sub_saturate_s":
		inst = &I16x8SubSaturateS{}
	case "i16x8.sub_saturate_u":
		inst = &I16x8SubSaturateU{}
	case "i16x8.mul":
		inst = &I16x8Mul{}
	case "i32x4.neg":
		inst = &I32x4Neg{}
	case "i32x4.any_true":
		inst = &I32x4AnyTrue{}
	case "i32x4.all_true":
		inst = &I32x4AllTrue{}
	case "i32x4.shl":
		inst = &I32x4Shl{}
	case "i32x4.shr_s":
		inst = &I32x4ShrS{}
	case "i32x4.shr_u":
		inst = &I32x4ShrU{}
	case "i32x4.add":
		inst = &I32x4Add{}
	case "i32x4.sub":
		inst = &I32x4Sub{}
	case "i32x4.mul":
		inst = &I32x4Mul{}
	case "i64x2.neg":
		inst = &I64x2Neg{}
	case "i64x2.any_true":
		inst = &I64x2AnyTrue{}
	case "i64x2.all_true":
		inst = &I64x2AllTrue{}
	case "i64x2.shl":
		inst = &I64x2Shl{}
	case "i64x2.shr_s":
		inst = &I64x2ShrS{}
	case "i64x2.shr_u":
		inst = &I64x2ShrU{}
	case "i64x2.add":
		inst = &I64x2Add{}
	case "i64x2.sub":
		inst = &I64x2Sub{}
	case "i64x2.mul":
		inst = &I64x2Mul{}
	case "f32x4.abs":
		inst = &F32x4Abs{}
	case "f32x4.neg":
		inst = &F32x4Neg{}
	case "f32x4.sqrt":
		inst = &F32x4Sqrt{}
	case "f32x4.add":
		inst = &F32x4Add{}
	case "f32x4.sub":
		inst = &F32x4Sub{}
	case "f32x4.mul":
		inst = &F32x4Mul{}
	case "f32x4.div":
		inst = &F32x4Div{}
	case "f32x4.min":
		inst = &F32x4Min{}
	case "f32x4.max":
		inst = &F32x4Max{}
	case "f64x2.abs":
		inst = &F64x2Abs{}
	case "f64x2.neg":
		inst = &F64x2Neg{}
	case "f64x2.sqrt":
		inst = &F64x2Sqrt{}
	case "f64x2.add":
		inst = &F64x2Add{}
	case "f64x2.sub":
		inst = &F64x2Sub{}
	case "f64x2.mul":
		inst = &F64x2Mul{}
	case "f64x2.div":
		inst = &F64x2Div{}
	case "f64x2.min":
		inst = &F64x2Min{}
	case "f64x2.max":
		inst = &F64x2Max{}
	case "i32x4.trunc_sat_f32x4_s":
		inst = &I32x4TruncSatF32x4S{}
	case "i32x4.trunc_sat_f32x4_u":
		inst = &I32x4TruncSatF32x4U{}
	case "i64x2.trunc_sat_f64x2_s":
		inst = &I64x2TruncSatF64x2S{}
	case "i64x2.trunc_sat_f64x2_u":
		inst = &I64x2TruncSatF64x2U{}
	case "f32x4.convert_i32x4_s":
		inst = &F32x4ConvertI32x4S{}
	case "f32x4.convert_i32x4_u":
		inst = &F32x4ConvertI32x4U{}
	case "f64x2.convert_i64x2_s":
		inst = &F64x2ConvertI64x2S{}
	case "f64x2.convert_i64x2_u":
		inst = &F64x2ConvertI64x2U{}
	case "v8x16.swizzle":
		inst = &V8x16Swizzle{}
	case "v8x16.load_splat":
		inst = &V8x16LoadSplat{}
	case "v16x8.load_splat":
		inst = &V16x8LoadSplat{}
	case "v32x4.load_splat":
		inst = &V32x4LoadSplat{}
	case "v64x2.load_splat":
		inst = &V64x2LoadSplat{}
	case "i8x16.narrow_i16x8_s":
		inst = &I8x16NarrowI16x8S{}
	case "i8x16.narrow_i16x8_u":
		inst = &I8x16NarrowI16x8U{}
	case "i16x8.narrow_i32x4_s":
		inst = &I16x8NarrowI32x4S{}
	case "i16x8.narrow_i32x4_u":
		inst = &I16x8NarrowI32x4U{}
	case "i16x8.widen_low_i8x16_s":
		inst = &I16x8WidenLowI8x16S{}
	case "i16x8.widen_high_i8x16_s":
		inst = &I16x8WidenHighI8x16S{}
	case "i16x8.widen_low_i8x16_u":
		inst = &I16x8WidenLowI8x16U{}
	case "i16x8.widen_high_i8x16_u":
		inst = &I16x8WidenHighI8x16u{}
	case "i32x4.widen_low_i16x8_s":
		inst = &I32x4WidenLowI16x8S{}
	case "i32x4.widen_high_i16x8_s":
		inst = &I32x4WidenHighI16x8S{}
	case "i32x4.widen_low_i16x8_u":
		inst = &I32x4WidenLowI16x8U{}
	case "i32x4.widen_high_i16x8_u":
		inst = &I32x4WidenHighI16x8u{}
	case "i16x8.load8x8_s":
		inst = &I16x8Load8x8S{}
	case "i16x8.load8x8_u":
		inst = &I16x8Load8x8U{}
	case "i32x4.load16x4_s":
		inst = &I32x4Load16x4S{}
	case "i32x4.load16x4_u":
		inst = &I32x4Load16x4U{}
	case "i64x2.load32x2_s":
		inst = &I64x2Load32x2S{}
	case "i64x2.load32x2_u":
		inst = &I64x2Load32x2U{}
	case "v128.andnot":
		inst = &V128Andnot{}
	case "nan:canonical":
		inst = &CanonicalNan{}
	case "nan:arithmetic":
		inst = &ArithmeticNan{}
	default:
		panic(fmt.Sprintf("todo: implement instruction %s", kw))
	}
	err = inst.parseInstrBody(ps)
	if err != nil {
		return nil, err
	}
	return inst, nil
}
