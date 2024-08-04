package token

type TokenType string

const (
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"
	Colon        = ":"
	Semicolon    = ";"
	Equal        = "="
	Comma        = ","

	// Keywords
	Graph   = "graph"
	Digraph = "digraph"
	Strict  = "strict"
)

type Token struct {
	Type    TokenType
	Literal string
}
