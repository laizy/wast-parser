package ast

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/go-interpreter/wagon/wasm/leb128"
)

func (self *Module) Encode() []byte {
	var fields []ModuleField
	if t, ok := self.Kind.(ModuleKindText); ok {
		fields = t.Fields
	}

	magic := []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00}
	buf := new(bytes.Buffer)
	buf.Write(magic)

	var types []Type
	var imports []Import
	var funcs []Func
	var tables []Table
	var memories []Memory
	var globals []Global
	var exports []Export
	var start []StartField
	var elem []Elem
	var data []Data

	for _, field := range fields {
		if ty, ok := field.(Type); ok {
			types = append(types, ty)
		} else if imp, ok := field.(Import); ok {
			imports = append(imports, imp)
		} else if fun, ok := field.(Func); ok {
			funcs = append(funcs, fun)
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

	return buf.Bytes()
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

func (g *GlobalValType) Encode(w io.Writer) error {
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

func (f *FunctionType) Encode(w io.Writer) error {
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

func (t *TableType) Encode(w io.Writer) error {
	if err := t.Elem.Encode(w); err != nil {
		return err
	}

	if err := t.Limits.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *MemoryType) Encode(w io.Writer) error {
	if err := t.Limits.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *Table) Encode(w io.Writer) error {
	// check why can not zero
	if len(t.Exports.Names) == 0 {
		return errors.New("Name should not empty")
	}

	if x, ok := t.Kind.(TableKindNormal); ok {
		return x.Type.Encode(w)
	}

	return errors.New("TableKind should be normal during encoding")
}

func (t *Memory) Encode(w io.Writer) error {
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

func (t *Import) Encode(w io.Writer) error {
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

func (t *Global) Encode(w io.Writer) error {
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

func (t *Export) Encode(w io.Writer) error {
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

func (t *Elem) Encode(w io.Writer) error {
	return nil
}

func (t *ElemPayloadIndices) Encode(w io.Writer) error {
	return nil
}

func (t *ElemPayloadExprs) Encode(w io.Writer) error {
	return nil
}

func (t *Data) Encode(w io.Writer) error {
	return nil
}

func (t *Func) Encode(w io.Writer) error {
	return nil
}

func (t *Expression) Encode(w io.Writer) error {
	return nil
}

func (t *TypeUse) Encode(w io.Writer) error {
	return t.Index.Encode(w)
}

func (t Index) Encode(w io.Writer) error {
	if t.Isnum {
		if _, err := leb128.WriteVarUint32(w, t.Num); err != nil {
			return err
		}
	}

	return fmt.Errorf("unresolved index in emission %s", t.Id.Name)
}

func (t *OptionIndex) Encode(w io.Writer) error {
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
