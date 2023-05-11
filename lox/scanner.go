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
	var c rune = s.advance()
	switch c {
	case '(':
		s.addToken(LeftParen, "")
	case ')':
		s.addToken(RightParen, "")
	case '{':
		s.addToken(LeftBrace, "")
	case '}':
		s.addToken(LeftBrace, "")
	case ',':
		s.addToken(Comma, "")
	case '.':
		s.addToken(Dot, "")
	case '-':
		s.addToken(Minus, "")
	case '+':
		s.addToken(Plus, "")
	case ';':
		s.addToken(SemiColon, "")
	case '*':
		s.addToken(Star, "")
	case '!':
		if s.match('=') {
			s.addToken(BangEqual, "")
		} else {
			s.addToken(Bang, "")
		}
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual, "")
		} else {
			s.addToken(Equal, "")
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual, "")
		} else {
			s.addToken(Less, "")
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual, "")
		} else {
			s.addToken(LessEqual, "")
		}
	default:
		error(s.line, "Unexpected character.")
	}

}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s Scanner) advance() rune {
	nextPos := s.current + 1
	return rune(s.source[nextPos])
}

func (s Scanner) addToken(ttype TokenType, literal string) {
	var text string = s.source[s.start:s.current]
	s.tokens = append(s.tokens, newToken(ttype, text, literal, s.line))
}

func (s Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if rune(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}
