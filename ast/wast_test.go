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

func LoadWastFilesFromDir(dir string) (map[string][]byte, error) {
	wasts := make(map[string][]byte)
	fnames, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}
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

func LoadWastFiles() (map[string][]byte, error) {
	wasts, err := LoadWastFilesFromDir("../tests/")
	if err != nil {
		return nil, err
	}
	spec, err := LoadWastFilesFromDir("../tests/spectestdata/")
	if err != nil {
		return nil, err
	}

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
	for name, content := range spec {
		if !notTestFile[name] {
			wasts[name] = content
		}
	}

	return wasts, nil
}

func TestWastParsing(t *testing.T) {
	wasts, err := LoadWastFiles()
	assert.Nil(t, err)
	for name, content := range wasts {
		fmt.Printf("test file name: %s\n", name)
		ps, err := parser.NewParserBuffer(string(content))
		assert.Nil(t, err)
		var wast Wast
		err = wast.Parse(ps)
		assert.Nil(t, err, fmt.Errorf("parse %s error", name))
	}
}
