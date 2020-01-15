package ast

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/go-interpreter/wagon/wasm/leb128"
)

type Section interface {
	Encode(w io.Writer) error
}

func (self *Module) Encode() ([]byte, error) {
	var fields []ModuleField
	if t, ok := self.Kind.(ModuleKindText); ok {
		fields = t.Fields
	}

	magic := []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00}
	buf := new(bytes.Buffer)
	if _, err := buf.Write(magic); err != nil {
		return nil, err
	}

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
		if ty, ok := field.(Type); ok {
			types = append(types, ty)
		} else if imp, ok := field.(Import); ok {
			imports = append(imports, imp)
		} else if fun, ok := field.(Func); ok {
			funcsTypes = append(funcs, fun.Type)
			funcs = append(funcs, fun.Type)
		} else if table, ok := field.(Table); ok {
			tables = append(tables, table)
		} else if mem, ok := field.(Memory); ok {
			memories = append(memories, mem)
		} else if global, ok := field.(Global); ok {
			globals = append(globals, global)
		} else if val, ok := field.(Export); ok {
			exports = append(exports, val)
		} else if val, ok := field.(StartField); ok {
			start = append(start, val)
		} else if val, ok := field.(Elem); ok {
			elem = append(elem, val)
		} else if val, ok := field.(Data); ok {
			data = append(data, val)
		}
	}

	SectionList(0x1, types, buf)
	SectionList(0x2, imports, buf)
	SectionList(0x3, funcsTypes, buf)
	SectionList(0x4, tables, buf)
	SectionList(0x5, memories, buf)
	SectionList(0x6, globals, buf)
	SectionList(0x7, exports, buf)
	SectionList(0x8, start, buf)
	SectionList(0x9, elem, buf)
	SectionList(0xa, funcs, buf)
	SectionList(0xb, data, buf)

	return buf.Bytes(), nil
}

func SectionList(id byte, l []Section, w io.Writer) error {
	if len(l) == 0 {
		return nil
	}

	tmp := new(bytes.Buffer)
	if err := ListEncode(l, tmp); err != nil {
		return err
	}

	if err := writeByte(w, id); err != nil {
		return err
	}

	if err := BytesEncode(tmp.Bytes(), w); err != nil {
		return err
	}

	return nil
}

func ListEncode(l []Section, w io.Writer) error {
	if _, err := leb128.WriteVarUint32(w, uint32(len(l))); err != nil {
		return err
	}

	for _, s := range l {
		s.Encode(w)
	}

	return nil
}

