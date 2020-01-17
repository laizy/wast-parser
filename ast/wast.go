package ast

import (
	"errors"
	"fmt"
	"github.com/ontio/wast-parser/lexer"
	"github.com/ontio/wast-parser/parser"
	"strings"
)

type Wat struct {
	Module Module
}

func (self *Wat) Parse(ps *parser.ParserBuffer) error {
	token := ps.Peek2Token()
	if matchKeyword(token, "module") == false {
		mfs := make([]ModuleField, 0)
		for !ps.Empty() {
			err := ps.Parens(func(ps *parser.ParserBuffer) error {
				mf, err := parseModuleField(ps)
				if err != nil {
					return err
				}
				mfs = append(mfs, mf)
				return nil
			})
			if err != nil {
				return err
			}
		}
		self.Module = Module{
			Name: NoneOptionId(),
			Kind: ModuleKindText{
				Fields: mfs,
			},
		}
		return nil
	}
	return ps.Parens(func(ps *parser.ParserBuffer) error {
		return self.Module.Parse(ps)
	})
}

type Wast struct {
	Directives []WastDirective
}

func (self *Wast) Parse(ps *parser.ParserBuffer) error {
	if isWastDirectiveToken(ps.Peek2Token()) {
		for !ps.Empty() {
			err := ps.Parens(func(ps *parser.ParserBuffer) error {
				dir, err := parseWastDirective(ps)
				if err != nil {
					return err
				}
				self.Directives = append(self.Directives, dir)
				return nil
			})
			if err != nil {
				return err
			}
		}
	} else {
		var wat Wat
		err := wat.Parse(ps)
		if err != nil {
			return err
		}
		self.Directives = append(self.Directives, wat.Module)
	}

	return nil
}

func isWastDirectiveToken(token lexer.Token) bool {
	if token == nil || token.Type() != lexer.KeywordType {
		return false
	}
	kw := token.(lexer.Keyword).Val
	if strings.HasPrefix(kw, "assert_") || kw == "module" || kw == "register" || kw == "invoke" {
		return true
	}

	return false
}

type WastDirective interface {
	wastDirective()
}

type implWastDirective struct{}
func (self implWastDirective) wastDirective() {}

type AssertInvalidDirective struct {
	implWastDirective
	Module Module
	Msg    string
}

type AssertExhaustionDirective struct {
	implWastDirective
	Call    WastInvoke
	Message string
}

type AssertUnlinkableDirective struct {
	implWastDirective
	Module Module
	Msg    string
}

type AssertMalformedDirective struct {
	implWastDirective
	Module QuoteModule
	Msg    string
}

type RegisterDirective struct {
	implWastDirective
	Name   string
	Module OptionId
}

type AssertTrapDirective struct {
	implWastDirective
	Exec WastExecute
	msg  string
}

type AssertReturnDirective struct {
	implWastDirective
	Exec    WastExecute
	Results []Expression
}

type AssertReturnCanonicalNanDirective struct {
	implWastDirective
	Invoke WastInvoke
}

type WastExecute interface {
	wastExecute()
}
type implWastExecute struct{}
func (self implWastExecute) wastExecute() {}

type WastInvoke struct {
	implWastDirective
	implWastExecute
	Module OptionId
	Name   string
	Args   []Expression
}

type WastExecuteGet struct {
	implWastExecute
	Module OptionId
	Global string
}

func (self *WastInvoke) Parse(ps *parser.ParserBuffer) error {
	err := ps.ExpectKeywordMatch("invoke")
	if err != nil {
		return err
	}

	self.Module.Parse(ps)
	self.Name, err = ps.ExpectString()
	if err != nil {
		return err
	}

	for !ps.Empty() {
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			var expr Expression
			err := expr.Parse(ps)
			if err != nil {
				return err
			}

			self.Args = append(self.Args, expr)
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

type QuoteModule interface {
	quoteModule()
}

type implQuoteModule struct{}
func (self implQuoteModule) quoteModule() {}

type Quote struct {
	implQuoteModule
	Data []string
}

func parseQuoteModule(ps *parser.ParserBuffer) (QuoteModule, error) {
	if matchKeyword(ps.Peek2Token(), "quote") {
		err := ps.ExpectKeywordMatch("module")
		if err != nil {
			return nil, err
		}
		_ = ps.ExpectKeywordMatch("quote")
		var quote Quote
		for !ps.Empty() {
			str, err := ps.ExpectString()
			if err != nil {
				return nil, err
			}
			quote.Data = append(quote.Data, str)
		}
		return quote, nil
	}

	var module Module
	err := module.Parse(ps)

	return module, err
}

func parseWastDirective(ps *parser.ParserBuffer) (WastDirective, error) {
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}
	switch kw {
	case "module":
		ps.StepBack(1)
		var module Module
		err := module.Parse(ps)
		return module, err
	case "assert_malformed":
		var result AssertMalformedDirective
		err = ps.Parens(func(ps *parser.ParserBuffer) error {
			result.Module, err = parseQuoteModule(ps)
			return err
		})
		if err != nil {
			return nil, err
		}
		result.Msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "assert_invalid":
		var result AssertInvalidDirective
		err = ps.Parens(result.Module.Parse)
		if err != nil {
			return nil, err
		}
		result.Msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "register":
		var result RegisterDirective
		result.Name, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		result.Module.Parse(ps)
		return result, nil
	case "invoke":
		ps.StepBack(1)
		var invoke WastInvoke
		err := invoke.Parse(ps)
		if err != nil {
			return nil, err
		}
		return invoke, nil
	case "assert_trap":
		var trap AssertTrapDirective
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			trap.Exec, err = parseWastExecute(ps)
			return err
		})
		if err != nil {
			return nil, err
		}

		trap.msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}

		return trap, nil
	case "assert_return":
		var ret AssertReturnDirective
		err := ps.Parens(func(ps *parser.ParserBuffer) error {
			ret.Exec, err = parseWastExecute(ps)
			return err
		})
		if err != nil {
			return nil, err
		}
		for !ps.Empty() {
			err := ps.Parens(func(ps *parser.ParserBuffer) error {
				var expr Expression
				err := expr.Parse(ps)
				if err != nil {
					return err
				}
				ret.Results = append(ret.Results, expr)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}

		return ret, nil
	case "assert_unlinkable":
		var result AssertUnlinkableDirective
		err = ps.Parens(result.Module.Parse)
		if err != nil {
			return nil, err
		}
		result.Msg, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "assert_exhaustion":
		var invoke WastInvoke
		err := ps.Parens(invoke.Parse)
		if err != nil {
			return nil, err
		}
		msg, err := ps.ExpectString()
		if err != nil {
			return nil, err
		}
		aed := AssertExhaustionDirective{
			Call:    invoke,
			Message: msg,
		}
		return aed, nil
	default:
		return nil, fmt.Errorf("parse wast directive error: unexpected keyword %s", kw)
	}
}


func parseWastExecute(ps *parser.ParserBuffer) (WastExecute, error) {
	kw, err := ps.ExpectKeyword()
	if err != nil {
		return nil, err
	}

	switch kw {
	case "invoke":
		ps.StepBack(1)
		var invoke WastInvoke
		err := invoke.Parse(ps)
		return invoke, err
	case "module":
		ps.StepBack(1)
		var module Module
		err := module.Parse(ps)
		return module, err
	case "get":
		var get WastExecuteGet
		get.Module.Parse(ps)
		get.Global, err = ps.ExpectString()
		if err != nil {
			return nil, err
		}

		return get, nil
	default:
		return nil, errors.New("parse wast execute error")
	}
}

