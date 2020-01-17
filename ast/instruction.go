
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

func (self *Block) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2}
	sink.WriteBytes(inst)
	self.BlockType.Encode(sink)

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

func (self *If) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4}
	sink.WriteBytes(inst)
	self.BlockType.Encode(sink)

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

func (self *Else) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5}
	sink.WriteBytes(inst)
	self.Id.Encode(sink)

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

func (self *Loop) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3}
	sink.WriteBytes(inst)
	self.BlockType.Encode(sink)

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

func (self *End) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb}
	sink.WriteBytes(inst)
	self.Id.Encode(sink)

}

type Unreachable struct {
	
}

func (self *Unreachable) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *Unreachable) String() string {
	return "unreachable"
}

func (self *Unreachable) Encode(sink *ZeroCopySink) {
	inst := []byte{0x0}
	sink.WriteBytes(inst)
	
}

type Nop struct {
	
}

func (self *Nop) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *Nop) String() string {
	return "nop"
}

func (self *Nop) Encode(sink *ZeroCopySink) {
	inst := []byte{0x1}
	sink.WriteBytes(inst)
	
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

func (self *Br) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *BrIf) Encode(sink *ZeroCopySink) {
	inst := []byte{0xd}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *BrTable) Encode(sink *ZeroCopySink) {
	inst := []byte{0xe}
	sink.WriteBytes(inst)
	self.Indices.Encode(sink)

}

type Return struct {
	
}

func (self *Return) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *Return) String() string {
	return "return"
}

func (self *Return) Encode(sink *ZeroCopySink) {
	inst := []byte{0xf}
	sink.WriteBytes(inst)
	
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

func (self *Call) Encode(sink *ZeroCopySink) {
	inst := []byte{0x10}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *CallIndirect) Encode(sink *ZeroCopySink) {
	inst := []byte{0x11}
	sink.WriteBytes(inst)
	self.Impl.Encode(sink)

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

func (self *ReturnCall) Encode(sink *ZeroCopySink) {
	inst := []byte{0x12}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *ReturnCallIndirect) Encode(sink *ZeroCopySink) {
	inst := []byte{0x13}
	sink.WriteBytes(inst)
	self.Impl.Encode(sink)

}

type Drop struct {
	
}

func (self *Drop) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *Drop) String() string {
	return "drop"
}

func (self *Drop) Encode(sink *ZeroCopySink) {
	inst := []byte{0x1a}
	sink.WriteBytes(inst)
	
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

func (self *Select) Encode(sink *ZeroCopySink) {
	inst := []byte{}
	sink.WriteBytes(inst)
	self.SelectTypes.Encode(sink)

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

func (self *LocalGet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x20}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *LocalSet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x21}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *LocalTee) Encode(sink *ZeroCopySink) {
	inst := []byte{0x22}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *GlobalGet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x23}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *GlobalSet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x24}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *TableGet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x25}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *TableSet) Encode(sink *ZeroCopySink) {
	inst := []byte{0x26}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *I32Load) Encode(sink *ZeroCopySink) {
	inst := []byte{0x28}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load) Encode(sink *ZeroCopySink) {
	inst := []byte{0x29}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *F32Load) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *F64Load) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Load8s) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Load8u) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Load16s) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Load16u) Encode(sink *ZeroCopySink) {
	inst := []byte{0x2f}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load8s) Encode(sink *ZeroCopySink) {
	inst := []byte{0x30}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load8u) Encode(sink *ZeroCopySink) {
	inst := []byte{0x31}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load16s) Encode(sink *ZeroCopySink) {
	inst := []byte{0x32}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load16u) Encode(sink *ZeroCopySink) {
	inst := []byte{0x33}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load32s) Encode(sink *ZeroCopySink) {
	inst := []byte{0x34}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Load32u) Encode(sink *ZeroCopySink) {
	inst := []byte{0x35}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Store) Encode(sink *ZeroCopySink) {
	inst := []byte{0x36}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Store) Encode(sink *ZeroCopySink) {
	inst := []byte{0x37}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *F32Store) Encode(sink *ZeroCopySink) {
	inst := []byte{0x38}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *F64Store) Encode(sink *ZeroCopySink) {
	inst := []byte{0x39}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Store8) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32Store16) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Store8) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Store16) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64Store32) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

}

type MemorySize struct {
	
}

func (self *MemorySize) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *MemorySize) String() string {
	return "memory.size"
}

func (self *MemorySize) Encode(sink *ZeroCopySink) {
	inst := []byte{0x3f,0x0}
	sink.WriteBytes(inst)
	
}

type MemoryGrow struct {
	
}

func (self *MemoryGrow) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *MemoryGrow) String() string {
	return "memory.grow"
}

func (self *MemoryGrow) Encode(sink *ZeroCopySink) {
	inst := []byte{0x40,0x0}
	sink.WriteBytes(inst)
	
}

type MemoryCopy struct {
	
}

func (self *MemoryCopy) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *MemoryCopy) String() string {
	return "memory.copy"
}

func (self *MemoryCopy) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0xa,0x0,0x0}
	sink.WriteBytes(inst)
	
}

type MemoryFill struct {
	
}

func (self *MemoryFill) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *MemoryFill) String() string {
	return "memory.fill"
}

func (self *MemoryFill) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0xb,0x0}
	sink.WriteBytes(inst)
	
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

func (self *DataDrop) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x9}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *ElemDrop) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0xd}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

}

type TableCopy struct {
	
}

func (self *TableCopy) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *TableCopy) String() string {
	return "table.copy"
}

func (self *TableCopy) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0xe,0x0,0x0}
	sink.WriteBytes(inst)
	
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

func (self *TableFill) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x11}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *TableSize) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x10}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *TableGrow) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0xf}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

}

type RefNull struct {
	
}

func (self *RefNull) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *RefNull) String() string {
	return "ref.null"
}

func (self *RefNull) Encode(sink *ZeroCopySink) {
	inst := []byte{0xd0}
	sink.WriteBytes(inst)
	
}

type RefIsNull struct {
	
}

func (self *RefIsNull) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *RefIsNull) String() string {
	return "ref.is_null"
}

func (self *RefIsNull) Encode(sink *ZeroCopySink) {
	inst := []byte{0xd1}
	sink.WriteBytes(inst)
	
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

func (self *RefHost) Encode(sink *ZeroCopySink) {
	inst := []byte{0xff}
	sink.WriteBytes(inst)
	sink.WriteInt32(self.Val)

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

func (self *RefFunc) Encode(sink *ZeroCopySink) {
	inst := []byte{0xd2}
	sink.WriteBytes(inst)
	self.Index.Encode(sink)

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

func (self *I32Const) Encode(sink *ZeroCopySink) {
	inst := []byte{0x41}
	sink.WriteBytes(inst)
	sink.WriteInt32(self.Val)

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

func (self *I64Const) Encode(sink *ZeroCopySink) {
	inst := []byte{0x42}
	sink.WriteBytes(inst)
	sink.WriteInt64(self.Val)

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

func (self *F32Const) Encode(sink *ZeroCopySink) {
	inst := []byte{0x43}
	sink.WriteBytes(inst)
	self.Val.Encode(sink)

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

func (self *F64Const) Encode(sink *ZeroCopySink) {
	inst := []byte{0x44}
	sink.WriteBytes(inst)
	self.Val.Encode(sink)

}

type I32Clz struct {
	
}

func (self *I32Clz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Clz) String() string {
	return "i32.clz"
}

func (self *I32Clz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x67}
	sink.WriteBytes(inst)
	
}

type I32Ctz struct {
	
}

func (self *I32Ctz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Ctz) String() string {
	return "i32.ctz"
}

func (self *I32Ctz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x68}
	sink.WriteBytes(inst)
	
}

type I32Pocnt struct {
	
}

func (self *I32Pocnt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Pocnt) String() string {
	return "i32.popcnt"
}

func (self *I32Pocnt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x69}
	sink.WriteBytes(inst)
	
}

type I32Add struct {
	
}

func (self *I32Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Add) String() string {
	return "i32.add"
}

func (self *I32Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6a}
	sink.WriteBytes(inst)
	
}

type I32Sub struct {
	
}

func (self *I32Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Sub) String() string {
	return "i32.sub"
}

func (self *I32Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6b}
	sink.WriteBytes(inst)
	
}

type I32Mul struct {
	
}

func (self *I32Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Mul) String() string {
	return "i32.mul"
}

func (self *I32Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6c}
	sink.WriteBytes(inst)
	
}

type I32DivS struct {
	
}

func (self *I32DivS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32DivS) String() string {
	return "i32.div_s"
}

func (self *I32DivS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6d}
	sink.WriteBytes(inst)
	
}

type I32DivU struct {
	
}

