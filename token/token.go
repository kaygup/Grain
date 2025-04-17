package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of fiel

	IDENT = "IDENT" // add, foobar, x y
	INT   = "INT"   // 1 , 2 , 3 , 4 and so on

	ASSIGN = "=" // for giving things value
	PLUS   = "+" // for adding duh
)
