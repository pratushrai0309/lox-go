package lox

type TokenType int

const (
	//Single-Character Tokens
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	SemiColon
	Slash
	Star

	//One or two character Tokens
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	//Literals
	Indentifier
	String
	Number

	//Keywords
	And
	Class
	Else
	False
	Fun
	For
	IF
	NIL
	Or
	Print
	Return
	Super
	This
	True
	Let
	While

	EOF
)