func (self *I32DivU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32DivU) String() string {
	return "i32.div_u"
}

func (self *I32DivU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6e}
	sink.WriteBytes(inst)
	
}

type I32RemS struct {
	
}

func (self *I32RemS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32RemS) String() string {
	return "i32.rem_s"
}

func (self *I32RemS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x6f}
	sink.WriteBytes(inst)
	
}

type I32RemU struct {
	
}

func (self *I32RemU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32RemU) String() string {
	return "i32.rem_u"
}

func (self *I32RemU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x70}
	sink.WriteBytes(inst)
	
}

type I32And struct {
	
}

func (self *I32And) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32And) String() string {
	return "i32.and"
}

func (self *I32And) Encode(sink *ZeroCopySink) {
	inst := []byte{0x71}
	sink.WriteBytes(inst)
	
}

type I32Or struct {
	
}

func (self *I32Or) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Or) String() string {
	return "i32.or"
}

func (self *I32Or) Encode(sink *ZeroCopySink) {
	inst := []byte{0x72}
	sink.WriteBytes(inst)
	
}

type I32Xor struct {
	
}

func (self *I32Xor) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Xor) String() string {
	return "i32.xor"
}

func (self *I32Xor) Encode(sink *ZeroCopySink) {
	inst := []byte{0x73}
	sink.WriteBytes(inst)
	
}

type I32Shl struct {
	
}

func (self *I32Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Shl) String() string {
	return "i32.shl"
}

func (self *I32Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0x74}
	sink.WriteBytes(inst)
	
}

type I32ShrS struct {
	
}

func (self *I32ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32ShrS) String() string {
	return "i32.shr_s"
}

func (self *I32ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x75}
	sink.WriteBytes(inst)
	
}

type I32ShrU struct {
	
}

func (self *I32ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32ShrU) String() string {
	return "i32.shr_u"
}

func (self *I32ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x76}
	sink.WriteBytes(inst)
	
}

type I32Rotl struct {
	
}

func (self *I32Rotl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Rotl) String() string {
	return "i32.rotl"
}

func (self *I32Rotl) Encode(sink *ZeroCopySink) {
	inst := []byte{0x77}
	sink.WriteBytes(inst)
	
}

type I32Rotr struct {
	
}

func (self *I32Rotr) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Rotr) String() string {
	return "i32.rotr"
}

func (self *I32Rotr) Encode(sink *ZeroCopySink) {
	inst := []byte{0x78}
	sink.WriteBytes(inst)
	
}

type I64Clz struct {
	
}

func (self *I64Clz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Clz) String() string {
	return "i64.clz"
}

func (self *I64Clz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x79}
	sink.WriteBytes(inst)
	
}

type I64Ctz struct {
	
}

func (self *I64Ctz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Ctz) String() string {
	return "i64.ctz"
}

func (self *I64Ctz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7a}
	sink.WriteBytes(inst)
	
}

type I64Popcnt struct {
	
}

func (self *I64Popcnt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Popcnt) String() string {
	return "i64.popcnt"
}

func (self *I64Popcnt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7b}
	sink.WriteBytes(inst)
	
}

type I64Add struct {
	
}

func (self *I64Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Add) String() string {
	return "i64.add"
}

func (self *I64Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7c}
	sink.WriteBytes(inst)
	
}

type I64Sub struct {
	
}

func (self *I64Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Sub) String() string {
	return "i64.sub"
}

func (self *I64Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7d}
	sink.WriteBytes(inst)
	
}

type I64Mul struct {
	
}

func (self *I64Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Mul) String() string {
	return "i64.mul"
}

func (self *I64Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7e}
	sink.WriteBytes(inst)
	
}

type I64DivS struct {
	
}

func (self *I64DivS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64DivS) String() string {
	return "i64.div_s"
}

func (self *I64DivS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x7f}
	sink.WriteBytes(inst)
	
}

type I64DivU struct {
	
}

func (self *I64DivU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64DivU) String() string {
	return "i64.div_u"
}

func (self *I64DivU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x80}
	sink.WriteBytes(inst)
	
}

type I64RemS struct {
	
}

func (self *I64RemS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64RemS) String() string {
	return "i64.rem_s"
}

func (self *I64RemS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x81}
	sink.WriteBytes(inst)
	
}

type I64RemU struct {
	
}

func (self *I64RemU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64RemU) String() string {
	return "i64.rem_u"
}

func (self *I64RemU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x82}
	sink.WriteBytes(inst)
	
}

type I64And struct {
	
}

func (self *I64And) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64And) String() string {
	return "i64.and"
}

func (self *I64And) Encode(sink *ZeroCopySink) {
	inst := []byte{0x83}
	sink.WriteBytes(inst)
	
}

type I64Or struct {
	
}

func (self *I64Or) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Or) String() string {
	return "i64.or"
}

func (self *I64Or) Encode(sink *ZeroCopySink) {
	inst := []byte{0x84}
	sink.WriteBytes(inst)
	
}

type I64Xor struct {
	
}

func (self *I64Xor) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Xor) String() string {
	return "i64.xor"
}

func (self *I64Xor) Encode(sink *ZeroCopySink) {
	inst := []byte{0x85}
	sink.WriteBytes(inst)
	
}

type I64Shl struct {
	
}

func (self *I64Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Shl) String() string {
	return "i64.shl"
}

func (self *I64Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0x86}
	sink.WriteBytes(inst)
	
}

type I64ShrS struct {
	
}

func (self *I64ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64ShrS) String() string {
	return "i64.shr_s"
}

func (self *I64ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x87}
	sink.WriteBytes(inst)
	
}

type I64ShrU struct {
	
}

func (self *I64ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64ShrU) String() string {
	return "i64.shr_u"
}

func (self *I64ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x88}
	sink.WriteBytes(inst)
	
}

type I64Rotl struct {
	
}

func (self *I64Rotl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Rotl) String() string {
	return "i64.rotl"
}

func (self *I64Rotl) Encode(sink *ZeroCopySink) {
	inst := []byte{0x89}
	sink.WriteBytes(inst)
	
}

type I64Rotr struct {
	
}

func (self *I64Rotr) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Rotr) String() string {
	return "i64.rotr"
}

func (self *I64Rotr) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8a}
	sink.WriteBytes(inst)
	
}

type F32Abs struct {
	
}

func (self *F32Abs) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Abs) String() string {
	return "f32.abs"
}

func (self *F32Abs) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8b}
	sink.WriteBytes(inst)
	
}

type F32Neg struct {
	
}

func (self *F32Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Neg) String() string {
	return "f32.neg"
}

func (self *F32Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8c}
	sink.WriteBytes(inst)
	
}

type F32Ceil struct {
	
}

func (self *F32Ceil) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Ceil) String() string {
	return "f32.ceil"
}

func (self *F32Ceil) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8d}
	sink.WriteBytes(inst)
	
}

type F32Floor struct {
	
}

func (self *F32Floor) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Floor) String() string {
	return "f32.floor"
}

func (self *F32Floor) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8e}
	sink.WriteBytes(inst)
	
}

type F32Trunc struct {
	
}

func (self *F32Trunc) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Trunc) String() string {
	return "f32.trunc"
}

func (self *F32Trunc) Encode(sink *ZeroCopySink) {
	inst := []byte{0x8f}
	sink.WriteBytes(inst)
	
}

type F32Nearest struct {
	
}

func (self *F32Nearest) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Nearest) String() string {
	return "f32.nearest"
}

func (self *F32Nearest) Encode(sink *ZeroCopySink) {
	inst := []byte{0x90}
	sink.WriteBytes(inst)
	
}

type F32Sqrt struct {
	
}

func (self *F32Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Sqrt) String() string {
	return "f32.sqrt"
}

func (self *F32Sqrt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x91}
	sink.WriteBytes(inst)
	
}

type F32Add struct {
	
}

func (self *F32Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Add) String() string {
	return "f32.add"
}

func (self *F32Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0x92}
	sink.WriteBytes(inst)
	
}

type F32Sub struct {
	
}

func (self *F32Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Sub) String() string {
	return "f32.sub"
}

func (self *F32Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0x93}
	sink.WriteBytes(inst)
	
}

type F32Mul struct {
	
}

func (self *F32Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Mul) String() string {
	return "f32.mul"
}

func (self *F32Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0x94}
	sink.WriteBytes(inst)
	
}

type F32Div struct {
	
}

func (self *F32Div) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Div) String() string {
	return "f32.div"
}

func (self *F32Div) Encode(sink *ZeroCopySink) {
	inst := []byte{0x95}
	sink.WriteBytes(inst)
	
}

type F32Min struct {
	
}

func (self *F32Min) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Min) String() string {
	return "f32.min"
}

func (self *F32Min) Encode(sink *ZeroCopySink) {
	inst := []byte{0x96}
	sink.WriteBytes(inst)
	
}

