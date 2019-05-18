package token

// MTokenType ...
type MTokenType string

// Token ...
type Token struct {
	Type    MTokenType
	Literal string
}

const (
	// ILLEGAL ...
	ILLEGAL = "ILLEGAL"
	// EOF ...
	EOF = "EOF"

	// 識別子 + リテラル

	// IDENT ...
	IDENT = "IDENT" // add, foobar, x, y, ...
	// INT ...
	INT = "INT" // 123456

	// 演算子

	// ASSIGN ...
	ASSIGN = "="
	// PLUS ...
	PLUS = "+"

	// デリミタ

	// COMMA ...
	COMMA = ","
	// SEMICOLON ...
	SEMICOLON = ";"

	// LPAREN ...
	LPAREN = "(" // left parenthesis
	// RPAREN ...
	RPAREN = ")" // right parenthesis
	// LBRACE ...
	LBRACE = "{" // left brace
	// RBRACE ...
	RBRACE = "}" // right brace

	// キーワード

	// FUNCTION ...
	FUNCTION = "FUNCTION"
	// LET ...
	LET = "LET"
)

var keywords = map[string]MTokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent ...
func LookupIdent(ident string) MTokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
