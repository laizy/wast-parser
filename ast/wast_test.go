package ast

import (
	"fmt"
	"github.com/ontio/wast-parser/parser"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func LoadWastFiles(dir string) (map[string][]byte, error) {
	wasts := make(map[string][]byte)
	fnames, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
	//fnames = []string{
	//	//"../tests/spectestdata/data.wast",
	//	//"../tests/spectestdata/address.wast",
	//	//"../tests/spectestdata/memory.wast",
	//	//"../tests/spectestdata/func.wast",
	//	"../tests/spectestdata/const.wast",
	//}
	for _, name := range fnames {
		if !strings.HasSuffix(name, ".wast") {
			continue
		}
		raw, err := ioutil.ReadFile(name)
		if err != nil {
			return nil, err
		}
		wasts[path.Base(name)] = raw
	}

	return wasts, nil
}

func TestWastParsing(t *testing.T) {
	wasts, err := LoadWastFiles("../tests/spectestdata/")
	assert.Nil(t, err)
	notTestFile := map[string]bool{
		"call.wast":              true,
		"conversions.wast":       true,
		"const.wast":             true,
		"endianness.wast":        true,
		"f32_cmp.wast":           true,
		"memory_redundancy.wast": true,
		"call_indirect.wast":     true,
		"f64_cmp.wast":           true,
		"f32_bitwise.wast":       true,
		"float_exprs.wast":       true,
		"float_literals.wast":    true,
		"f32.wast":               true,
		"f64.wast":               true,
		"f64_bitwise.wast":       true,
		"float_misc.wast":        true,
	}
	for name, content := range wasts {
		if notTestFile[name] {
			continue
		}
		fmt.Printf("test file name: %s\n", name)
		ps, err := parser.NewParserBuffer(string(content))
		assert.Nil(t, err)
		var wast Wast
		err = wast.Parse(ps)
		assert.Nil(t, err, fmt.Errorf("parse %s error", name))
	}
}

func TestWastParsing2(t *testing.T) {
	wasts, err := LoadWastFiles("../tests/")
	assert.Nil(t, err)
	for name, content := range wasts {
		fmt.Printf("2test file name: %s\n", name)
		ps, err := parser.NewParserBuffer(string(content))
		assert.Nil(t, err)
		var wast Wast
		err = wast.Parse(ps)
		assert.Nil(t, err, fmt.Errorf("parse %s error", name))
	}
}