type F32Max struct {
	
}

func (self *F32Max) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Max) String() string {
	return "f32.max"
}

func (self *F32Max) Encode(sink *ZeroCopySink) {
	inst := []byte{0x97}
	sink.WriteBytes(inst)
	
}

type F32Copysign struct {
	
}

func (self *F32Copysign) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Copysign) String() string {
	return "f32.copysign"
}

func (self *F32Copysign) Encode(sink *ZeroCopySink) {
	inst := []byte{0x98}
	sink.WriteBytes(inst)
	
}

type F64Abs struct {
	
}

func (self *F64Abs) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Abs) String() string {
	return "f64.abs"
}

func (self *F64Abs) Encode(sink *ZeroCopySink) {
	inst := []byte{0x99}
	sink.WriteBytes(inst)
	
}

type F64Neg struct {
	
}

func (self *F64Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Neg) String() string {
	return "f64.neg"
}

func (self *F64Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9a}
	sink.WriteBytes(inst)
	
}

type F64Ceil struct {
	
}

func (self *F64Ceil) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Ceil) String() string {
	return "f64.ceil"
}

func (self *F64Ceil) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9b}
	sink.WriteBytes(inst)
	
}

type F64Floor struct {
	
}

func (self *F64Floor) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Floor) String() string {
	return "f64.floor"
}

func (self *F64Floor) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9c}
	sink.WriteBytes(inst)
	
}

type F64Trunc struct {
	
}

func (self *F64Trunc) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Trunc) String() string {
	return "f64.trunc"
}

func (self *F64Trunc) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9d}
	sink.WriteBytes(inst)
	
}

type F64Nearest struct {
	
}

func (self *F64Nearest) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Nearest) String() string {
	return "f64.nearest"
}

func (self *F64Nearest) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9e}
	sink.WriteBytes(inst)
	
}

type F64Sqrt struct {
	
}

func (self *F64Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Sqrt) String() string {
	return "f64.sqrt"
}

func (self *F64Sqrt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x9f}
	sink.WriteBytes(inst)
	
}

type F64Add struct {
	
}

func (self *F64Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Add) String() string {
	return "f64.add"
}

func (self *F64Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa0}
	sink.WriteBytes(inst)
	
}

type F64Sub struct {
	
}

func (self *F64Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Sub) String() string {
	return "f64.sub"
}

func (self *F64Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa1}
	sink.WriteBytes(inst)
	
}

type F64Mul struct {
	
}

func (self *F64Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Mul) String() string {
	return "f64.mul"
}

func (self *F64Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa2}
	sink.WriteBytes(inst)
	
}

type F64Div struct {
	
}

func (self *F64Div) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Div) String() string {
	return "f64.div"
}

func (self *F64Div) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa3}
	sink.WriteBytes(inst)
	
}

type F64Min struct {
	
}

func (self *F64Min) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Min) String() string {
	return "f64.min"
}

func (self *F64Min) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa4}
	sink.WriteBytes(inst)
	
}

type F64Max struct {
	
}

func (self *F64Max) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Max) String() string {
	return "f64.max"
}

func (self *F64Max) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa5}
	sink.WriteBytes(inst)
	
}

type F64Copysign struct {
	
}

func (self *F64Copysign) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Copysign) String() string {
	return "f64.copysign"
}

func (self *F64Copysign) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa6}
	sink.WriteBytes(inst)
	
}

type I32Eqz struct {
	
}

func (self *I32Eqz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Eqz) String() string {
	return "i32.eqz"
}

func (self *I32Eqz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x45}
	sink.WriteBytes(inst)
	
}

type I32Eq struct {
	
}

func (self *I32Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Eq) String() string {
	return "i32.eq"
}

func (self *I32Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0x46}
	sink.WriteBytes(inst)
	
}

type I32Ne struct {
	
}

func (self *I32Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Ne) String() string {
	return "i32.ne"
}

func (self *I32Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0x47}
	sink.WriteBytes(inst)
	
}

type I32LtS struct {
	
}

func (self *I32LtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32LtS) String() string {
	return "i32.lt_s"
}

func (self *I32LtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x48}
	sink.WriteBytes(inst)
	
}

type I32LtU struct {
	
}

func (self *I32LtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32LtU) String() string {
	return "i32.lt_u"
}

func (self *I32LtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x49}
	sink.WriteBytes(inst)
	
}

type I32GtS struct {
	
}

func (self *I32GtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32GtS) String() string {
	return "i32.gt_s"
}

func (self *I32GtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4a}
	sink.WriteBytes(inst)
	
}

type I32GtU struct {
	
}

func (self *I32GtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32GtU) String() string {
	return "i32.gt_u"
}

func (self *I32GtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4b}
	sink.WriteBytes(inst)
	
}

type I32LeS struct {
	
}

func (self *I32LeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32LeS) String() string {
	return "i32.le_s"
}

func (self *I32LeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4c}
	sink.WriteBytes(inst)
	
}

type I32LeU struct {
	
}

func (self *I32LeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32LeU) String() string {
	return "i32.le_u"
}

func (self *I32LeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4d}
	sink.WriteBytes(inst)
	
}

type I32GeS struct {
	
}

func (self *I32GeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32GeS) String() string {
	return "i32.ge_s"
}

func (self *I32GeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4e}
	sink.WriteBytes(inst)
	
}

type I32GeU struct {
	
}

func (self *I32GeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32GeU) String() string {
	return "i32.ge_u"
}

func (self *I32GeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x4f}
	sink.WriteBytes(inst)
	
}

type I64Eqz struct {
	
}

func (self *I64Eqz) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Eqz) String() string {
	return "i64.eqz"
}

func (self *I64Eqz) Encode(sink *ZeroCopySink) {
	inst := []byte{0x50}
	sink.WriteBytes(inst)
	
}

type I64Eq struct {
	
}

func (self *I64Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Eq) String() string {
	return "i64.eq"
}

func (self *I64Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0x51}
	sink.WriteBytes(inst)
	
}

type I64Ne struct {
	
}

func (self *I64Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Ne) String() string {
	return "i64.ne"
}

func (self *I64Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0x52}
	sink.WriteBytes(inst)
	
}

type I64LtS struct {
	
}

func (self *I64LtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64LtS) String() string {
	return "i64.lt_s"
}

func (self *I64LtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x53}
	sink.WriteBytes(inst)
	
}

type I64LtU struct {
	
}

func (self *I64LtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64LtU) String() string {
	return "i64.lt_u"
}

func (self *I64LtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x54}
	sink.WriteBytes(inst)
	
}

type I64GtS struct {
	
}

func (self *I64GtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64GtS) String() string {
	return "i64.gt_s"
}

func (self *I64GtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x55}
	sink.WriteBytes(inst)
	
}

type I64GtU struct {
	
}

func (self *I64GtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64GtU) String() string {
	return "i64.gt_u"
}

func (self *I64GtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x56}
	sink.WriteBytes(inst)
	
}

type I64LeS struct {
	
}

func (self *I64LeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64LeS) String() string {
	return "i64.le_s"
}

func (self *I64LeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x57}
	sink.WriteBytes(inst)
	
}

type I64LeU struct {
	
}

func (self *I64LeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64LeU) String() string {
	return "i64.le_u"
}

func (self *I64LeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x58}
	sink.WriteBytes(inst)
	
}

type I64GeS struct {
	
}

func (self *I64GeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64GeS) String() string {
	return "i64.ge_s"
}

func (self *I64GeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0x59}
	sink.WriteBytes(inst)
	
}

type I64GeU struct {
	
}

func (self *I64GeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64GeU) String() string {
	return "i64.ge_u"
}

func (self *I64GeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5a}
	sink.WriteBytes(inst)
	
}

type F32Eq struct {
	
}

func (self *F32Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Eq) String() string {
	return "f32.eq"
}

func (self *F32Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5b}
	sink.WriteBytes(inst)
	
}

type F32Ne struct {
	
}

func (self *F32Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Ne) String() string {
	return "f32.ne"
}

func (self *F32Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5c}
	sink.WriteBytes(inst)
	
}

type F32Lt struct {
	
}

func (self *F32Lt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Lt) String() string {
	return "f32.lt"
}

func (self *F32Lt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5d}
	sink.WriteBytes(inst)
	
}

type F32Gt struct {
	
}

func (self *F32Gt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Gt) String() string {
	return "f32.gt"
}

func (self *F32Gt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5e}
	sink.WriteBytes(inst)
	
}

type F32Le struct {
	
}

func (self *F32Le) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Le) String() string {
	return "f32.le"
}

func (self *F32Le) Encode(sink *ZeroCopySink) {
	inst := []byte{0x5f}
	sink.WriteBytes(inst)
	
}

type F32Ge struct {
	
}

