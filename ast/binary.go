package ast

import (
	"errors"
	"fmt"
)

type Section interface {
	Encode(sink *ZeroCopySink)
}

func (self *Module) Encode() ([]byte, error) {
	var fields []ModuleField
	if t, ok := self.Kind.(ModuleKindText); ok {
		fields = t.Fields
	}

	magic := []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00}
	sink := NewZeroCopySink(nil)
	sink.WriteBytes(magic)

	var types []Section
	var imports []Section
	var funcs []Section
	var funcsTypes []Section
	var tables []Section
	var memories []Section
	var globals []Section
	var exports []Section
	var start []Section
	var elem []Section
	var data []Section

	for _, field := range fields {
		switch field.(type) {
		case Type:
			types = append(types, field.(Type))
		case Import:
			imports = append(imports, field.(Import))
		case Func:
			funcsTypes = append(funcsTypes, field.(Func).Type)
			funcs = append(funcs, field.(Func))
		case Table:
			tables = append(tables, field.(Table))
		case Memory:
			memories = append(memories, field.(Memory))
		case Global:
			globals = append(globals, field.(Global))
		case Export:
			exports = append(exports, field.(Export))
		case StartField:
			start = append(start, field.(StartField))
		case Elem:
			elem = append(elem, field.(Elem))
		case Data:
			data = append(data, field.(Data))
		default:
			return nil, errors.New("err section")
		}
	}

	SectionList(0x1, types, sink)
	SectionList(0x2, imports, sink)
	SectionList(0x3, funcsTypes, sink)
	SectionList(0x4, tables, sink)
	SectionList(0x5, memories, sink)
	SectionList(0x6, globals, sink)
	SectionList(0x7, exports, sink)
	SectionList(0x8, start, sink)
	SectionList(0x9, elem, sink)
	SectionList(0xa, funcs, sink)
	SectionList(0xb, data, sink)

	return sink.Bytes(), nil
}

func SectionList(id byte, l []Section, sink *ZeroCopySink) error {
	if len(l) == 0 {
		return nil
	}

	tmpSink := NewZeroCopySink(nil)
	ListEncode(l, tmpSink)

	sink.WriteByte(id)
	sink.WriteVarBytes(tmpSink.Bytes())

	return nil
}

func ListEncode(l []Section, sink *ZeroCopySink) {
	sink.WriteUint32(uint32(len(l)))

	for _, s := range l {
		s.Encode(sink)
	}
}

const TypeFunc uint8 = 0x60

func (t ValType) Encode(sink *ZeroCopySink) {
	var vt byte
	switch t {
	case I32:
		vt = 0x7f
	case I64:
		vt = 0x7e
	case F32:
		vt = 0x7d
	case F64:
		vt = 0x7c
	case Anyref:
		vt = 0x6f
	case Funcref:
		vt = 0x70
	case V128:
		vt = 0x7b
	}

	sink.WriteByte(vt)
}

func (t Type) Encode(sink *ZeroCopySink) {
	t.Func.Encode(sink)
}

func (g GlobalValType) Encode(sink *ZeroCopySink) {
	g.Type.Encode(sink)

	var m uint8
	if g.Mutable {
		m = 1
	}

	sink.WriteUint8(m)
}

func (t TableElemType) Encode(sink *ZeroCopySink) {
	var vt ValType
	switch t {
	case FuncRef:
		vt = Funcref
	case AnyRef:
		vt = Anyref
	}

	vt.Encode(sink)
}

func (f FunctionType) Encode(sink *ZeroCopySink) {
	sink.WriteByte(0x60)
	sink.WriteUint32(uint32(len(f.Params)))

	for _, p := range f.Params {
		p.Val.Encode(sink)
	}

	sink.WriteUint32(uint32(len(f.Results)))

	for _, p := range f.Results {
		p.Encode(sink)
	}
}

func (t Limits) Encode(sink *ZeroCopySink) {
	if t.Max == 0 {
		sink.WriteByte(0x00)
		sink.WriteUint32(t.Min)
	} else {
		sink.WriteByte(0x01)
		sink.WriteUint32(t.Min)
		sink.WriteUint32(t.Max)
	}
}

