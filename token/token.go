// Token unit
package token

type TokenType string

type Token struct {
	Type    TokenType // 类型属性 区分"整数"或者 "(" ")" 等词法单元
	Literal string    // 词法单元字面量
}

// identifier
const (
	IDENT = "IDENT" // like x, y, z, foo, bar...
	INT   = "INT"   // 123456
)

// keywords
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// illegal identifier
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file
)

// separator
const (
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

// operator
const (
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	GT = ">"
	LT = "<"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"if":       IF,
	"else":     ELSE,
	"true":     TRUE,
	"false":    FALSE,
	"return":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