func (self *F32Ge) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32Ge) String() string {
	return "f32.ge"
}

func (self *F32Ge) Encode(sink *ZeroCopySink) {
	inst := []byte{0x60}
	sink.WriteBytes(inst)
	
}

type F64Eq struct {
	
}

func (self *F64Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Eq) String() string {
	return "f64.eq"
}

func (self *F64Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0x61}
	sink.WriteBytes(inst)
	
}

type F64Ne struct {
	
}

func (self *F64Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Ne) String() string {
	return "f64.ne"
}

func (self *F64Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0x62}
	sink.WriteBytes(inst)
	
}

type F64Lt struct {
	
}

func (self *F64Lt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Lt) String() string {
	return "f64.lt"
}

func (self *F64Lt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x63}
	sink.WriteBytes(inst)
	
}

type F64Gt struct {
	
}

func (self *F64Gt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Gt) String() string {
	return "f64.gt"
}

func (self *F64Gt) Encode(sink *ZeroCopySink) {
	inst := []byte{0x64}
	sink.WriteBytes(inst)
	
}

type F64Le struct {
	
}

func (self *F64Le) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Le) String() string {
	return "f64.le"
}

func (self *F64Le) Encode(sink *ZeroCopySink) {
	inst := []byte{0x65}
	sink.WriteBytes(inst)
	
}

type F64Ge struct {
	
}

func (self *F64Ge) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64Ge) String() string {
	return "f64.ge"
}

func (self *F64Ge) Encode(sink *ZeroCopySink) {
	inst := []byte{0x66}
	sink.WriteBytes(inst)
	
}

type I32WrapI64 struct {
	
}

func (self *I32WrapI64) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32WrapI64) String() string {
	return "i32.wrap_i64"
}

func (self *I32WrapI64) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa7}
	sink.WriteBytes(inst)
	
}

type I32TruncF32S struct {
	
}

func (self *I32TruncF32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncF32S) String() string {
	return "i32.trunc_f32_s"
}

func (self *I32TruncF32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa8}
	sink.WriteBytes(inst)
	
}

type I32TruncF32U struct {
	
}

func (self *I32TruncF32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncF32U) String() string {
	return "i32.trunc_f32_u"
}

func (self *I32TruncF32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xa9}
	sink.WriteBytes(inst)
	
}

type I32TruncF64S struct {
	
}

func (self *I32TruncF64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncF64S) String() string {
	return "i32.trunc_f64_s"
}

func (self *I32TruncF64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xaa}
	sink.WriteBytes(inst)
	
}

type I32TruncF64U struct {
	
}

func (self *I32TruncF64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncF64U) String() string {
	return "i32.trunc_f64_u"
}

func (self *I32TruncF64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xab}
	sink.WriteBytes(inst)
	
}

type I64ExtendI32S struct {
	
}

func (self *I64ExtendI32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64ExtendI32S) String() string {
	return "i64.extend_i32_s"
}

func (self *I64ExtendI32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xac}
	sink.WriteBytes(inst)
	
}

type I64ExtendI32U struct {
	
}

func (self *I64ExtendI32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64ExtendI32U) String() string {
	return "i64.extend_i32_u"
}

func (self *I64ExtendI32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xad}
	sink.WriteBytes(inst)
	
}

type I64TruncF32S struct {
	
}

func (self *I64TruncF32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncF32S) String() string {
	return "i64.trunc_f32_s"
}

func (self *I64TruncF32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xae}
	sink.WriteBytes(inst)
	
}

type I64TruncF32U struct {
	
}

func (self *I64TruncF32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncF32U) String() string {
	return "i64.trunc_f32_u"
}

func (self *I64TruncF32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xaf}
	sink.WriteBytes(inst)
	
}

type I64TruncF64S struct {
	
}

func (self *I64TruncF64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncF64S) String() string {
	return "i64.trunc_f64_s"
}

func (self *I64TruncF64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb0}
	sink.WriteBytes(inst)
	
}

type I64TruncF64U struct {
	
}

func (self *I64TruncF64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncF64U) String() string {
	return "i64.trunc_f64_u"
}

func (self *I64TruncF64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb1}
	sink.WriteBytes(inst)
	
}

type F32ConvertI32S struct {
	
}

func (self *F32ConvertI32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32ConvertI32S) String() string {
	return "f32.convert_i32_s"
}

func (self *F32ConvertI32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb2}
	sink.WriteBytes(inst)
	
}

type F32ConvertI32U struct {
	
}

func (self *F32ConvertI32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32ConvertI32U) String() string {
	return "f32.convert_i32_u"
}

func (self *F32ConvertI32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb3}
	sink.WriteBytes(inst)
	
}

type F32ConvertI64S struct {
	
}

func (self *F32ConvertI64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32ConvertI64S) String() string {
	return "f32.convert_i64_s"
}

func (self *F32ConvertI64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb4}
	sink.WriteBytes(inst)
	
}

type F32ConvertI64U struct {
	
}

func (self *F32ConvertI64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32ConvertI64U) String() string {
	return "f32.convert_i64_u"
}

func (self *F32ConvertI64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb5}
	sink.WriteBytes(inst)
	
}

type F32DemoteF64 struct {
	
}

func (self *F32DemoteF64) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32DemoteF64) String() string {
	return "f32.demote_f64"
}

func (self *F32DemoteF64) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb6}
	sink.WriteBytes(inst)
	
}

type F64ConvertI32S struct {
	
}

func (self *F64ConvertI32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64ConvertI32S) String() string {
	return "f64.convert_i32_s"
}

func (self *F64ConvertI32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb7}
	sink.WriteBytes(inst)
	
}

type F64ConvertI32U struct {
	
}

func (self *F64ConvertI32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64ConvertI32U) String() string {
	return "f64.convert_i32_u"
}

func (self *F64ConvertI32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb8}
	sink.WriteBytes(inst)
	
}

type F64ConvertI64S struct {
	
}

func (self *F64ConvertI64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64ConvertI64S) String() string {
	return "f64.convert_i64_s"
}

func (self *F64ConvertI64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xb9}
	sink.WriteBytes(inst)
	
}

type F64ConvertI64U struct {
	
}

func (self *F64ConvertI64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64ConvertI64U) String() string {
	return "f64.convert_i64_u"
}

func (self *F64ConvertI64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xba}
	sink.WriteBytes(inst)
	
}

type F64PromoteF32 struct {
	
}

func (self *F64PromoteF32) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64PromoteF32) String() string {
	return "f64.promote_f32"
}

func (self *F64PromoteF32) Encode(sink *ZeroCopySink) {
	inst := []byte{0xbb}
	sink.WriteBytes(inst)
	
}

type I32ReinterpretF32 struct {
	
}

func (self *I32ReinterpretF32) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32ReinterpretF32) String() string {
	return "i32.reinterpret_f32"
}

func (self *I32ReinterpretF32) Encode(sink *ZeroCopySink) {
	inst := []byte{0xbc}
	sink.WriteBytes(inst)
	
}

type I64ReinterpretF64 struct {
	
}

func (self *I64ReinterpretF64) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64ReinterpretF64) String() string {
	return "i64.reinterpret_f64"
}

func (self *I64ReinterpretF64) Encode(sink *ZeroCopySink) {
	inst := []byte{0xbd}
	sink.WriteBytes(inst)
	
}

type F32ReinterpretI32 struct {
	
}

func (self *F32ReinterpretI32) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32ReinterpretI32) String() string {
	return "f32.reinterpret_i32"
}

func (self *F32ReinterpretI32) Encode(sink *ZeroCopySink) {
	inst := []byte{0xbe}
	sink.WriteBytes(inst)
	
}

type F64ReinterpretI64 struct {
	
}

func (self *F64ReinterpretI64) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64ReinterpretI64) String() string {
	return "f64.reinterpret_i64"
}

func (self *F64ReinterpretI64) Encode(sink *ZeroCopySink) {
	inst := []byte{0xbf}
	sink.WriteBytes(inst)
	
}

type I32TruncSatF32S struct {
	
}

func (self *I32TruncSatF32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncSatF32S) String() string {
	return "i32.trunc_sat_f32_s"
}

func (self *I32TruncSatF32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x0}
	sink.WriteBytes(inst)
	
}

type I32TruncSatF32U struct {
	
}

func (self *I32TruncSatF32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncSatF32U) String() string {
	return "i32.trunc_sat_f32_u"
}

func (self *I32TruncSatF32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x1}
	sink.WriteBytes(inst)
	
}

type I32TruncSatF64S struct {
	
}

func (self *I32TruncSatF64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncSatF64S) String() string {
	return "i32.trunc_sat_f64_s"
}

func (self *I32TruncSatF64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x2}
	sink.WriteBytes(inst)
	
}

