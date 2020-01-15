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
  (type $t0 (func (param i32) (param i32) (result i32)))
  (func $add (export "add") (type $t0) (param $p0 i32) (param $p1 i32) (result i32)
    get_local $p0
    get_local $p1
    i32.add)
)
`)
	assert.Nil(t, err)

	var module Wat
	err = module.Parse(ps)
	assert.Nil(t, err)

	_, err = module.Module.Encode()
	assert.Nil(t, err)

	fmt.Printf("tokens: %v", module.Module)
}