func (t TableType) Encode(sink *ZeroCopySink) {
	t.Elem.Encode(sink)
	t.Limits.Encode(sink)
}

func (t MemoryType) Encode(sink *ZeroCopySink) {
	t.Limits.Encode(sink)
}

func (t Table) Encode(sink *ZeroCopySink) {
	// check why can not zero
	if len(t.Exports.Names) == 0 {
		panic("Name should not empty")
	}

	if x, ok := t.Kind.(TableKindNormal); ok {
		x.Type.Encode(sink)
		return
	}

	panic("TableKind should be normal during encoding")
}

func (t Memory) Encode(sink *ZeroCopySink) {
	if mem, ok := t.Kind.(*MemoryKindNormal); ok {
		mem.Type.Encode(sink)
		return
	}

	panic("MemoryKind should be normal during encoding")
}

func (t Import) Encode(sink *ZeroCopySink) {
	sink.WriteString(t.Module)
	sink.WriteString(t.Field)

	impType := t.Item.ImportType()
	switch impType {
	case "func":
		sink.WriteByte(0x00)
	case "table":
		sink.WriteByte(0x01)
	case "memory":
		sink.WriteByte(0x02)
	case "global":
		sink.WriteByte(0x03)
	}

	t.Item.Encode(sink)
}

func (t Global) Encode(sink *ZeroCopySink) {
	// check why can not zero
	if len(t.Exports.Names) == 0 {
		panic("Name should not empty")
	}

	t.ValType.Encode(sink)

	exp, ok := t.Kind.(GlobalKindInline)

	if !ok {
		panic("GlobalKind should be inline during encoding")
	}

	exp.Expr.Encode(sink)
}

func (t Export) Encode(sink *ZeroCopySink) {
	sink.WriteString(t.Name)
	sink.WriteByte(byte(t.Type))
	t.Index.Encode(sink)
}

func (t Elem) Encode(sink *ZeroCopySink) {
	switch t.Kind.(type) {
	case ElemKindActive:
		active, _ := t.Kind.(ElemKindActive)
		if !active.Table.Isnum {
			panic("expect num in Elem kind")
		}

		switch t.Payload.(type) {
		case ElemPayloadIndices:
			if active.Table.Num == 0 && !t.forceNonZero {
				sink.WriteByte(byte(0x00))
				active.Offset.Encode(sink)
			} else {
				sink.WriteByte(byte(0x02))
				active.Table.Encode(sink)
				active.Offset.Encode(sink)
				sink.WriteByte(byte(0x00))
			}
		case ElemPayloadExprs:
			expr, _ := t.Payload.(ElemPayloadExprs)
			if active.Table.Num == 0 && expr.Type == FuncRef {
				sink.WriteByte(byte(0x04))
				active.Offset.Encode(sink)
			} else {
				sink.WriteByte(byte(0x06))
				active.Table.Encode(sink)
				active.Offset.Encode(sink)
				expr.Type.Encode(sink)
			}
		default:
			panic("error Elem payload Kind")
		}
	case ElemKindPassive:
		switch t.Payload.(type) {
		case ElemPayloadIndices:
			sink.WriteByte(byte(0x01))
			sink.WriteByte(byte(0x00))
		case ElemPayloadExprs:
			expr, _ := t.Payload.(ElemPayloadExprs)
			sink.WriteByte(byte(0x05))
			expr.Type.Encode(sink)
		default:
			panic("error Elem payload Kind")
		}
	default:
		panic("error Elem Kind")
	}

	t.Payload.Encode(sink)
}

func (t ElemPayloadIndices) Encode(sink *ZeroCopySink) {
	var indices []Section
	for _, i := range t.Indices {
		indices = append(indices, i)
	}

	ListEncode(indices, sink)
}

func (t ElemPayloadExprs) Encode(sink *ZeroCopySink) {
}

func (t Data) Encode(sink *ZeroCopySink) {
	switch t.Kind.(type) {
	case DataKindPassive:
		sink.WriteByte(byte(0x01))
	case DataKindActive:
		active, _ := t.Kind.(DataKindActive)
		if active.Memory.Isnum && active.Memory.Num == 0 {
			sink.WriteByte(byte(0x00))
		} else {
			sink.WriteByte(byte(0x02))
			active.Memory.Encode(sink)
		}
	default:
		panic("error data kind")
	}

	var l uint32
	for _, v := range t.Val {
		l = l + uint32(len(v))
	}

	sink.WriteUint32(l)

	for _, v := range t.Val {
		sink.WriteBytes(v)
	}
}

