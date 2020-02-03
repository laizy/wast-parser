package lexer

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type TokenType byte

const (
	LParenType TokenType = iota
	RParenType
	StringType
	IntegerType
	FloatType
	KeywordType
	IdType
	ReservedType
)

type Lexer struct {
	buf *bytes.Buffer
}

func NewLexer(source string) *Lexer {
	return &Lexer{buf: bytes.NewBufferString(source)}
}

type Token interface {
	String() string
	Type() TokenType
}

type LParen struct{}

func (self LParen) String() string {
	return "("
}

func (self LParen) Type() TokenType {
	return LParenType
}

type String struct {
	Val string
}

func (self String) String() string {
	return fmt.Sprintf("string(%s)", self.Val)
}

func (self String) Type() TokenType {
	return StringType
}

type RParen struct{}

func (self RParen) String() string {
	return ")"
}

func (self RParen) Type() TokenType {
	return RParenType
}

type Identifier struct {
	Val string
}

func (self Identifier) String() string {
	return fmt.Sprintf("id(%s)", self.Val)
}

func (self Identifier) Type() TokenType {
	return IdType
}

type Keyword struct {
	Val string
}

func (self Keyword) String() string {
	return fmt.Sprintf("keyword(%s)", self.Val)
}

func (self Keyword) Type() TokenType {
	return KeywordType
}

type Reserved struct {
	Val string
}

func (self Reserved) Type() TokenType {
	return ReservedType
}

func (self Reserved) String() string {
	return fmt.Sprintf("reserved(%s)", self.Val)
}

type Integer struct {
	Val string
	Hex bool
}

func (self Integer) Type() TokenType {
	return IntegerType
}

func (self Integer) String() string {
	if self.Hex {
		return "0x" + self.Val
	}

	return self.Val
}

func (self *Integer) ToUint(bitSize int) (uint64, error) {
	val, err := self.ToInt(64)
	if err != nil {
		return 0, err
	}

	shift := 64 - uint(bitSize)

	return (uint64(val) << shift) >> shift, nil
}

func (self *Integer) ToInt(bitSize int) (int64, error) {
	base := 10
	if self.Hex {
		base = 16
	}
	if strings.HasPrefix(self.Val, "-") {
		return strconv.ParseInt(self.Val, base, bitSize)
	}
	temp, err := strconv.ParseUint(self.Val, base, bitSize)
	if err != nil {
		return 0, err
	}
	return int64(temp), nil
}

type Float interface {
	Token
	ImplementFloat()
}
type implFloat struct{}

func (self implFloat) ImplementFloat() {}
func (self implFloat) String() string {
	return "float"
}

func (self implFloat) Type() TokenType {
	return FloatType
}

type Nan struct {
	implFloat
	SpecBit bool //whether use val to encode for this float
	Val     uint64
	Neg     bool
}

type Inf struct {
	implFloat
	Neg bool
}

type FloatVal struct {
	implFloat
	Hex      bool
	Integral string
	Decimal  string
	Exponent string
}