func BytesEncode(b []byte, w io.Writer) error {
	if _, err := leb128.WriteVarUint32(w, uint32(len(b))); err != nil {
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

const TypeFunc uint8 = 0x60

func (t ValType) Encode(w io.Writer) error {
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

	err := writeByte(w, vt)
	if err != nil {
		return err
	}

	return err
}

func (t Type) Encode(w io.Writer) error {
	return t.Func.Encode(w)
}

func (g GlobalValType) Encode(w io.Writer) error {
	if err := g.Type.Encode(w); err != nil {
		return err
	}

	var m uint8
	if g.Mutable {
		m = 1
	}

	return writeByte(w, m)
}

func (t TableElemType) Encode(w io.Writer) error {
	var vt ValType
	switch t {
	case FuncRef:
		vt = Funcref
	case AnyRef:
		vt = Anyref
	}

	return vt.Encode(w)
}

func (f FunctionType) Encode(w io.Writer) error {
	err := writeByte(w, 0x60)
	if err != nil {
		return err
	}

	_, err = leb128.WriteVarUint32(w, uint32(len(f.Params)))
	if err != nil {
		return err
	}
	for _, p := range f.Params {
		err = p.Val.Encode(w)
		if err != nil {
			return err
		}
	}

	_, err = leb128.WriteVarUint32(w, uint32(len(f.Results)))
	if err != nil {
		return err
	}
	for _, p := range f.Results {
		err = p.Encode(w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t Limits) Encode(w io.Writer) error {
	if t.Max == 0 {
		err := writeByte(w, 0x00)
		if err != nil {
			return err
		}

		_, err = leb128.WriteVarUint32(w, t.Min)
		if err != nil {
			return err
		}
	} else {
		err := writeByte(w, 0x01)
		if err != nil {
			return err
		}

		_, err = leb128.WriteVarUint32(w, t.Min)
		if err != nil {
			return err
		}

		_, err = leb128.WriteVarUint32(w, t.Max)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t TableType) Encode(w io.Writer) error {
	if err := t.Elem.Encode(w); err != nil {
		return err
	}

	if err := t.Limits.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t MemoryType) Encode(w io.Writer) error {
	if err := t.Limits.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t Table) Encode(w io.Writer) error {
	// check why can not zero
	if len(t.Exports.Names) == 0 {
		return errors.New("Name should not empty")
	}

	if x, ok := t.Kind.(TableKindNormal); ok {
		return x.Type.Encode(w)
	}

	return errors.New("TableKind should be normal during encoding")
}

func (t Memory) Encode(w io.Writer) error {
	if mem, ok := t.Kind.(*MemoryKindNormal); ok {
		err := mem.Type.Encode(w)
		if err != nil {
			return err
		}
	} else {
		return errors.New("MemoryKind should be normal during encoding")
	}

	return nil
}

func (t Import) Encode(w io.Writer) error {
	if err := writeStringUint(w, t.Module); err != nil {
		return err
	}

	if err := writeStringUint(w, t.Field); err != nil {
		return err
	}

	impType := t.Item.ImportType()
	switch impType {
	case "func":
		if err := writeByte(w, 0x00); err != nil {
			return err
		}
	case "table":
		if err := writeByte(w, 0x01); err != nil {
			return err
		}
	case "memory":
		if err := writeByte(w, 0x02); err != nil {
			return err
		}
	case "global":
		if err := writeByte(w, 0x03); err != nil {
			return err
		}
	}

	if err := t.Item.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t Global) Encode(w io.Writer) error {
	// check why can not zero
	if len(t.Exports.Names) == 0 {
		return errors.New("Name should not empty")
	}

	if err := t.ValType.Encode(w); err != nil {
		return err
	}

	exp, ok := t.Kind.(GlobalKindInline)

	if !ok {
		return errors.New("GlobalKind should be inline during encoding")
	}

	return exp.Expr.Encode(w)
}

func (t Export) Encode(w io.Writer) error {
	if err := writeStringUint(w, t.Name); err != nil {
		return err
	}

	if err := writeByte(w, byte(t.Type)); err != nil {
		return err
	}

	if err := t.Index.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t Elem) Encode(w io.Writer) error {
	switch t.Kind.(type) {
	case ElemKindActive:
		active, _ := t.Kind.(ElemKindActive)
		if !active.Table.Isnum {
			return errors.New("expect num in Elem kind")
		}

		switch t.Payload.(type) {
		case ElemPayloadIndices:
			if active.Table.Num == 0 && !t.forceNonZero {
				if err := writeByte(w, byte(0x00)); err != nil {
					return err
				}
				if err := active.Offset.Encode(w); err != nil {
					return err
				}
			} else {
				if err := writeByte(w, byte(0x02)); err != nil {
					return err
				}
				if err := active.Table.Encode(w); err != nil {
					return err
				}
				if err := active.Offset.Encode(w); err != nil {
					return err
				}
				if err := writeByte(w, byte(0x00)); err != nil {
					return err
				}
			}
		case ElemPayloadExprs:
			expr, _ := t.Payload.(ElemPayloadExprs)
			if active.Table.Num == 0 && expr.Type == FuncRef {
				if err := writeByte(w, byte(0x04)); err != nil {
					return err
				}

				if err := active.Offset.Encode(w); err != nil {
					return err
				}
			} else {
				if err := writeByte(w, byte(0x06)); err != nil {
					return err
				}

				if err := active.Table.Encode(w); err != nil {
					return err
				}
				if err := active.Offset.Encode(w); err != nil {
					return err
				}
				if err := expr.Type.Encode(w); err != nil {
					return err
				}
			}
		default:
			return errors.New("error Elem payload Kind")
		}
	case ElemKindPassive:
		switch t.Payload.(type) {
		case ElemPayloadIndices:
			if err := writeByte(w, byte(0x01)); err != nil {
				return err
			}
			if err := writeByte(w, byte(0x00)); err != nil {
				return err
			}
		case ElemPayloadExprs:
			expr, _ := t.Payload.(ElemPayloadExprs)
			if err := writeByte(w, byte(0x05)); err != nil {
				return err
			}
			if err := expr.Type.Encode(w); err != nil {
				return err
			}
		default:
			return errors.New("error Elem payload Kind")
		}
	default:
		return errors.New("error Elem Kind")
	}
	//

	if active, ok := t.Kind.(ElemKindActive); ok {
		if !active.Table.Isnum {
			return errors.New("expect num in Elem kind")
		}

		if _, ok := t.Payload.(ElemPayloadIndices); ok {
			if active.Table.Num == 0 && !t.forceNonZero {
				if err := writeByte(w, byte(0x00)); err != nil {
					return err
				}
				return active.Offset.Encode(w)
			} else {
				if err := writeByte(w, byte(0x02)); err != nil {
					return err
				}
				if err := active.Table.Encode(w); err != nil {
					return err
				}
				if err := active.Offset.Encode(w); err != nil {
					return err
				}
				if err := writeByte(w, byte(0x00)); err != nil {
					return err
				}
			}
		} else if expr, ok := t.Payload.(ElemPayloadExprs); ok {
			if active.Table.Num == 0 && expr.Type == FuncRef {
				if err := writeByte(w, byte(0x04)); err != nil {
					return err
				}

				return active.Offset.Encode(w)
			}
		} else {
			return errors.New("error Elem payload Kind")
		}
	} else if _, ok := t.Kind.(ElemKindPassive); ok {
		if _, ok := t.Payload.(ElemPayloadIndices); ok {
			if err := writeByte(w, byte(0x01)); err != nil {
				return err
			}
			if err := writeByte(w, byte(0x00)); err != nil {
				return err
			}
		} else if expr, ok := t.Payload.(ElemPayloadExprs); ok {
			if err := writeByte(w, byte(0x05)); err != nil {
				return err
			}

			if err := expr.Type.Encode(w); err != nil {
				return err
			}

		}
	} else {
		return errors.New("error Elem Kind")
	}

	if err := t.Payload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t ElemPayloadIndices) Encode(w io.Writer) error {
	var indices []Section
	for _, i := range t.Indices {
		indices = append(indices, i)
	}

	if err := ListEncode(indices, w); err != nil {
		return err
	}

	return nil
}

func (t ElemPayloadExprs) Encode(w io.Writer) error {
	// to do
	return nil
}

func (t Data) Encode(w io.Writer) error {
	// todo
	return nil
}

func (t Func) Encode(w io.Writer) error {
	// to do
	return nil
}

func (t Expression) Encode(w io.Writer) error {
	// to do
	return nil
}

func (t TypeUse) Encode(w io.Writer) error {
	return t.Index.Encode(w)
}

func (t StartField) Encode(w io.Writer) error {
	return nil
}

func (self ImportFunc) Encode(w io.Writer) error {
	return self.TypeUse.Encode(w)
}

func (self ImportGlobal) Encode(w io.Writer) error {
	return self.Global.Encode(w)
}

func (self ImportMemory) Encode(w io.Writer) error {
	return self.Mem.Encode(w)
}

func (self ImportTable) Encode(w io.Writer) error {
	return self.Table.Encode(w)
}

func (t Index) Encode(w io.Writer) error {
	if t.Isnum {
		if _, err := leb128.WriteVarUint32(w, t.Num); err != nil {
			return err
		}
	}

	return fmt.Errorf("unresolved index in emission %s", t.Id.Name)
}

func (t OptionIndex) Encode(w io.Writer) error {
	return t.ToIndex().Encode(w)
}

func writeStringUint(w io.Writer, s string) error {
	return writeBytesUint(w, []byte(s))
}

func writeBytesUint(w io.Writer, p []byte) error {
	_, err := leb128.WriteVarUint32(w, uint32(len(p)))
	if err != nil {
		return err
	}
	_, err = w.Write(p)
	return err
}

func writeU32(w io.Writer, n uint32) error {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], n)
	_, err := w.Write(buf[:])
	return err
}

func writeU64(w io.Writer, n uint64) error {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], n)
	_, err := w.Write(buf[:])
	return err
}