type I32TruncSatF64U struct {
	
}

func (self *I32TruncSatF64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32TruncSatF64U) String() string {
	return "i32.trunc_sat_f64_u"
}

func (self *I32TruncSatF64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x3}
	sink.WriteBytes(inst)
	
}

type I64TruncSatF32S struct {
	
}

func (self *I64TruncSatF32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncSatF32S) String() string {
	return "i64.trunc_sat_f32_s"
}

func (self *I64TruncSatF32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x4}
	sink.WriteBytes(inst)
	
}

type I64TruncSatF32U struct {
	
}

func (self *I64TruncSatF32U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncSatF32U) String() string {
	return "i64.trunc_sat_f32_u"
}

func (self *I64TruncSatF32U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x5}
	sink.WriteBytes(inst)
	
}

type I64TruncSatF64S struct {
	
}

func (self *I64TruncSatF64S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncSatF64S) String() string {
	return "i64.trunc_sat_f64_s"
}

func (self *I64TruncSatF64S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x6}
	sink.WriteBytes(inst)
	
}

type I64TruncSatF64U struct {
	
}

func (self *I64TruncSatF64U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64TruncSatF64U) String() string {
	return "i64.trunc_sat_f64_u"
}

func (self *I64TruncSatF64U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfc,0x7}
	sink.WriteBytes(inst)
	
}

type I32Extend8S struct {
	
}

func (self *I32Extend8S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Extend8S) String() string {
	return "i32.extend8_s"
}

func (self *I32Extend8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc0}
	sink.WriteBytes(inst)
	
}

type I32Extend16S struct {
	
}

func (self *I32Extend16S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32Extend16S) String() string {
	return "i32.extend16_s"
}

func (self *I32Extend16S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc1}
	sink.WriteBytes(inst)
	
}

type I64Extend8S struct {
	
}

func (self *I64Extend8S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Extend8S) String() string {
	return "i64.extend8_s"
}

func (self *I64Extend8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc2}
	sink.WriteBytes(inst)
	
}

type I64Extend16S struct {
	
}

func (self *I64Extend16S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Extend16S) String() string {
	return "i64.extend16_s"
}

func (self *I64Extend16S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc3}
	sink.WriteBytes(inst)
	
}

type I64Extend32S struct {
	
}

func (self *I64Extend32S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64Extend32S) String() string {
	return "i64.extend32_s"
}

func (self *I64Extend32S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xc4}
	sink.WriteBytes(inst)
	
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

func (self *AtomicNotify) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x0}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicWait) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicWait) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

}

type AtomicFence struct {
	
}

func (self *AtomicFence) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *AtomicFence) String() string {
	return "atomic.fence"
}

func (self *AtomicFence) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3}
	sink.WriteBytes(inst)
	
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

func (self *I32AtomicLoad) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x10}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicLoad) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x11}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicLoad8u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x12}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicLoad16u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x13}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicLoad8u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x14}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicLoad16u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x15}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicLoad32u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x16}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicStore) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x17}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicStore) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x18}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicStore8) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x19}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicStore16) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicStore8) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicStore16) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicStore32) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwAdd) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwAdd) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x1f}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8AddU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x20}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16AddU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x21}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8AddU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x22}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16AddU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x23}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32AddU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x24}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwSub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x25}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwSub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x26}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8SubU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x27}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16SubU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x28}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8SubU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x29}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16SubU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32SubU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwAnd) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwAnd) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8AndU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16AndU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x2f}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8AndU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x30}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16AndU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x31}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32AndU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x32}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwOr) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x33}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwOr) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x34}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8OrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x35}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16OrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x36}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8OrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x37}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16OrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x38}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32OrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x39}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwXor) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwXor) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8XorU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16XorU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8XorU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16XorU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x3f}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32XorU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x40}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwXchg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x41}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwXchg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x42}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8XchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x43}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16XchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x44}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8XchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x45}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16XchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x46}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32XchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x47}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmwCmpxchg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x48}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmwCmpxchg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x49}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw8CmpxchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x4a}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32AtomicRmw16CmpxchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x4b}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw8CmpxchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x4c}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw16CmpxchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x4d}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64AtomicRmw32CmpxchgU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfe,0x4e}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *V128Load) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x0}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *V128Store) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

}

type I8x16Eq struct {
	
}

func (self *I8x16Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Eq) String() string {
	return "i8x16.eq"
}

func (self *I8x16Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x18}
	sink.WriteBytes(inst)
	
}

type I8x16Ne struct {
	
}

func (self *I8x16Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Ne) String() string {
	return "i8x16.ne"
}

func (self *I8x16Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x19}
	sink.WriteBytes(inst)
	
}

type I8x16LtS struct {
	
}

func (self *I8x16LtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16LtS) String() string {
	return "i8x16.lt_s"
}

func (self *I8x16LtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1a}
	sink.WriteBytes(inst)
	
}

type I8x16LtU struct {
	
}

func (self *I8x16LtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16LtU) String() string {
	return "i8x16.lt_u"
}

func (self *I8x16LtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1b}
	sink.WriteBytes(inst)
	
}

type I8x16GtS struct {
	
}

func (self *I8x16GtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16GtS) String() string {
	return "i8x16.gt_s"
}

func (self *I8x16GtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1c}
	sink.WriteBytes(inst)
	
}

type I8x16GtU struct {
	
}

func (self *I8x16GtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16GtU) String() string {
	return "i8x16.gt_u"
}

func (self *I8x16GtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1d}
	sink.WriteBytes(inst)
	
}

type I8x16LeS struct {
	
}

func (self *I8x16LeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16LeS) String() string {
	return "i8x16.le_s"
}

func (self *I8x16LeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1e}
	sink.WriteBytes(inst)
	
}

type I8x16LeU struct {
	
}

func (self *I8x16LeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16LeU) String() string {
	return "i8x16.le_u"
}

func (self *I8x16LeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x1f}
	sink.WriteBytes(inst)
	
}

type I8x16GeS struct {
	
}

func (self *I8x16GeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16GeS) String() string {
	return "i8x16.ge_s"
}

func (self *I8x16GeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x20}
	sink.WriteBytes(inst)
	
}

type I8x16GeU struct {
	
}

func (self *I8x16GeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16GeU) String() string {
	return "i8x16.ge_u"
}

func (self *I8x16GeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x21}
	sink.WriteBytes(inst)
	
}

type I16x8Eq struct {
	
}

func (self *I16x8Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Eq) String() string {
	return "i16x8.eq"
}

func (self *I16x8Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x22}
	sink.WriteBytes(inst)
	
}

type I16x8Ne struct {
	
}

func (self *I16x8Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Ne) String() string {
	return "i16x8.ne"
}

func (self *I16x8Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x23}
	sink.WriteBytes(inst)
	
}

type I16x8LtS struct {
	
}

func (self *I16x8LtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8LtS) String() string {
	return "i16x8.lt_s"
}

func (self *I16x8LtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x24}
	sink.WriteBytes(inst)
	
}

type I16x8LtU struct {
	
}

func (self *I16x8LtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8LtU) String() string {
	return "i16x8.lt_u"
}

func (self *I16x8LtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x25}
	sink.WriteBytes(inst)
	
}

type I16x8GtS struct {
	
}

func (self *I16x8GtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8GtS) String() string {
	return "i16x8.gt_s"
}

func (self *I16x8GtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x26}
	sink.WriteBytes(inst)
	
}

type I16x8GtU struct {
	
}

func (self *I16x8GtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8GtU) String() string {
	return "i16x8.gt_u"
}

func (self *I16x8GtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x27}
	sink.WriteBytes(inst)
	
}

type I16x8LeS struct {
	
}

func (self *I16x8LeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8LeS) String() string {
	return "i16x8.le_s"
}

func (self *I16x8LeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x28}
	sink.WriteBytes(inst)
	
}

type I16x8LeU struct {
	
}

func (self *I16x8LeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8LeU) String() string {
	return "i16x8.le_u"
}

func (self *I16x8LeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x29}
	sink.WriteBytes(inst)
	
}

type I16x8GeS struct {
	
}

func (self *I16x8GeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8GeS) String() string {
	return "i16x8.ge_s"
}

func (self *I16x8GeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2a}
	sink.WriteBytes(inst)
	
}

type I16x8GeU struct {
	
}

func (self *I16x8GeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8GeU) String() string {
	return "i16x8.ge_u"
}

func (self *I16x8GeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2b}
	sink.WriteBytes(inst)
	
}

type I32x4Eq struct {
	
}

func (self *I32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Eq) String() string {
	return "i32x4.eq"
}

func (self *I32x4Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2c}
	sink.WriteBytes(inst)
	
}