func number(num string) Token {
	negative := false
	if strings.HasPrefix(num, "+") {
		num = num[1:]
	} else if strings.HasPrefix(num, "-") {
		negative = true
		num = num[1:]
	}

	skipUnderscore := func(num string, negative bool, valid func(b byte) bool) (string, string) {
		if len(num) == 0 {
			return "", ""
		}
		var result []byte
		if negative {
			result = []byte("-")
		}
		lastUnderscore := false
		if !valid(num[0]) {
			return "", ""
		}
		result = append(result, num[0])
		num = num[1:]
		last := len(num)
		for i, val := range []byte(num) {
			if val == '_' && !lastUnderscore {
				lastUnderscore = true
				continue
			}

			if !valid(val) {
				last = i
				break
			}

			lastUnderscore = false
			result = append(result, val)
		}

		if lastUnderscore {
			return "", ""
		}

		return num[last:], string(result)
	}

	if num == "inf" {
		return Inf{Neg: negative}
	} else if num == "nan" {
		return Nan{Neg: negative}
	} else if strings.HasPrefix(num, "nan:0x") {
		left, val := skipUnderscore(num[6:], false, validHexDigit)
		if len(left) != 0 {
			return nil
		}
		val = strings.ToLower(val)
		n, err := strconv.ParseUint(val, 16, 64)
		if err != nil {
			return nil
		}
		return Nan{Neg: negative, Val: n, SpecBit: true}
	}

	hex := false
	valid := validDigit
	if strings.HasPrefix(num, "0x") {
		num = num[2:]
		hex = true
		valid = validHexDigit
	}

	n, val := skipUnderscore(num, negative, valid)
	if val == "" {
		return nil
	}
	num = n
	integral := val

	if num == "" {
		return Integer{Hex: hex, Val: integral}
	}

	decimal := ""
	if num[0] == '.' && len(num) > 1 && valid(num[1]) {
		num, decimal = skipUnderscore(num[1:], false, valid)
		if val == "" {
			return nil
		}
	}

	exponent := ""
	if len(num) > 0 {
		if (hex && (num[0] == 'p' || num[0] == 'P')) || (!hex && (num[0] == 'e' || num[0] == 'E')) {
			neg := false
			num = num[1:]
			if len(num) > 0 {
				if num[0] == '-' {
					neg = true
					num = num[1:]
				} else if num[0] == '+' {
					num = num[1:]
				}
			}

			num, val = skipUnderscore(num, neg, validDigit)
			if val == "" {
				return nil
			}
			exponent = val
		}
	}

	if num != "" {
		return nil
	}

	return FloatVal{
		Hex:      hex,
		Integral: integral,
		Decimal:  decimal,
		Exponent: exponent,
	}
}

func validDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func validHexDigit(b byte) bool {
	return validDigit(b) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}

func (self *Lexer) PeekChar() (rune, error) {
	r, _, err := self.buf.ReadRune()
	if err != nil {
		return r, err
	}
	_ = self.buf.UnreadRune()

	return r, nil
}

func (self *Lexer) Eof() bool {
	return self.buf.Len() == 0
}

func (self *Lexer) StartWith(pref string) bool {
	return bytes.HasPrefix(self.buf.Bytes(), []byte(pref))
}

func (self *Lexer) ReadChar() (rune, error) {
	r, _, err := self.buf.ReadRune()
	if err != nil {
		return r, err
	}

	return r, nil
}

func (self *Lexer) ReadByte() (byte, error) {
	return self.buf.ReadByte()
}

// peek the next byte
func (self *Lexer) NextByte() byte {
	b, err := self.buf.ReadByte()
	if err != nil {
		panic(err)
	}

	_ = self.buf.UnreadByte()

	return b
}

func (self *Lexer) ReadWhile(test func(byte) bool) string {
	var buf []byte
	for b, err := self.ReadByte(); err == nil; b, err = self.ReadByte() {
		if !test(b) {
			_ = self.buf.UnreadByte()
			break
		}

		buf = append(buf, b)
	}

	return string(buf)
}

func (self *Lexer) SkipWhiteSpace() bool {
	str := self.ReadWhile(func(b byte) bool {
		return b == ' ' || b == '\n' || b == '\t' || b == '\r'
	})

	return len(str) != 0
}

func (self *Lexer) SkipComment() bool {
	skipped := false
	for {
		if self.SkipPrefix(";;") {
			self.ReadWhile(func(b byte) bool {
				return b != '\n'
			})
			self.SkipPrefix("\n")
			skipped = true
			continue
		} else if self.SkipPrefix("(;") {
			level := 1
			finished := false
			self.ReadWhile(func(b byte) bool {
				if finished {
					return false
				}
				if b == '(' && self.SkipPrefix(";") {
					level += 1
				}
				if b == ';' && self.SkipPrefix(")") {
					level -= 1
					if level == 0 {
						finished = true
					}
				}
				return true
			})
			skipped = true
			continue
		}

		break
	}

	return skipped
}

