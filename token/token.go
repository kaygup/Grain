package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	SOW     = "SOW"
	PLANT   = "PLANT"
	HARVEST = "HARVEST"
	SPROUT  = "SPROUT"
	WILT    = "WILT"
	LOOP    = "LOOP"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"sow":     SOW,
	"plant":   PLANT,
	"harvest": HARVEST,
	"sprout":  SPROUT,
	"wilt":    WILT,
	"loop":    LOOP,
	"true":    TRUE,
	"false":   FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
