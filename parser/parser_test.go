package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	parser, err := NewParserBuffer(`
(module
  (memory 1)
  (data (i32.const 0) "abcdefghijklmnopqrstuvwxyz")

  (func (export "8u_good1") (param $i i32) (result i32)
    (i32.load8_u offset=0 (local.get $i))                   ;; 97 'a'
  )
)
`)

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	fmt.Printf("tokens: %v", parser)
}