func (t Func) Encode(sink *ZeroCopySink) {
	switch t.Kind.(type) {
	case FuncKindInline:
		fun := t.Kind.(FuncKindInline)

		type compressLocal struct {
			num   uint32
			local Local
		}
		tmpSink := NewZeroCopySink(nil)
		var comL []compressLocal

		for _, ct := range fun.Locals {
			if len(comL) > 0 && ct.ValType == comL[len(comL)-1].local.ValType {
				comL[len(comL)-1].num += 1
				continue
			}

			comL = append(comL, compressLocal{
				num:   1,
				local: ct,
			})
		}

		tmpSink.WriteUint32(uint32(len(comL)))
		for _, v := range comL {
			tmpSink.WriteUint32(v.num)
			v.local.ValType.Encode(tmpSink)
		}

		fun.Expr.Encode(tmpSink)
		sink.WriteVarBytes(tmpSink.Bytes())
	default:
		panic("should only have inline functions in emission")
	}

}

func (t Expression) Encode(sink *ZeroCopySink) {
	for _, inst := range t.Instrs {
		inst.Encode(sink)
	}

	sink.WriteByte(byte(0x0b))
}

func (t TypeUse) Encode(sink *ZeroCopySink) {
	t.Index.Encode(sink)
}

func (t StartField) Encode(sink *ZeroCopySink) {
}

func (self ImportFunc) Encode(sink *ZeroCopySink) {
	self.TypeUse.Encode(sink)
}

func (self ImportGlobal) Encode(sink *ZeroCopySink) {
	self.Global.Encode(sink)
}

func (self ImportMemory) Encode(sink *ZeroCopySink) {
	self.Mem.Encode(sink)
}

func (self ImportTable) Encode(sink *ZeroCopySink) {
	self.Table.Encode(sink)
}

func (t Index) Encode(sink *ZeroCopySink) {
	if t.Isnum {
		sink.WriteUint32(t.Num)
		return
	}

	panic(fmt.Errorf("unresolved index in emission %s", t.Id.Name))
}

func (t OptionIndex) Encode(sink *ZeroCopySink) {
	t.ToIndex().Encode(sink)
}

func (t BlockType) Encode(sink *ZeroCopySink) {
	if t.Ty.Index.IsSome() {
		t.Ty.Index.Encode(sink)
		return
	}

	if len(t.Ty.Type.Params) == 0 && len(t.Ty.Type.Results) == 0 {
		sink.WriteByte(byte(0x40))
	}

	if len(t.Ty.Type.Params) == 0 && len(t.Ty.Type.Results) == 1 {
		t.Ty.Type.Results[0].Encode(sink)
	}

	panic("multi-value block types should have an index")
}

func (t MemArg) Encode(sink *ZeroCopySink) {
	sink.WriteUint32(t.Align)
	sink.WriteUint32(t.Offset)
}

func (t CallIndirectInner) Encode(sink *ZeroCopySink) {
	t.Table.Encode(sink)
	t.Type.Encode(sink)
}

func (t BrTableIndices) Encode(sink *ZeroCopySink) {
	sink.WriteUint32(uint32(len(t.Labels)))
	for _, s := range t.Labels {
		s.Encode(sink)
	}

	t.Default.Encode(sink)
}

func (t OptionId) Encode(sink *ZeroCopySink) {
	//Do nothing
}

func (t SelectTypes) Encode(sink *ZeroCopySink) {
	if len(t.Types) == 0 {
		sink.WriteByte(byte(0x1b))
	} else {
		sink.WriteByte(byte(0x1c))
		sink.WriteUint32(uint32(len(t.Types)))
		for _, s := range t.Types {
			s.Encode(sink)
		}
	}
}

func (t Float32) Encode(sink *ZeroCopySink) {
	sink.WriteFloat32(t.Bits)
}

func (t Float64) Encode(sink *ZeroCopySink) {
	sink.WriteFloat64(t.Bits)
}
