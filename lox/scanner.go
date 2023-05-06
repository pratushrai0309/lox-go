package lox

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func newScanner(source string) Scanner {
	return Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    0,
	}
}

func (s Scanner) scanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, eof(s.line))
	return s.tokens
}

func (s Scanner) scanToken() {

}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
