package lexer

import (
	"unicode"

	"github.com/guerilla/token"
)

// Lexer -
type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

// New returns a new instance of the lexer object
func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

// ReadChar gives the next character and advance the  lexer input field position
func (s *Lexer) ReadChar() {
	if s.readPosition >= len(s.input) {
		s.char = 0
	} else {
		s.char = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition++
}

//NextToken -
func (s *Lexer) NextToken() token.Token {
	var tok token.Token
	s.skipWhitespace()
	switch s.char {
	case ')':
		tok = token.NewToken(token.RPAREN, s.char)
	case '(':
		tok = token.NewToken(token.LPAREN, s.char)
	case '+':
		tok = token.NewToken(token.PLUS, s.char)
	case ',':
		tok = token.NewToken(token.COMMA, s.char)
	case '=':
		tok = token.NewToken(token.ASSIGN, s.char)
	case ';':
		tok = token.NewToken(token.SEMICOLON, s.char)
	case '}':
		tok = token.NewToken(token.RBRACE, s.char)
	case '{':
		tok = token.NewToken(token.LBRACE, s.char)
	default:
		if unicode.IsLetter(rune(s.char)) {
			tok.Literal = s.readIdentifier()
			return tok
		} else if unicode.IsDigit(rune(s.char)) {
			tok.Literal = s.readNumber()
			return tok
		}
		tok = token.NewToken(token.ILLEGAL, s.char)
	}
	s.ReadChar()
	return tok
}

//readIdentifier -
func (s *Lexer) readIdentifier() string {
	pos := s.position
	for unicode.IsLetter(rune(s.char)) {
		s.ReadChar()
	}
	ret := s.input[pos:s.position]
	return ret
}

// readNumber -
func (s *Lexer) readNumber() string {
	pos := s.position
	for unicode.IsNumber(rune(s.char)) {
		s.ReadChar()
	}
	return s.input[pos:s.position]
}

func (s *Lexer) skipWhitespace() {
	for s.char == ' ' || s.char == '\t' || s.char == '\n' || s.char == '\r' {
		s.ReadChar()
	}
}
