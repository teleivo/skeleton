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

	Identifier = "identifier"

	// Keywords
	Digraph  = "digraph"
	Edge     = "edge"
	Graph    = "graph"
	Node     = "node"
	Strict   = "strict"
	Subgraph = "subgraph"
)

// Token represents a token of the dot language.
type Token struct {
	Type    TokenType
	Literal string
}

// MaxKeywordLen is the length of the longest dot keyword which is "subgraph".
const MaxKeywordLen = 8

var keywords = map[string]TokenType{
	"digraph":  Digraph,
	"edge":     Edge,
	"graph":    Graph,
	"node":     Node,
	"strict":   Strict,
	"subgraph": Subgraph,
}

// LookupIdentifier returns the token type associated with given identifier which is either a dot
// keyword or a dot id. This function expects that the input is a valid dot id as specified in
// https://graphviz.org/doc/info/lang.html#ids.
func LookupIdentifier(identifier string) TokenType {
	v, ok := keywords[identifier]
	if ok {
		return v
	}

	return Identifier
}
