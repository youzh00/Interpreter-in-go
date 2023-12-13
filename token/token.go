package token

// Define the token type as string, that allows us to many differents values as TokenType.
// Also easy to debug, just print a string.
// But not the same performace as using an Int or Byte.
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// As we have a limited number of different token types, then we can define them as constats

const (
	//signifies a charachter we do not know about
	ILLEGAL = "ILLEGAL"
	// EOF means end of file, which will tell the parser when to stop
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
