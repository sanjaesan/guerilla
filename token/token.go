package token

// const -
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

//TType -
type TType string

//Token defines a struct to hold tokens
type Token struct {
	Type    TType
	Literal string
}

// NewToken is a constructor, returns a new token
func NewToken(ttype TType, literal byte) Token {
	return Token{Type: ttype, Literal: string(literal)}
}

var keywords = map[string]TType{
	"func":  FUNCTION,
	"let": LET,
}

// LookUpIdent -
func LookUpIdent(ident string) TType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