type I32x4Ne struct {
	
}

func (self *I32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Ne) String() string {
	return "i32x4.ne"
}

func (self *I32x4Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2d}
	sink.WriteBytes(inst)
	
}

type I32x4LtS struct {
	
}

func (self *I32x4LtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4LtS) String() string {
	return "i32x4.lt_s"
}

func (self *I32x4LtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2e}
	sink.WriteBytes(inst)
	
}

type I32x4LtU struct {
	
}

func (self *I32x4LtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4LtU) String() string {
	return "i32x4.lt_u"
}

func (self *I32x4LtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x2f}
	sink.WriteBytes(inst)
	
}

type I32x4GtS struct {
	
}

func (self *I32x4GtS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4GtS) String() string {
	return "i32x4.gt_s"
}

func (self *I32x4GtS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x30}
	sink.WriteBytes(inst)
	
}

type I32x4GtU struct {
	
}

func (self *I32x4GtU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4GtU) String() string {
	return "i32x4.gt_u"
}

func (self *I32x4GtU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x31}
	sink.WriteBytes(inst)
	
}

type I32x4LeS struct {
	
}

func (self *I32x4LeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4LeS) String() string {
	return "i32x4.le_s"
}

func (self *I32x4LeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x32}
	sink.WriteBytes(inst)
	
}

type I32x4LeU struct {
	
}

func (self *I32x4LeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4LeU) String() string {
	return "i32x4.le_u"
}

func (self *I32x4LeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x33}
	sink.WriteBytes(inst)
	
}

type I32x4GeS struct {
	
}

func (self *I32x4GeS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4GeS) String() string {
	return "i32x4.ge_s"
}

func (self *I32x4GeS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x34}
	sink.WriteBytes(inst)
	
}

type I32x4GeU struct {
	
}

func (self *I32x4GeU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4GeU) String() string {
	return "i32x4.ge_u"
}

func (self *I32x4GeU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x35}
	sink.WriteBytes(inst)
	
}

type F32x4Eq struct {
	
}

func (self *F32x4Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Eq) String() string {
	return "f32x4.eq"
}

func (self *F32x4Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x40}
	sink.WriteBytes(inst)
	
}

type F32x4Ne struct {
	
}

func (self *F32x4Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Ne) String() string {
	return "f32x4.ne"
}

func (self *F32x4Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x41}
	sink.WriteBytes(inst)
	
}

type F32x4Lt struct {
	
}

func (self *F32x4Lt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Lt) String() string {
	return "f32x4.lt"
}

func (self *F32x4Lt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x42}
	sink.WriteBytes(inst)
	
}

type F32x4Gt struct {
	
}

func (self *F32x4Gt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Gt) String() string {
	return "f32x4.gt"
}

func (self *F32x4Gt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x43}
	sink.WriteBytes(inst)
	
}

type F32x4Le struct {
	
}

func (self *F32x4Le) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Le) String() string {
	return "f32x4.le"
}

func (self *F32x4Le) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x44}
	sink.WriteBytes(inst)
	
}

type F32x4Ge struct {
	
}

func (self *F32x4Ge) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Ge) String() string {
	return "f32x4.ge"
}

func (self *F32x4Ge) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x45}
	sink.WriteBytes(inst)
	
}

type F64x2Eq struct {
	
}

func (self *F64x2Eq) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Eq) String() string {
	return "f64x2.eq"
}

func (self *F64x2Eq) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x46}
	sink.WriteBytes(inst)
	
}

type F64x2Ne struct {
	
}

func (self *F64x2Ne) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Ne) String() string {
	return "f64x2.ne"
}

func (self *F64x2Ne) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x47}
	sink.WriteBytes(inst)
	
}

type F64x2Lt struct {
	
}

func (self *F64x2Lt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Lt) String() string {
	return "f64x2.lt"
}

func (self *F64x2Lt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x48}
	sink.WriteBytes(inst)
	
}

type F64x2Gt struct {
	
}

func (self *F64x2Gt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Gt) String() string {
	return "f64x2.gt"
}

func (self *F64x2Gt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x49}
	sink.WriteBytes(inst)
	
}

type F64x2Le struct {
	
}

func (self *F64x2Le) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Le) String() string {
	return "f64x2.le"
}

func (self *F64x2Le) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4a}
	sink.WriteBytes(inst)
	
}

type F64x2Ge struct {
	
}

func (self *F64x2Ge) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Ge) String() string {
	return "f64x2.ge"
}

func (self *F64x2Ge) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4b}
	sink.WriteBytes(inst)
	
}

type V128Not struct {
	
}

func (self *V128Not) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128Not) String() string {
	return "v128.not"
}

func (self *V128Not) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4c}
	sink.WriteBytes(inst)
	
}

type V128And struct {
	
}

func (self *V128And) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128And) String() string {
	return "v128.and"
}

func (self *V128And) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4d}
	sink.WriteBytes(inst)
	
}

type V128Or struct {
	
}

func (self *V128Or) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128Or) String() string {
	return "v128.or"
}

func (self *V128Or) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4e}
	sink.WriteBytes(inst)
	
}

type V128Xor struct {
	
}

func (self *V128Xor) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128Xor) String() string {
	return "v128.xor"
}

func (self *V128Xor) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x4f}
	sink.WriteBytes(inst)
	
}

type V128Bitselect struct {
	
}

func (self *V128Bitselect) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128Bitselect) String() string {
	return "v128.bitselect"
}

func (self *V128Bitselect) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x50}
	sink.WriteBytes(inst)
	
}

type I8x16Neg struct {
	
}

func (self *I8x16Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Neg) String() string {
	return "i8x16.neg"
}

func (self *I8x16Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x51}
	sink.WriteBytes(inst)
	
}

type I8x16AnyTrue struct {
	
}

func (self *I8x16AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16AnyTrue) String() string {
	return "i8x16.any_true"
}

func (self *I8x16AnyTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x52}
	sink.WriteBytes(inst)
	
}

type I8x16AllTrue struct {
	
}

func (self *I8x16AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16AllTrue) String() string {
	return "i8x16.all_true"
}

func (self *I8x16AllTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x53}
	sink.WriteBytes(inst)
	
}

type I8x16Shl struct {
	
}

func (self *I8x16Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Shl) String() string {
	return "i8x16.shl"
}

func (self *I8x16Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x54}
	sink.WriteBytes(inst)
	
}

type I8x16ShrS struct {
	
}

func (self *I8x16ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16ShrS) String() string {
	return "i8x16.shr_s"
}

func (self *I8x16ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x55}
	sink.WriteBytes(inst)
	
}

type I8x16ShrU struct {
	
}

func (self *I8x16ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16ShrU) String() string {
	return "i8x16.shr_u"
}

func (self *I8x16ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x56}
	sink.WriteBytes(inst)
	
}

type I8x16Add struct {
	
}

func (self *I8x16Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Add) String() string {
	return "i8x16.add"
}

func (self *I8x16Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x57}
	sink.WriteBytes(inst)
	
}

type I8x16AddSaturateS struct {
	
}

func (self *I8x16AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16AddSaturateS) String() string {
	return "i8x16.add_saturate_s"
}

func (self *I8x16AddSaturateS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x58}
	sink.WriteBytes(inst)
	
}

type I8x16AddSaturateU struct {
	
}

func (self *I8x16AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16AddSaturateU) String() string {
	return "i8x16.add_saturate_u"
}

func (self *I8x16AddSaturateU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x59}
	sink.WriteBytes(inst)
	
}

type I8x16Sub struct {
	
}

func (self *I8x16Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Sub) String() string {
	return "i8x16.sub"
}

func (self *I8x16Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x5a}
	sink.WriteBytes(inst)
	
}

type I8x16SubSaturateS struct {
	
}

func (self *I8x16SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16SubSaturateS) String() string {
	return "i8x16.sub_saturate_s"
}

func (self *I8x16SubSaturateS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x5b}
	sink.WriteBytes(inst)
	
}

type I8x16SubSaturateU struct {
	
}

func (self *I8x16SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16SubSaturateU) String() string {
	return "i8x16.sub_saturate_u"
}

func (self *I8x16SubSaturateU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x5c}
	sink.WriteBytes(inst)
	
}

type I8x16Mul struct {
	
}

func (self *I8x16Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16Mul) String() string {
	return "i8x16.mul"
}

func (self *I8x16Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x5d}
	sink.WriteBytes(inst)
	
}

type I16x8Neg struct {
	
}

func (self *I16x8Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Neg) String() string {
	return "i16x8.neg"
}

func (self *I16x8Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x62}
	sink.WriteBytes(inst)
	
}

type I16x8AnyTrue struct {
	
}

func (self *I16x8AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8AnyTrue) String() string {
	return "i16x8.any_true"
}

