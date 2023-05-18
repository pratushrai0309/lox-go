package lox

import "fmt"

type Token struct {
	ttype   TokenType
	lexeme  string
	literal interface{}
	line    int
}

func newToken(ttype TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		ttype:   ttype,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t Token) toString() string {
	return fmt.Sprintf("{ ttype: %d, lexeme: %s, literal: %s, line: %d }", t.ttype, t.lexeme, t.literal, t.line)
}

func eof(line int) Token {
	return Token{
		ttype:   EOF,
		lexeme:  "",
		literal: "",
		line:    line,
	}
}
