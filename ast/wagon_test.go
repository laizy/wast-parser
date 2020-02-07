package ast

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/go-interpreter/wagon/wasm"
	"github.com/ontio/wast-parser/parser"
	"github.com/stretchr/testify/assert"
)

func EvalExpr(expr Expression) (int64, error) {
	if len(expr.Instrs) ==  1 {
		switch val := expr.Instrs[0].(type) {
		case *I64Const:
			return val.Val, nil
		case *I32Const:
			return int64(val.Val), nil
		default:
			return 0, errors.New("not supported expression eval")
		}
	}

	return 0, errors.New("not supported expression eval")
}

func EvalExprs(exprs []Expression) ([]int64, error) {
	var result  [] int64
	for _, expr := range exprs {
		val, err := EvalExpr(expr)
		if err != nil {
			return nil, err
		}

		result = append(result, val)
	}

	return result, nil
}

func TestEncode(t *testing.T) {
	wasts, err := LoadWastFiles()
	assert.Nil(t, err)
	for name, content := range wasts {
		fmt.Printf("test file name: %s\n", name)
		ps, err := parser.NewParserBuffer(string(content))
		assert.Nil(t, err)
		var wast Wast
		err = wast.Parse(ps)
		assert.Nil(t, err, fmt.Errorf("parse %s error", name))
		for _, item := range wast.Directives {
			switch direc := item.(type) {
			case Module:
				encode, err := direc.Encode()
				assert.Nil(t, err)
				r := bytes.NewReader(encode)
				_, err = wasm.ReadModule(r, nil)
				if err != nil {
					t.Fatalf("error reading module %v", err)
				}
			}
		}
	}
}