func (self *I16x8AnyTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x63}
	sink.WriteBytes(inst)
	
}

type I16x8AllTrue struct {
	
}

func (self *I16x8AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8AllTrue) String() string {
	return "i16x8.all_true"
}

func (self *I16x8AllTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x64}
	sink.WriteBytes(inst)
	
}

type I16x8Shl struct {
	
}

func (self *I16x8Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Shl) String() string {
	return "i16x8.shl"
}

func (self *I16x8Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x65}
	sink.WriteBytes(inst)
	
}

type I16x8ShrS struct {
	
}

func (self *I16x8ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8ShrS) String() string {
	return "i16x8.shr_s"
}

func (self *I16x8ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x66}
	sink.WriteBytes(inst)
	
}

type I16x8ShrU struct {
	
}

func (self *I16x8ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8ShrU) String() string {
	return "i16x8.shr_u"
}

func (self *I16x8ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x67}
	sink.WriteBytes(inst)
	
}

type I16x8Add struct {
	
}

func (self *I16x8Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Add) String() string {
	return "i16x8.add"
}

func (self *I16x8Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x68}
	sink.WriteBytes(inst)
	
}

type I16x8AddSaturateS struct {
	
}

func (self *I16x8AddSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8AddSaturateS) String() string {
	return "i16x8.add_saturate_s"
}

func (self *I16x8AddSaturateS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x69}
	sink.WriteBytes(inst)
	
}

type I16x8AddSaturateU struct {
	
}

func (self *I16x8AddSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8AddSaturateU) String() string {
	return "i16x8.add_saturate_u"
}

func (self *I16x8AddSaturateU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x6a}
	sink.WriteBytes(inst)
	
}

type I16x8Sub struct {
	
}

func (self *I16x8Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Sub) String() string {
	return "i16x8.sub"
}

func (self *I16x8Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x6b}
	sink.WriteBytes(inst)
	
}

type I16x8SubSaturateS struct {
	
}

func (self *I16x8SubSaturateS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8SubSaturateS) String() string {
	return "i16x8.sub_saturate_s"
}

func (self *I16x8SubSaturateS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x6c}
	sink.WriteBytes(inst)
	
}

type I16x8SubSaturateU struct {
	
}

func (self *I16x8SubSaturateU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8SubSaturateU) String() string {
	return "i16x8.sub_saturate_u"
}

func (self *I16x8SubSaturateU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x6d}
	sink.WriteBytes(inst)
	
}

type I16x8Mul struct {
	
}

func (self *I16x8Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8Mul) String() string {
	return "i16x8.mul"
}

func (self *I16x8Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x6e}
	sink.WriteBytes(inst)
	
}

type I32x4Neg struct {
	
}

func (self *I32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Neg) String() string {
	return "i32x4.neg"
}

func (self *I32x4Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x73}
	sink.WriteBytes(inst)
	
}

type I32x4AnyTrue struct {
	
}

func (self *I32x4AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4AnyTrue) String() string {
	return "i32x4.any_true"
}

func (self *I32x4AnyTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x74}
	sink.WriteBytes(inst)
	
}

type I32x4AllTrue struct {
	
}

func (self *I32x4AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4AllTrue) String() string {
	return "i32x4.all_true"
}

func (self *I32x4AllTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x75}
	sink.WriteBytes(inst)
	
}

type I32x4Shl struct {
	
}

func (self *I32x4Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Shl) String() string {
	return "i32x4.shl"
}

func (self *I32x4Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x76}
	sink.WriteBytes(inst)
	
}

type I32x4ShrS struct {
	
}

func (self *I32x4ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4ShrS) String() string {
	return "i32x4.shr_s"
}

func (self *I32x4ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x77}
	sink.WriteBytes(inst)
	
}

type I32x4ShrU struct {
	
}

func (self *I32x4ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4ShrU) String() string {
	return "i32x4.shr_u"
}

func (self *I32x4ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x78}
	sink.WriteBytes(inst)
	
}

type I32x4Add struct {
	
}

func (self *I32x4Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Add) String() string {
	return "i32x4.add"
}

func (self *I32x4Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x79}
	sink.WriteBytes(inst)
	
}

type I32x4Sub struct {
	
}

func (self *I32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Sub) String() string {
	return "i32x4.sub"
}

func (self *I32x4Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x7c}
	sink.WriteBytes(inst)
	
}

type I32x4Mul struct {
	
}

func (self *I32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4Mul) String() string {
	return "i32x4.mul"
}

func (self *I32x4Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x7f}
	sink.WriteBytes(inst)
	
}

type I64x2Neg struct {
	
}

func (self *I64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2Neg) String() string {
	return "i64x2.neg"
}

func (self *I64x2Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x84}
	sink.WriteBytes(inst)
	
}

type I64x2AnyTrue struct {
	
}

func (self *I64x2AnyTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2AnyTrue) String() string {
	return "i64x2.any_true"
}

func (self *I64x2AnyTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x85}
	sink.WriteBytes(inst)
	
}

type I64x2AllTrue struct {
	
}

func (self *I64x2AllTrue) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2AllTrue) String() string {
	return "i64x2.all_true"
}

func (self *I64x2AllTrue) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x86}
	sink.WriteBytes(inst)
	
}

type I64x2Shl struct {
	
}

func (self *I64x2Shl) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2Shl) String() string {
	return "i64x2.shl"
}

func (self *I64x2Shl) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x87}
	sink.WriteBytes(inst)
	
}

type I64x2ShrS struct {
	
}

func (self *I64x2ShrS) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2ShrS) String() string {
	return "i64x2.shr_s"
}

func (self *I64x2ShrS) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x88}
	sink.WriteBytes(inst)
	
}

type I64x2ShrU struct {
	
}

func (self *I64x2ShrU) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2ShrU) String() string {
	return "i64x2.shr_u"
}

func (self *I64x2ShrU) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x89}
	sink.WriteBytes(inst)
	
}

type I64x2Add struct {
	
}

func (self *I64x2Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2Add) String() string {
	return "i64x2.add"
}

func (self *I64x2Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x8a}
	sink.WriteBytes(inst)
	
}

type I64x2Sub struct {
	
}

func (self *I64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2Sub) String() string {
	return "i64x2.sub"
}

func (self *I64x2Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x8d}
	sink.WriteBytes(inst)
	
}

type I64x2Mul struct {
	
}

func (self *I64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2Mul) String() string {
	return "i64x2.mul"
}

func (self *I64x2Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x90}
	sink.WriteBytes(inst)
	
}

type F32x4Abs struct {
	
}

func (self *F32x4Abs) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Abs) String() string {
	return "f32x4.abs"
}

func (self *F32x4Abs) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x95}
	sink.WriteBytes(inst)
	
}

type F32x4Neg struct {
	
}

func (self *F32x4Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Neg) String() string {
	return "f32x4.neg"
}

func (self *F32x4Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x96}
	sink.WriteBytes(inst)
	
}

type F32x4Sqrt struct {
	
}

func (self *F32x4Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Sqrt) String() string {
	return "f32x4.sqrt"
}

func (self *F32x4Sqrt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x97}
	sink.WriteBytes(inst)
	
}

type F32x4Add struct {
	
}

func (self *F32x4Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Add) String() string {
	return "f32x4.add"
}

func (self *F32x4Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9a}
	sink.WriteBytes(inst)
	
}

type F32x4Sub struct {
	
}

func (self *F32x4Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Sub) String() string {
	return "f32x4.sub"
}

func (self *F32x4Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9b}
	sink.WriteBytes(inst)
	
}

type F32x4Mul struct {
	
}

func (self *F32x4Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Mul) String() string {
	return "f32x4.mul"
}

func (self *F32x4Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9c}
	sink.WriteBytes(inst)
	
}

type F32x4Div struct {
	
}

func (self *F32x4Div) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Div) String() string {
	return "f32x4.div"
}

func (self *F32x4Div) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9d}
	sink.WriteBytes(inst)
	
}

type F32x4Min struct {
	
}

func (self *F32x4Min) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Min) String() string {
	return "f32x4.min"
}

func (self *F32x4Min) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9e}
	sink.WriteBytes(inst)
	
}

type F32x4Max struct {
	
}

func (self *F32x4Max) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4Max) String() string {
	return "f32x4.max"
}

func (self *F32x4Max) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0x9f}
	sink.WriteBytes(inst)
	
}

type F64x2Abs struct {
	
}

func (self *F64x2Abs) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Abs) String() string {
	return "f64x2.abs"
}

func (self *F64x2Abs) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa0}
	sink.WriteBytes(inst)
	
}

type F64x2Neg struct {
	
}

func (self *F64x2Neg) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Neg) String() string {
	return "f64x2.neg"
}

func (self *F64x2Neg) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa1}
	sink.WriteBytes(inst)
	
}

type F64x2Sqrt struct {
	
}

