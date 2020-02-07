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
	//todo : fix tests/*.wast data
	wasts = make(map[string][]byte)
	spec, err := LoadWastFilesFromDir("../tests/spectestdata/")
	if err != nil {
		return nil, err
	}

	notTestFile := []string{
		"call.wast",
		"conversions.wast",
		"const.wast",
		"endianness.wast",
		"f32_cmp.wast",
		"memory_redundancy.wast",
		"call_indirect.wast",
		"f64_cmp.wast",
		"f32_bitwise.wast",
		"float_exprs.wast",
		"float_literals.wast",
		"f32.wast",
		"f64.wast",
		"f64_bitwise.wast",
		"float_misc.wast",
	}
	for name, content := range spec {
		allow := true
		for _, f := range notTestFile {
			if strings.HasSuffix(name, f) {
				allow = false
				break
			}
		}
		if allow {
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
