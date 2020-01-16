package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
	"github.com/valyala/fasttemplate"
)

type Instruction struct {
	Name   string
	Id     []string
	Fields []Field
	Inst   []byte
}

type Field struct {
	Name string
	Type string
}

func expectKeyword(ps *parser.ParserBuffer) (string, error) {
	token := ps.ReadToken()
	if token == nil {
		return "", errors.New("expect keyword")
	}
	switch val := token.(type) {
	case lexer.Keyword:
		return val.Val, nil
	case lexer.Reserved:
		return val.Val, nil
	default:
		return "", errors.New("expect keyword")
	}
}

func (self *Instruction) Parse(ps *parser.ParserBuffer) error {
	kw, err := expectKeyword(ps)
	if err != nil {
		return err
	}
	self.Name = kw
	if err := ps.Parens(func(ps *parser.ParserBuffer) error {
		for !ps.Empty() {
			v, err := ps.ExpectUint32()
			if err != nil {
				return err
			}
			if v > uint32(0xff) {
				return fmt.Errorf("err Instruction encode %d", v)
			}

			self.Inst = append(self.Inst, byte(v))
		}
		return nil
	}); err != nil {
		return err
	}

	if ps.PeekToken().Type() == lexer.LParenType {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			for !ps.Empty() {
				id, err := expectKeyword(ps)
				if err != nil {
					return err
				}
				self.Id = append(self.Id, id)
			}
			return nil
		})

		if err != nil {
			return err
		}
	} else {
		id, err := expectKeyword(ps)
		if err != nil {
			return err
		}
		self.Id = append(self.Id, id)
	}

	for !ps.Empty() {
		err = ps.Parens(func(ps *parser.ParserBuffer) error {
			field, err := expectKeyword(ps)
			if err != nil {
				return err
			}
			ty, err := expectKeyword(ps)
			if err != nil {
				return err
			}

			self.Fields = append(self.Fields, Field{Name: field, Type: ty})
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func generate(template string, m map[string]interface{}) string {
	t := fasttemplate.New(template, "[", "]")
	return t.ExecuteString(m)
}

func (self Instruction) Generate() string {
	template := `
type [Name] struct {
	[Fields]
}

func (self *[Name]) parseInstrBody(ps *parser.ParserBuffer) error {
	[parseBody]
	return nil
}

func (self *[Name]) String() string {
	return "[Id]"
}

func (self *[Name]) Encode(sink *ZeroCopySink) {
	inst := [Instruction]
	sink.WriteBytes(inst)
	[FieldsEncode]
}
`
	return generate(template, map[string]interface{}{
		"Name":         self.Name,
		"Fields":       self.generateFields(),
		"parseBody":    self.generateParseBody(),
		"Id":           self.Id[0],
		"Instruction":  self.generateInstr(),
		"FieldsEncode": self.generateEncode(),
	})
}

func (self Instruction) generateEncode() string {
	fieldsEncode := ""
	for _, field := range self.Fields {
		switch field.Type {
		case "uint32":
			fieldsEncode += "sink.WriteUint32(self." + field.Name + ")" + "\n"
		case "int64":
			fieldsEncode += "sink.WriteInt64(self." + field.Name + ")" + "\n"
		default:
			fieldsEncode += "self." + field.Name + ".Encode()" + "\n"
		}
	}

	return fieldsEncode
}

func (self Instruction) generateInstr() string {
	var instr []string
	for _, b := range self.Inst {
		instr = append(instr, fmt.Sprintf("0x%x", b))
	}

	s := strings.Join(instr, ",")
	return "[]byte{" + s + "}"
}

func (self Instruction) generateFields() string {
	var fields []string
	for _, field := range self.Fields {
		ty := field.Type
		if strings.HasPrefix(ty, "MemArg") {
			ty = "MemArg"
		}
		fields = append(fields, fmt.Sprintf("%s %s", field.Name, ty))
	}

	return strings.Join(fields, "\n")
}

func (self Instruction) generateParseBody() string {
	body := ""
	for _, field := range self.Fields {
		switch field.Type {
		case "uint32":
			body += parseInt(field.Name, "Uint32")
		case "int64":
			body += parseInt(field.Name, "Int64")
		case "OptionId":
			body += parseOptionId(field.Name)
		default:
			if strings.HasPrefix(field.Type, "MemArg") {
				body += parseMemArg(field.Name, field.Type)
			} else {
				body += parseGeneral(field.Name)
			}
		}
	}

	return body
}

func parseMemArg(name string, ty string) string {
	offset := strings.Trim(ty, "MemArg<>")

	return generate(
		`err := self.[Name].Parse(ps, [offset])
	if err != nil {
		return err
	}
`, map[string]interface{}{"Name": name, "offset": offset})
}

func parseGeneral(name string) string {
	return generate(
		`err := self.[Name].Parse(ps)
	if err != nil {
		return err
	}
`, map[string]interface{}{"Name": name})
}

func parseOptionId(name string) string {
	return generate(`self.[Name].Parse(ps)
`, map[string]interface{}{"Name": name})
}

func parseInt(name string, ty string) string {
	return generate(`val, err := ps.Expect[Type]()
	if err != nil {
		return err
	}
	self.[Name] = val
`, map[string]interface{}{"Name": name, "Type": ty})
}

func mustParseInstrs(source string) []Instruction {
	ps, err := parser.NewParserBuffer(source)
	if err != nil {
		panic(err)
	}
	var instrs []Instruction
	for !ps.Empty() {
		var inst Instruction
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			return inst.Parse(ps)
		})
		if err != nil {
			panic(err)
		}
		instrs = append(instrs, inst)
	}

	return instrs
}

func generateParseInstrution(instrs []Instruction) string {
	var cases []string
	for _, instr := range instrs {
		cases = append(cases, generate(` case "[Id]":
		inst = &[Name]{}`, map[string]interface{}{"Id": strings.Join(instr.Id, "\", \""), "Name": instr.Name}))
	}

	return generate(`
func parseInstr(ps *parser.ParserBuffer) (Instruction, error) {
	var inst Instruction
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	[cases]
	default:
		panic(fmt.Sprintf("todo: implement instruction %s", kw))
	}
	err = inst.parseInstrBody(ps)
	if err != nil {
		return nil, err
	}
	return inst, nil
}
`, map[string]interface{}{"cases": strings.Join(cases, "\n")})
}

func main() {
	instrs := `
(Block (0x02) block (BlockType BlockType))
(If (0x04) if (BlockType BlockType))
(Else (0x05) else (Id OptionId))
(Loop (0x03) loop (BlockType BlockType))
(End (0x0b) end (Id OptionId))

(Unreachable (0x00) unreachable)
(Nop (0x01) nop)
(Br (0x0c) br (Index Index))
(BrIf (0x0d) br_if (Index Index))
(BrTable (0x0e) br_table (Indices BrTableIndices)) 
(Return (0x0f) return)
(Call (0x10) call (Index Index))
(CallIndirect (0x11) call_indirect (Impl CallIndirectInner))
(ReturnCall (0x12) return_call (Index Index))
(ReturnCallIndirect (0x13) return_call_indirect (Impl CallIndirectInner))
(Drop (0x1a) drop)
(Select () select (SelectTypes SelectTypes))
(LocalGet (0x20) (local.get get_local) (Index Index)) 
(LocalSet (0x21) (local.set set_local) (Index Index)) 
(LocalTee (0x22) (local.tee tee_local) (Index Index)) 
(GlobalGet (0x23) (global.get get_global) (Index Index)) 
(GlobalSet (0x24) (global.set set_global) (Index Index)) 

(TableGet (0x25) table.get (Index Index))
(TableSet (0x26) table.set (Index Index))

(I32Load (0x28) i32.load (MemArg MemArg<4>))
(I64Load (0x29) i64.load (MemArg MemArg<8>))
(F32Load (0x2a) f32.load (MemArg MemArg<4>))
(F64Load (0x2b) f64.load (MemArg MemArg<8>))
(I32Load8s (0x2c) i32.load8_s (MemArg MemArg<1>))
(I32Load8u (0x2d) i32.load8_u (MemArg MemArg<1>))
(I32Load16s (0x2e) i32.load16_s (MemArg MemArg<2>))
(I32Load16u (0x2f) i32.load16_u (MemArg MemArg<2>))
(I64Load8s (0x30) i64.load8_s (MemArg MemArg<1>))
(I64Load8u (0x31) i64.load8_u (MemArg MemArg<1>))
(I64Load16s (0x32) i64.load16_s (MemArg MemArg<2>))
(I64Load16u (0x33) i64.load16_u (MemArg MemArg<2>))
(I64Load32s (0x34) i64.load32_s (MemArg MemArg<4>))
(I64Load32u (0x35) i64.load32_u (MemArg MemArg<4>))
(I32Store (0x36) i32.store (MemArg MemArg<4>))
(I64Store (0x37) i64.store (MemArg MemArg<8>))
(F32Store (0x38) f32.store (MemArg MemArg<4>))
(F64Store (0x39) f64.store (MemArg MemArg<8>))
(I32Store8 (0x3a) i32.store8 (MemArg MemArg<1>))
(I32Store16 (0x3b) i32.store16 (MemArg MemArg<2>))
(I64Store8 (0x3c) i64.store8 (MemArg MemArg<1>))
(I64Store16 (0x3d) i64.store16 (MemArg MemArg<2>))
(I64Store32 (0x3e) i64.store32 (MemArg MemArg<4>))

;; Lots of bulk memory proposal here as well
(MemorySize (0x3f 0x00) (memory.size current_memory))
(MemoryGrow (0x40 0x00) (memory.grow grow_memory))
;;MemoryInit(MemoryInit<'a>) : [0xfc 0x08] : "memory.init"
(MemoryCopy (0xfc 0x0a 0x00 0x00) memory.copy)
(MemoryFill (0xfc 0x0b 0x00) memory.fill)
(DataDrop (0xfc 0x09) data.drop (Index Index))
(ElemDrop (0xfc 0x0d) elem.drop (Index Index))
;;(TableInit table.init (Impl TableInit))
(TableCopy (0xfc 0x0e 0x00 0x00) table.copy)
(TableFill (0xfc 0x11) table.fill (Index Index))
(TableSize (0xfc 0x10) table.size (Index Index))
(TableGrow (0xfc 0x0f) table.grow (Index Index))

(RefNull (0xd0) ref.null)
(RefIsNull (0xd1) ref.is_null)
(RefHost (0xff) ref.host (Val uint32))
(RefFunc (0xd2) ref.func (Index Index))

(I32Const (0x41) i32.const (Val uint32))
(I64Const (0x42) i64.const (Val int64))
(F32Const (0x43) f32.const (Val Float32))
(F64Const (0x44) f64.const (Val Float64))

(I32Clz (0x67) i32.clz)
(I32Ctz (0x68) i32.ctz)
(I32Pocnt (0x69) i32.popcnt)
(I32Add (0x6a) i32.add)
(I32Sub (0x6b) i32.sub)
(I32Mul (0x6c) i32.mul)
(I32DivS (0x6d) i32.div_s)
(I32DivU (0x6e) i32.div_u)
(I32RemS (0x6f) i32.rem_s)
(I32RemU (0x70) i32.rem_u)
(I32And (0x71) i32.and)
(I32Or (0x72) i32.or)
(I32Xor (0x73) i32.xor)
(I32Shl (0x74) i32.shl)
(I32ShrS (0x75) i32.shr_s)
(I32ShrU (0x76) i32.shr_u)
(I32Rotl (0x77) i32.rotl)
(I32Rotr (0x78) i32.rotr)

(I64Clz (0x79) i64.clz)
(I64Ctz (0x7a) i64.ctz)
(I64Popcnt (0x7b) i64.popcnt)
(I64Add (0x7c) i64.add)
(I64Sub (0x7d) i64.sub)
(I64Mul (0x7e) i64.mul)
(I64DivS (0x7f) i64.div_s)
(I64DivU (0x80) i64.div_u)
(I64RemS (0x81) i64.rem_s)
(I64RemU (0x82) i64.rem_u)
(I64And (0x83) i64.and)
(I64Or (0x84) i64.or)
(I64Xor (0x85) i64.xor)
(I64Shl (0x86) i64.shl)
(I64ShrS (0x87) i64.shr_s)
(I64ShrU (0x88) i64.shr_u)
(I64Rotl (0x89) i64.rotl)
(I64Rotr (0x8a) i64.rotr)

(F32Abs (0x8b) f32.abs)
(F32Neg (0x8c) f32.neg)
(F32Ceil (0x8d) f32.ceil)
(F32Floor (0x8e) f32.floor)
(F32Trunc (0x8f) f32.trunc)
(F32Nearest (0x90) f32.nearest)
(F32Sqrt (0x91) f32.sqrt)
(F32Add (0x92) f32.add)
(F32Sub (0x93) f32.sub)
(F32Mul (0x94) f32.mul)
(F32Div (0x95) f32.div)
(F32Min (0x96) f32.min)
(F32Max (0x97) f32.max)
(F32Copysign (0x98) f32.copysign)

(F64Abs (0x99) f64.abs)
(F64Neg (0x9a) f64.neg)
(F64Ceil (0x9b) f64.ceil)
(F64Floor (0x9c) f64.floor)
(F64Trunc (0x9d) f64.trunc)
(F64Nearest (0x9e) f64.nearest)
(F64Sqrt (0x9f) f64.sqrt)
(F64Add (0xa0) f64.add)
(F64Sub (0xa1) f64.sub)
(F64Mul (0xa2) f64.mul)
(F64Div (0xa3) f64.div)
(F64Min (0xa4) f64.min)
(F64Max (0xa5) f64.max)
(F64Copysign (0xa6) f64.copysign)

(I32Eqz (0x45) i32.eqz)
(I32Eq (0x46) i32.eq)
(I32Ne (0x47) i32.ne)
(I32LtS (0x48) i32.lt_s)
(I32LtU (0x49) i32.lt_u)
(I32GtS (0x4a) i32.gt_s)
(I32GtU (0x4b) i32.gt_u)
(I32LeS (0x4c) i32.le_s)
(I32LeU (0x4d) i32.le_u)
(I32GeS (0x4e) i32.ge_s)
(I32GeU (0x4f) i32.ge_u)

(I64Eqz (0x50) i64.eqz)
(I64Eq (0x51) i64.eq)
(I64Ne (0x52) i64.ne)
(I64LtS (0x53) i64.lt_s)
(I64LtU (0x54) i64.lt_u)
(I64GtS (0x55) i64.gt_s)
(I64GtU (0x56) i64.gt_u)
(I64LeS (0x57) i64.le_s)
(I64LeU (0x58) i64.le_u)
(I64GeS (0x59) i64.ge_s)
(I64GeU (0x5a) i64.ge_u)

(F32Eq (0x5b) f32.eq)
(F32Ne (0x5c) f32.ne)
(F32Lt (0x5d) f32.lt)
(F32Gt (0x5e) f32.gt)
(F32Le (0x5f) f32.le)
(F32Ge (0x60) f32.ge)

(F64Eq (0x61) f64.eq)
(F64Ne (0x62) f64.ne)
(F64Lt (0x63) f64.lt)
(F64Gt (0x64) f64.gt)
(F64Le (0x65) f64.le)
(F64Ge (0x66) f64.ge)

(I32WrapI64 (0xa7) (i32.wrap_i64 i32.wrap/i64))
(I32TruncF32S (0xa8) (i32.trunc_f32_s i32.trunc_s/f32))
(I32TruncF32U (0xa9) (i32.trunc_f32_u i32.trunc_u/f32))
(I32TruncF64S (0xaa) (i32.trunc_f64_s i32.trunc_s/f64))
(I32TruncF64U (0xab) (i32.trunc_f64_u i32.trunc_u/f64))
(I64ExtendI32S (0xac) (i64.extend_i32_s i64.extend_s/i32))
(I64ExtendI32U (0xad) (i64.extend_i32_u i64.extend_u/i32))
(I64TruncF32S (0xae) (i64.trunc_f32_s i64.trunc_s/f32))
(I64TruncF32U (0xaf) (i64.trunc_f32_u i64.trunc_u/f32))
(I64TruncF64S (0xb0) (i64.trunc_f64_s i64.trunc_s/f64))
(I64TruncF64U (0xb1) (i64.trunc_f64_u i64.trunc_u/f64))
(F32ConvertI32S (0xb2) (f32.convert_i32_s f32.convert_s/i32))
(F32ConvertI32U (0xb3) (f32.convert_i32_u f32.convert_u/i32))
(F32ConvertI64S (0xb4) (f32.convert_i64_s f32.convert_s/i64))
(F32ConvertI64U (0xb5) (f32.convert_i64.u f32.convert_u/i64))
(F32DemoteF64 (0xb6) (f32.demote_f64 f32.demote/f64))
(F64ConvertI32S (0xb7) (f64.convert_i32_s f64.convert_s/i32))
(F64ConvertI32U (0xb8) (f64.convert_i32_u f64.convert_u/i32))
(F64ConvertI64S (0xb9) (f64.convert_i64_s f64.convert_s/i64))
(F64ConvertI64U (0xba) (f64.convert_i64_u f64.convert_u/i64))
(F64PromoteF32 (0xbb) (f64.promote_f32 f64.promote/f32))
(I32ReinterpretF32 (0xbc) (i32.reinterpret_f32 i32.reinterpret/f32))
(I64ReinterpretF64 (0xbd) (i64.reinterpret_f64 i64.reinterpret/f64))
(F32ReinterpretI32 (0xbe) (f32.reinterpret_i32 f32.reinterpret/i32))
(F64ReinterpretI64 (0xbf) (f64.reinterpret_i64 f64.reinterpret/i64))

;;(// non-trapping float to int
(I32TruncSatF32S (0xfc 0x00) (i32.trunc_sat_f32_s i32.trunc_s:sat/f32))
(I32TruncSatF32U (0xfc 0x01) (i32.trunc_sat_f32_u i32.trunc_u:sat/f32))
(I32TruncSatF64S (0xfc 0x02) (i32.trunc_sat_f64_s i32.trunc_s:sat/f64))
(I32TruncSatF64U (0xfc 0x03) (i32.trunc_sat_f64_u i32.trunc_u:sat/f64))
(I64TruncSatF32S (0xfc 0x04) (i64.trunc_sat_f32_s i64.trunc_s:sat/f32))
(I64TruncSatF32U (0xfc 0x05) (i64.trunc_sat_f32_u i64.trunc_u:sat/f32))
(I64TruncSatF64S (0xfc 0x06) (i64.trunc_sat_f64_s i64.trunc_s:sat/f64))
(I64TruncSatF64U (0xfc 0x07) (i64.trunc_sat_f64_u i64.trunc_u:sat/f64))

;; sign extension proposal
(I32Extend8S (0xc0) i32.extend8_s)
(I32Extend16S (0xc1) i32.extend16_s)
(I64Extend8S (0xc2) i64.extend8_s)
(I64Extend16S (0xc3) i64.extend16_s)
(I64Extend32S (0xc4) i64.extend32_s)

;; atomics proposal
(AtomicNotify (0xfe 0x00) atomic.notify (MemArg MemArg<4>))
(I32AtomicWait (0xfe 0x01) i32.atomic.wait (MemArg MemArg<4>))
(I64AtomicWait (0xfe 0x02) i64.atomic.wait (MemArg MemArg<8>))
(AtomicFence (0xfe 0x03) atomic.fence)

(I32AtomicLoad (0xfe 0x10) i32.atomic.load (MemArg MemArg<4>))
(I64AtomicLoad (0xfe 0x11) i64.atomic.load (MemArg MemArg<8>))
(I32AtomicLoad8u (0xfe 0x12) i32.atomic.load8_u (MemArg MemArg<1>))
(I32AtomicLoad16u (0xfe 0x13) i32.atomic.load16_u (MemArg MemArg<2>))
(I64AtomicLoad8u (0xfe 0x14) i64.atomic.load8_u (MemArg MemArg<1>))
(I64AtomicLoad16u (0xfe 0x15) i64.atomic.load16_u (MemArg MemArg<2>))
(I64AtomicLoad32u (0xfe 0x16) i64.atomic.load32_u (MemArg MemArg<4>))
(I32AtomicStore (0xfe 0x17) i32.atomic.store (MemArg MemArg<4>))
(I64AtomicStore (0xfe 0x18) i64.atomic.store (MemArg MemArg<8>))
(I32AtomicStore8 (0xfe 0x19) i32.atomic.store8 (MemArg MemArg<1>))
(I32AtomicStore16 (0xfe 0x1a) i32.atomic.store16 (MemArg MemArg<2>))
(I64AtomicStore8 (0xfe 0x1b) i64.atomic.store8 (MemArg MemArg<1>))
(I64AtomicStore16 (0xfe 0x1c) i64.atomic.store16 (MemArg MemArg<2>))
(I64AtomicStore32 (0xfe 0x1d) i64.atomic.store32 (MemArg MemArg<4>))

(I32AtomicRmwAdd (0xfe 0x1e) i32.atomic.rmw.add (MemArg MemArg<4>))
(I64AtomicRmwAdd (0xfe 0x1f) i64.atomic.rmw.add (MemArg MemArg<8>))
(I32AtomicRmw8AddU (0xfe 0x20) i32.atomic.rmw8.add_u (MemArg MemArg<1>))
(I32AtomicRmw16AddU (0xfe 0x21) i32.atomic.rmw16.add_u (MemArg MemArg<2>))
(I64AtomicRmw8AddU (0xfe 0x22) i64.atomic.rmw8.add_u (MemArg MemArg<1>))
(I64AtomicRmw16AddU (0xfe 0x23) i64.atomic.rmw16.add_u (MemArg MemArg<2>))
(I64AtomicRmw32AddU (0xfe 0x24) i64.atomic.rmw32.add_u (MemArg MemArg<4>))

(I32AtomicRmwSub (0xfe 0x25) i32.atomic.rmw.sub (MemArg MemArg<4>))
(I64AtomicRmwSub (0xfe 0x26) i64.atomic.rmw.sub (MemArg MemArg<8>))
(I32AtomicRmw8SubU (0xfe 0x27) i32.atomic.rmw8.sub_u (MemArg MemArg<1>))
(I32AtomicRmw16SubU (0xfe 0x28) i32.atomic.rmw16.sub_u (MemArg MemArg<2>))
(I64AtomicRmw8SubU (0xfe 0x29) i64.atomic.rmw8.sub_u (MemArg MemArg<1>))
(I64AtomicRmw16SubU (0xfe 0x2a) i64.atomic.rmw16.sub_u (MemArg MemArg<2>))
(I64AtomicRmw32SubU (0xfe 0x2b) i64.atomic.rmw32.sub_u (MemArg MemArg<4>))

(I32AtomicRmwAnd (0xfe 0x2c) i32.atomic.rmw.and (MemArg MemArg<4>))
(I64AtomicRmwAnd (0xfe 0x2d) i64.atomic.rmw.and (MemArg MemArg<8>))
(I32AtomicRmw8AndU (0xfe 0x2e) i32.atomic.rmw8.and_u (MemArg MemArg<1>))
(I32AtomicRmw16AndU (0xfe 0x2f) i32.atomic.rmw16.and_u (MemArg MemArg<2>))
(I64AtomicRmw8AndU (0xfe 0x30) i64.atomic.rmw8.and_u (MemArg MemArg<1>))
(I64AtomicRmw16AndU (0xfe 0x31) i64.atomic.rmw16.and_u (MemArg MemArg<2>))
(I64AtomicRmw32AndU (0xfe 0x32) i64.atomic.rmw32.and_u (MemArg MemArg<4>))

(I32AtomicRmwOr (0xfe 0x33) i32.atomic.rmw.or (MemArg MemArg<4>))
(I64AtomicRmwOr (0xfe 0x34) i64.atomic.rmw.or (MemArg MemArg<8>))
(I32AtomicRmw8OrU (0xfe 0x35) i32.atomic.rmw8.or_u (MemArg MemArg<1>))
(I32AtomicRmw16OrU (0xfe 0x36) i32.atomic.rmw16.or_u (MemArg MemArg<2>))
(I64AtomicRmw8OrU (0xfe 0x37) i64.atomic.rmw8.or_u (MemArg MemArg<1>))
(I64AtomicRmw16OrU (0xfe 0x38) i64.atomic.rmw16.or_u (MemArg MemArg<2>))
(I64AtomicRmw32OrU (0xfe 0x39) i64.atomic.rmw32.or_u (MemArg MemArg<4>))

(I32AtomicRmwXor (0xfe 0x3a) i32.atomic.rmw.xor (MemArg MemArg<4>))
(I64AtomicRmwXor (0xfe 0x3b) i64.atomic.rmw.xor (MemArg MemArg<8>))
(I32AtomicRmw8XorU (0xfe 0x3c) i32.atomic.rmw8.xor_u (MemArg MemArg<1>))
(I32AtomicRmw16XorU (0xfe 0x3d) i32.atomic.rmw16.xor_u (MemArg MemArg<2>))
(I64AtomicRmw8XorU (0xfe 0x3e) i64.atomic.rmw8.xor_u (MemArg MemArg<1>))
(I64AtomicRmw16XorU (0xfe 0x3f) i64.atomic.rmw16.xor_u (MemArg MemArg<2>))
(I64AtomicRmw32XorU (0xfe 0x40) i64.atomic.rmw32.xor_u (MemArg MemArg<4>))

(I32AtomicRmwXchg (0xfe 0x41) i32.atomic.rmw.xchg (MemArg MemArg<4>))
(I64AtomicRmwXchg (0xfe 0x42) i64.atomic.rmw.xchg (MemArg MemArg<8>))
(I32AtomicRmw8XchgU (0xfe 0x43) i32.atomic.rmw8.xchg_u (MemArg MemArg<1>))
(I32AtomicRmw16XchgU (0xfe 0x44) i32.atomic.rmw16.xchg_u (MemArg MemArg<2>))
(I64AtomicRmw8XchgU (0xfe 0x45) i64.atomic.rmw8.xchg_u (MemArg MemArg<1>))
(I64AtomicRmw16XchgU (0xfe 0x46) i64.atomic.rmw16.xchg_u (MemArg MemArg<2>))
(I64AtomicRmw32XchgU (0xfe 0x47) i64.atomic.rmw32.xchg_u (MemArg MemArg<4>))

(I32AtomicRmwCmpxchg (0xfe 0x48) i32.atomic.rmw.cmpxchg (MemArg MemArg<4>))
(I64AtomicRmwCmpxchg (0xfe 0x49) i64.atomic.rmw.cmpxchg (MemArg MemArg<8>))
(I32AtomicRmw8CmpxchgU (0xfe 0x4a) i32.atomic.rmw8.cmpxchg_u (MemArg MemArg<1>))
(I32AtomicRmw16CmpxchgU (0xfe 0x4b) i32.atomic.rmw16.cmpxchg_u (MemArg MemArg<2>))
(I64AtomicRmw8CmpxchgU (0xfe 0x4c) i64.atomic.rmw8.cmpxchg_u (MemArg MemArg<1>))
(I64AtomicRmw16CmpxchgU (0xfe 0x4d) i64.atomic.rmw16.cmpxchg_u (MemArg MemArg<2>))
(I64AtomicRmw32CmpxchgU (0xfe 0x4e) i64.atomic.rmw32.cmpxchg_u (MemArg MemArg<4>))

(V128Load (0xfd 0x00) v128.load (MemArg MemArg<16>))
(V128Store (0xfd 0x01) v128.store (MemArg MemArg<16>))
;;(V128Const (0xfd 0x02) v128.const (V128Const V128Const))

;;I8x16Splat : [0xfd 0x04] : "i8x16.splat"
;;I8x16ExtractLaneS(i32) : [0xfd 0x05] : "i8x16.extract_lane_s"
;;I8x16ExtractLaneU(i32) : [0xfd 0x06] : "i8x16.extract_lane_u"
;;I8x16ReplaceLane(i32) : [0xfd 0x07] : "i8x16.replace_lane"
;;I16x8Splat : [0xfd 0x08] : "i16x8.splat"
;;I16x8ExtractLaneS(i32) : [0xfd 0x09] : "i16x8.extract_lane_s"
;;I16x8ExtractLaneU(i32) : [0xfd 0x0a] : "i16x8.extract_lane_u"
;;I16x8ReplaceLane(i32) : [0xfd 0x0b] : "i16x8.replace_lane"
;;I32x4Splat : [0xfd 0x0c] : "i32x4.splat"
;;I32x4ExtractLane(i32) : [0xfd 0x0d] : "i32x4.extract_lane"
;;I32x4ReplaceLane(i32) : [0xfd 0x0e] : "i32x4.replace_lane"
;;I64x2Splat : [0xfd 0x0f] : "i64x2.splat"
;;I64x2ExtractLane(i32) : [0xfd 0x10] : "i64x2.extract_lane"
;;I64x2ReplaceLane(i32) : [0xfd 0x11] : "i64x2.replace_lane"
;;F32x4Splat : [0xfd 0x12] : "f32x4.splat"
;;F32x4ExtractLane(i32) : [0xfd 0x13] : "f32x4.extract_lane"
;;F32x4ReplaceLane(i32) : [0xfd 0x14] : "f32x4.replace_lane"
;;F64x2Splat : [0xfd 0x15] : "f64x2.splat"
;;F64x2ExtractLane(i32) : [0xfd 0x16] : "f64x2.extract_lane"
;;F64x2ReplaceLane(i32) : [0xfd 0x17] : "f64x2.replace_lane"

(I8x16Eq (0xfd 0x18) i8x16.eq)
(I8x16Ne (0xfd 0x19) i8x16.ne)
(I8x16LtS (0xfd 0x1a) i8x16.lt_s)
(I8x16LtU (0xfd 0x1b) i8x16.lt_u)
(I8x16GtS (0xfd 0x1c) i8x16.gt_s)
(I8x16GtU (0xfd 0x1d) i8x16.gt_u)
(I8x16LeS (0xfd 0x1e) i8x16.le_s)
(I8x16LeU (0xfd 0x1f) i8x16.le_u)
(I8x16GeS (0xfd 0x20) i8x16.ge_s)
(I8x16GeU (0xfd 0x21) i8x16.ge_u)
(I16x8Eq (0xfd 0x22) i16x8.eq)
(I16x8Ne (0xfd 0x23) i16x8.ne)
(I16x8LtS (0xfd 0x24) i16x8.lt_s)
(I16x8LtU (0xfd 0x25) i16x8.lt_u)
(I16x8GtS (0xfd 0x26) i16x8.gt_s)
(I16x8GtU (0xfd 0x27) i16x8.gt_u)
(I16x8LeS (0xfd 0x28) i16x8.le_s)
(I16x8LeU (0xfd 0x29) i16x8.le_u)
(I16x8GeS (0xfd 0x2a) i16x8.ge_s)
(I16x8GeU (0xfd 0x2b) i16x8.ge_u)
(I32x4Eq (0xfd 0x2c) i32x4.eq)
(I32x4Ne (0xfd 0x2d) i32x4.ne)
(I32x4LtS (0xfd 0x2e) i32x4.lt_s)
(I32x4LtU (0xfd 0x2f) i32x4.lt_u)
(I32x4GtS (0xfd 0x30) i32x4.gt_s)
(I32x4GtU (0xfd 0x31) i32x4.gt_u)
(I32x4LeS (0xfd 0x32) i32x4.le_s)
(I32x4LeU (0xfd 0x33) i32x4.le_u)
(I32x4GeS (0xfd 0x34) i32x4.ge_s)
(I32x4GeU (0xfd 0x35) i32x4.ge_u)

(F32x4Eq (0xfd 0x40) f32x4.eq)
(F32x4Ne (0xfd 0x41) f32x4.ne)
(F32x4Lt (0xfd 0x42) f32x4.lt)
(F32x4Gt (0xfd 0x43) f32x4.gt)
(F32x4Le (0xfd 0x44) f32x4.le)
(F32x4Ge (0xfd 0x45) f32x4.ge)
(F64x2Eq (0xfd 0x46) f64x2.eq)
(F64x2Ne (0xfd 0x47) f64x2.ne)
(F64x2Lt (0xfd 0x48) f64x2.lt)
(F64x2Gt (0xfd 0x49) f64x2.gt)
(F64x2Le (0xfd 0x4a) f64x2.le)
(F64x2Ge (0xfd 0x4b) f64x2.ge)

(V128Not (0xfd 0x4c) v128.not)
(V128And (0xfd 0x4d) v128.and)
(V128Or (0xfd 0x4e) v128.or)
(V128Xor (0xfd 0x4f) v128.xor)
(V128Bitselect (0xfd 0x50) v128.bitselect)

(I8x16Neg (0xfd 0x51) i8x16.neg)
(I8x16AnyTrue (0xfd 0x52) i8x16.any_true)
(I8x16AllTrue (0xfd 0x53) i8x16.all_true)
(I8x16Shl (0xfd 0x54) i8x16.shl)
(I8x16ShrS (0xfd 0x55) i8x16.shr_s)
(I8x16ShrU (0xfd 0x56) i8x16.shr_u)
(I8x16Add (0xfd 0x57) i8x16.add)
(I8x16AddSaturateS (0xfd 0x58) i8x16.add_saturate_s)
(I8x16AddSaturateU (0xfd 0x59) i8x16.add_saturate_u)
(I8x16Sub (0xfd 0x5a) i8x16.sub)
(I8x16SubSaturateS (0xfd 0x5b) i8x16.sub_saturate_s)
(I8x16SubSaturateU (0xfd 0x5c) i8x16.sub_saturate_u)
(I8x16Mul (0xfd 0x5d) i8x16.mul)

(I16x8Neg (0xfd 0x62) i16x8.neg)
(I16x8AnyTrue (0xfd 0x63) i16x8.any_true)
(I16x8AllTrue (0xfd 0x64) i16x8.all_true)
(I16x8Shl (0xfd 0x65) i16x8.shl)
(I16x8ShrS (0xfd 0x66) i16x8.shr_s)
(I16x8ShrU (0xfd 0x67) i16x8.shr_u)
(I16x8Add (0xfd 0x68) i16x8.add)
(I16x8AddSaturateS (0xfd 0x69) i16x8.add_saturate_s)
(I16x8AddSaturateU (0xfd 0x6a) i16x8.add_saturate_u)
(I16x8Sub (0xfd 0x6b) i16x8.sub)
(I16x8SubSaturateS (0xfd 0x6c) i16x8.sub_saturate_s)
(I16x8SubSaturateU (0xfd 0x6d) i16x8.sub_saturate_u)
(I16x8Mul (0xfd 0x6e) i16x8.mul)

(I32x4Neg (0xfd 0x73) i32x4.neg)
(I32x4AnyTrue (0xfd 0x74) i32x4.any_true)
(I32x4AllTrue (0xfd 0x75) i32x4.all_true)
(I32x4Shl (0xfd 0x76) i32x4.shl)
(I32x4ShrS (0xfd 0x77) i32x4.shr_s)
(I32x4ShrU (0xfd 0x78) i32x4.shr_u)
(I32x4Add (0xfd 0x79) i32x4.add)
(I32x4Sub (0xfd 0x7c) i32x4.sub)
(I32x4Mul (0xfd 0x7f) i32x4.mul)

(I64x2Neg (0xfd 0x84) i64x2.neg)
(I64x2AnyTrue (0xfd 0x85) i64x2.any_true)
(I64x2AllTrue (0xfd 0x86) i64x2.all_true)
(I64x2Shl (0xfd 0x87) i64x2.shl)
(I64x2ShrS (0xfd 0x88) i64x2.shr_s)
(I64x2ShrU (0xfd 0x89) i64x2.shr_u)
(I64x2Add (0xfd 0x8a) i64x2.add)
(I64x2Sub (0xfd 0x8d) i64x2.sub)
(I64x2Mul (0xfd 0x90) i64x2.mul)

(F32x4Abs (0xfd 0x95) f32x4.abs)
(F32x4Neg (0xfd 0x96) f32x4.neg)
(F32x4Sqrt (0xfd 0x97) f32x4.sqrt)
(F32x4Add (0xfd 0x9a) f32x4.add)
(F32x4Sub (0xfd 0x9b) f32x4.sub)
(F32x4Mul (0xfd 0x9c) f32x4.mul)
(F32x4Div (0xfd 0x9d) f32x4.div)
(F32x4Min (0xfd 0x9e) f32x4.min)
(F32x4Max (0xfd 0x9f) f32x4.max)

(F64x2Abs (0xfd 0xa0) f64x2.abs)
(F64x2Neg (0xfd 0xa1) f64x2.neg)
(F64x2Sqrt (0xfd 0xa2) f64x2.sqrt)
(F64x2Add (0xfd 0xa5) f64x2.add)
(F64x2Sub (0xfd 0xa6) f64x2.sub)
(F64x2Mul (0xfd 0xa7) f64x2.mul)
(F64x2Div (0xfd 0xa8) f64x2.div)
(F64x2Min (0xfd 0xa9) f64x2.min)
(F64x2Max (0xfd 0xaa) f64x2.max)

(I32x4TruncSatF32x4S (0xfd 0xab) i32x4.trunc_sat_f32x4_s)
(I32x4TruncSatF32x4U (0xfd 0xac) i32x4.trunc_sat_f32x4_u)
(I64x2TruncSatF64x2S (0xfd 0xad) i64x2.trunc_sat_f64x2_s)
(I64x2TruncSatF64x2U (0xfd 0xae) i64x2.trunc_sat_f64x2_u)
(F32x4ConvertI32x4S (0xfd 0xaf) f32x4.convert_i32x4_s)
(F32x4ConvertI32x4U (0xfd 0xb0) f32x4.convert_i32x4_u)
(F64x2ConvertI64x2S (0xfd 0xb1) f64x2.convert_i64x2_s)
(F64x2ConvertI64x2U (0xfd 0xb2) f64x2.convert_i64x2_u)
(V8x16Swizzle (0xfd 0xc0) v8x16.swizzle)

;;V8x16Shuffle(V8x16Shuffle) : [0xfd 0xc1] : "v8x16.shuffle"
(V8x16LoadSplat (0xfd 0xc2) v8x16.load_splat (MemArg MemArg<1>))
(V16x8LoadSplat (0xfd 0xc3) v16x8.load_splat (MemArg MemArg<2>))
(V32x4LoadSplat (0xfd 0xc4) v32x4.load_splat (MemArg MemArg<4>))
(V64x2LoadSplat (0xfd 0xc5) v64x2.load_splat (MemArg MemArg<8>))

(I8x16NarrowI16x8S (0xfd 0xc6) i8x16.narrow_i16x8_s)
(I8x16NarrowI16x8U (0xfd 0xc7) i8x16.narrow_i16x8_u)
(I16x8NarrowI32x4S (0xfd 0xc8) i16x8.narrow_i32x4_s)
(I16x8NarrowI32x4U (0xfd 0xc9) i16x8.narrow_i32x4_u)

(I16x8WidenLowI8x16S (0xfd 0xca) i16x8.widen_low_i8x16_s)
(I16x8WidenHighI8x16S (0xfd 0xcb) i16x8.widen_high_i8x16_s)
(I16x8WidenLowI8x16U (0xfd 0xcc) i16x8.widen_low_i8x16_u)
(I16x8WidenHighI8x16u (0xfd 0xcd) i16x8.widen_high_i8x16_u)
(I32x4WidenLowI16x8S (0xfd 0xce) i32x4.widen_low_i16x8_s)
(I32x4WidenHighI16x8S (0xfd 0xcf) i32x4.widen_high_i16x8_s)
(I32x4WidenLowI16x8U (0xfd 0xd0) i32x4.widen_low_i16x8_u)
(I32x4WidenHighI16x8u (0xfd 0xd1) i32x4.widen_high_i16x8_u)

(I16x8Load8x8S (0xfd 0xd2) i16x8.load8x8_s (MemArg MemArg<1>))
(I16x8Load8x8U (0xfd 0xd3) i16x8.load8x8_u (MemArg MemArg<1>))
(I32x4Load16x4S (0xfd 0xd4) i32x4.load16x4_s (MemArg MemArg<2>))
(I32x4Load16x4U (0xfd 0xd5) i32x4.load16x4_u (MemArg MemArg<2>))
(I64x2Load32x2S (0xfd 0xd6) i64x2.load32x2_s (MemArg MemArg<4>))
(I64x2Load32x2U (0xfd 0xd7) i64x2.load32x2_u (MemArg MemArg<4>))
(V128Andnot (0xfd 0xd8) v128.andnot)

`

	allInstrs := mustParseInstrs(instrs)
	all := ""
	for _, ins := range allInstrs {
		all += ins.Generate()
	}

	parseInstr := generateParseInstrution(allInstrs)

	goFile := generate(`
package ast

import (
	"fmt"

	"github.com/ontio/wast-parser/parser"
)

[Instrs]
[parseInstr]
`, map[string]interface{}{"Instrs": all, "parseInstr": parseInstr})

	err := ioutil.WriteFile("../ast/instruction.go", []byte(goFile), 0666)
	if err != nil {
		fmt.Printf("write file error: %s", err)
	}
}