func (self *Lexer) SkipPrefix(pref string) bool {
	b := self.StartWith(pref)
	if b {
		self.buf.Next(len(pref))
	}

	return b
}

func (self *Lexer) Parse() (Token, error) {
	skipped := true
	for skipped {
		skipped = self.SkipWhiteSpace() || self.SkipComment()
	}

	if self.Eof() {
		return nil, io.EOF
	}

	return self.ReadToken()
}

func (self *Lexer) ReadToken() (Token, error) {
	if self.SkipPrefix("(") {
		return LParen{}, nil
	} else if self.SkipPrefix(")") {
		return RParen{}, nil
	} else if self.SkipPrefix("\"") {
		str, err := self.ReadStringToken()
		if err != nil {
			return nil, err
		}

		return String{Val: str}, nil
	}

	str := self.ReadWhile(isIdChar)
	if len(str) == 0 {
		return nil, fmt.Errorf("unexpected bytes: %s", self.buf.Bytes())
	}

	if token := number(str); token != nil {
		return token, nil
	} else if str[0] == '$' && len(str) > 1 {
		return Identifier{Val: str}, nil
	} else if str[0] >= 'a' && str[0] <= 'z' {
		return Keyword{Val: str}, nil
	} else {
		return Reserved{Val: str}, nil
	}
}

func (self *Lexer) ReadStringToken() (string, error) {
	var result string
	for {
		c, err := self.ReadChar()
		if err != nil {
			return "", err
		}
		switch c {
		case '\\':
			ecs, err := self.ReadChar()
			if err != nil {
				return "", err
			}
			switch ecs {
			case '"':
				result += string('"')
			case '\'':
				result += string('\'')
			case 't':
				result += string('\t')
			case 'n':
				result += string('\n')
			case 'r':
				result += string('\r')
			case '\\':
				result += string('\\')
			case 'u':
				if !self.SkipPrefix("{") {
					return "", fmt.Errorf("expected start with {")
				}
				n, err := self.hexnum()
				if err != nil {
					return "", err
				}
				result += string(n)
				if self.SkipPrefix("}") {
					return "", fmt.Errorf("expected end with }")
				}
			default:
				if validHexDigit(byte(ecs)) {
					c2, err := self.hexdigit()
					if err != nil {
						return "", err
					}
					result += string(to_hex(ecs)*16 + c2)
				} else {
					return "", fmt.Errorf("UnexpectedEof")
				}
			}
		case '"':
			return result, nil
		default:
			if c < 0x20 || c == 0x7f {
				return "", fmt.Errorf("invalid string element %v", c)
			}
			result += string(c)
		}
	}
}

func (self *Lexer) hexnum() (uint32, error) {
	n, err := self.hexdigit()
	if err != nil {
		return 0, err
	}
	lastUnderscore := false
	for {
		c, err := self.ReadChar()
		if err != nil {
			return 0, err
		}
		if c == '_' {
			lastUnderscore = true
			continue
		}
		if !validHexDigit(byte(c)) {
			break
		}
		lastUnderscore = false
		n = n*16 + to_hex(c)
	}
	if lastUnderscore {
		return 0, fmt.Errorf("LoneUnderscore")
	}
	return uint32(n), nil
}

func (self *Lexer) hexdigit() (byte, error) {
	ch, err := self.ReadChar()
	if err != nil {
		return 0, err
	}
	if validHexDigit(byte(ch)) {
		return to_hex(ch), nil
	}
	return 0, fmt.Errorf("InvalidHexDigit")
}

func to_hex(c rune) byte {
	if 'a' <= c && c <= 'f' {
		return byte(c) - 'a' + 10
	} else if 'A' <= c && c <= 'F' {
		return byte(c) - 'A' + 10
	} else {
		return byte(c) - '0'
	}
}
func isIdChar(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	} else if b >= 'a' && b <= 'z' {
		return true
	} else if b >= 'A' && b <= 'Z' {
		return true
	}
	switch b {
	case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '.', '/', ':', '<',
		'=', '>', '?', '@', '\\', '^', '_', '`', ',', '~':
		return true
	default:
		return false
	}

}
