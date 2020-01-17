package ast

import (
	"fmt"
	"testing"

	"github.com/ontio/wast-parser/parser"
	"github.com/stretchr/testify/assert"
)

func TestModuleEncode(t *testing.T) {
	ps, err := parser.NewParserBuffer(`
(module
  (type (func (param i32) (param i32) (result i32)))
  (func (type 0)
    get_local 0
    get_local 1
    i32.add)
)
`)
	assert.Nil(t, err)

	var module Wat
	err = module.Parse(ps)
	assert.Nil(t, err)

	b, err := module.Module.Encode()
	assert.Nil(t, err)
	assert.Equal(t, b, []byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00, 0x01, 0x07, 0x01, 0x60, 0x02, 0x7f, 0x7f, 0x01, 0x7f, 0x03, 0x02, 0x01, 0x00, 0x0a, 0x09, 0x01, 0x07, 0x00, 0x20, 0x00, 0x20, 0x01, 0x6a, 0x0b})
	fmt.Printf("tokens: %v", module.Module)
}