func (self *F64x2Sqrt) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Sqrt) String() string {
	return "f64x2.sqrt"
}

func (self *F64x2Sqrt) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa2}
	sink.WriteBytes(inst)
	
}

type F64x2Add struct {
	
}

func (self *F64x2Add) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Add) String() string {
	return "f64x2.add"
}

func (self *F64x2Add) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa5}
	sink.WriteBytes(inst)
	
}

type F64x2Sub struct {
	
}

func (self *F64x2Sub) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Sub) String() string {
	return "f64x2.sub"
}

func (self *F64x2Sub) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa6}
	sink.WriteBytes(inst)
	
}

type F64x2Mul struct {
	
}

func (self *F64x2Mul) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Mul) String() string {
	return "f64x2.mul"
}

func (self *F64x2Mul) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa7}
	sink.WriteBytes(inst)
	
}

type F64x2Div struct {
	
}

func (self *F64x2Div) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Div) String() string {
	return "f64x2.div"
}

func (self *F64x2Div) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa8}
	sink.WriteBytes(inst)
	
}

type F64x2Min struct {
	
}

func (self *F64x2Min) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Min) String() string {
	return "f64x2.min"
}

func (self *F64x2Min) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xa9}
	sink.WriteBytes(inst)
	
}

type F64x2Max struct {
	
}

func (self *F64x2Max) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2Max) String() string {
	return "f64x2.max"
}

func (self *F64x2Max) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xaa}
	sink.WriteBytes(inst)
	
}

type I32x4TruncSatF32x4S struct {
	
}

func (self *I32x4TruncSatF32x4S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4TruncSatF32x4S) String() string {
	return "i32x4.trunc_sat_f32x4_s"
}

func (self *I32x4TruncSatF32x4S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xab}
	sink.WriteBytes(inst)
	
}

type I32x4TruncSatF32x4U struct {
	
}

func (self *I32x4TruncSatF32x4U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4TruncSatF32x4U) String() string {
	return "i32x4.trunc_sat_f32x4_u"
}

func (self *I32x4TruncSatF32x4U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xac}
	sink.WriteBytes(inst)
	
}

type I64x2TruncSatF64x2S struct {
	
}

func (self *I64x2TruncSatF64x2S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2TruncSatF64x2S) String() string {
	return "i64x2.trunc_sat_f64x2_s"
}

func (self *I64x2TruncSatF64x2S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xad}
	sink.WriteBytes(inst)
	
}

type I64x2TruncSatF64x2U struct {
	
}

func (self *I64x2TruncSatF64x2U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I64x2TruncSatF64x2U) String() string {
	return "i64x2.trunc_sat_f64x2_u"
}

func (self *I64x2TruncSatF64x2U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xae}
	sink.WriteBytes(inst)
	
}

type F32x4ConvertI32x4S struct {
	
}

func (self *F32x4ConvertI32x4S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4ConvertI32x4S) String() string {
	return "f32x4.convert_i32x4_s"
}

func (self *F32x4ConvertI32x4S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xaf}
	sink.WriteBytes(inst)
	
}

type F32x4ConvertI32x4U struct {
	
}

func (self *F32x4ConvertI32x4U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F32x4ConvertI32x4U) String() string {
	return "f32x4.convert_i32x4_u"
}

func (self *F32x4ConvertI32x4U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xb0}
	sink.WriteBytes(inst)
	
}

type F64x2ConvertI64x2S struct {
	
}

func (self *F64x2ConvertI64x2S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2ConvertI64x2S) String() string {
	return "f64x2.convert_i64x2_s"
}

func (self *F64x2ConvertI64x2S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xb1}
	sink.WriteBytes(inst)
	
}

type F64x2ConvertI64x2U struct {
	
}

func (self *F64x2ConvertI64x2U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *F64x2ConvertI64x2U) String() string {
	return "f64x2.convert_i64x2_u"
}

func (self *F64x2ConvertI64x2U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xb2}
	sink.WriteBytes(inst)
	
}

type V8x16Swizzle struct {
	
}

func (self *V8x16Swizzle) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V8x16Swizzle) String() string {
	return "v8x16.swizzle"
}

func (self *V8x16Swizzle) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc0}
	sink.WriteBytes(inst)
	
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

func (self *V8x16LoadSplat) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc2}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *V16x8LoadSplat) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc3}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *V32x4LoadSplat) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc4}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *V64x2LoadSplat) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc5}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

}

type I8x16NarrowI16x8S struct {
	
}

func (self *I8x16NarrowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16NarrowI16x8S) String() string {
	return "i8x16.narrow_i16x8_s"
}

func (self *I8x16NarrowI16x8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc6}
	sink.WriteBytes(inst)
	
}

type I8x16NarrowI16x8U struct {
	
}

func (self *I8x16NarrowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I8x16NarrowI16x8U) String() string {
	return "i8x16.narrow_i16x8_u"
}

func (self *I8x16NarrowI16x8U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc7}
	sink.WriteBytes(inst)
	
}

type I16x8NarrowI32x4S struct {
	
}

func (self *I16x8NarrowI32x4S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8NarrowI32x4S) String() string {
	return "i16x8.narrow_i32x4_s"
}

func (self *I16x8NarrowI32x4S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc8}
	sink.WriteBytes(inst)
	
}

type I16x8NarrowI32x4U struct {
	
}

func (self *I16x8NarrowI32x4U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8NarrowI32x4U) String() string {
	return "i16x8.narrow_i32x4_u"
}

func (self *I16x8NarrowI32x4U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xc9}
	sink.WriteBytes(inst)
	
}

type I16x8WidenLowI8x16S struct {
	
}

func (self *I16x8WidenLowI8x16S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8WidenLowI8x16S) String() string {
	return "i16x8.widen_low_i8x16_s"
}

func (self *I16x8WidenLowI8x16S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xca}
	sink.WriteBytes(inst)
	
}

type I16x8WidenHighI8x16S struct {
	
}

func (self *I16x8WidenHighI8x16S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8WidenHighI8x16S) String() string {
	return "i16x8.widen_high_i8x16_s"
}

func (self *I16x8WidenHighI8x16S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xcb}
	sink.WriteBytes(inst)
	
}

type I16x8WidenLowI8x16U struct {
	
}

func (self *I16x8WidenLowI8x16U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8WidenLowI8x16U) String() string {
	return "i16x8.widen_low_i8x16_u"
}

func (self *I16x8WidenLowI8x16U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xcc}
	sink.WriteBytes(inst)
	
}

type I16x8WidenHighI8x16u struct {
	
}

func (self *I16x8WidenHighI8x16u) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I16x8WidenHighI8x16u) String() string {
	return "i16x8.widen_high_i8x16_u"
}

func (self *I16x8WidenHighI8x16u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xcd}
	sink.WriteBytes(inst)
	
}

type I32x4WidenLowI16x8S struct {
	
}

func (self *I32x4WidenLowI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4WidenLowI16x8S) String() string {
	return "i32x4.widen_low_i16x8_s"
}

func (self *I32x4WidenLowI16x8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xce}
	sink.WriteBytes(inst)
	
}

type I32x4WidenHighI16x8S struct {
	
}

func (self *I32x4WidenHighI16x8S) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4WidenHighI16x8S) String() string {
	return "i32x4.widen_high_i16x8_s"
}

func (self *I32x4WidenHighI16x8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xcf}
	sink.WriteBytes(inst)
	
}

type I32x4WidenLowI16x8U struct {
	
}

func (self *I32x4WidenLowI16x8U) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4WidenLowI16x8U) String() string {
	return "i32x4.widen_low_i16x8_u"
}

func (self *I32x4WidenLowI16x8U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd0}
	sink.WriteBytes(inst)
	
}

type I32x4WidenHighI16x8u struct {
	
}

func (self *I32x4WidenHighI16x8u) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *I32x4WidenHighI16x8u) String() string {
	return "i32x4.widen_high_i16x8_u"
}

func (self *I32x4WidenHighI16x8u) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd1}
	sink.WriteBytes(inst)
	
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

func (self *I16x8Load8x8S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd2}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I16x8Load8x8U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd3}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32x4Load16x4S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd4}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I32x4Load16x4U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd5}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64x2Load32x2S) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd6}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

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

func (self *I64x2Load32x2U) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd7}
	sink.WriteBytes(inst)
	self.MemArg.Encode(sink)

}

type V128Andnot struct {
	
}

func (self *V128Andnot) parseInstrBody(ps *parser.ParserBuffer) error {
	
	return nil
}

func (self *V128Andnot) String() string {
	return "v128.andnot"
}

func (self *V128Andnot) Encode(sink *ZeroCopySink) {
	inst := []byte{0xfd,0xd8}
	sink.WriteBytes(inst)
	
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

