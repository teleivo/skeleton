package dot

import (
	"strings"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/skeleton/dot/token"
)

func TestLexer(t *testing.T) {
	tests := map[string]struct {
		in   string
		want []token.Token
	}{
		"OnlyWhitespace": {
			in:   "\t \n \t\t   ",
			want: []token.Token{},
		},
		"LiteralSingleCharacterTokens": {
			in: "{};=[],:",
			want: []token.Token{
				{Type: token.LeftBrace, Literal: "{"},
				{Type: token.RightBrace, Literal: "}"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Equal, Literal: "="},
				{Type: token.LeftBracket, Literal: "["},
				{Type: token.RightBracket, Literal: "]"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Colon, Literal: ":"},
			},
		},
		"KeywordsAreCaseInsensitive": {
			in: " graph Graph strict  Strict\ndigraph\tDigraph Subgraph  subgraph Node node edge Edge \n \t ",
			want: []token.Token{
				{Type: token.Graph, Literal: "graph"},
				{Type: token.Graph, Literal: "Graph"},
				{Type: token.Strict, Literal: "strict"},
				{Type: token.Strict, Literal: "Strict"},
				{Type: token.Digraph, Literal: "digraph"},
				{Type: token.Digraph, Literal: "Digraph"},
				{Type: token.Subgraph, Literal: "Subgraph"},
				{Type: token.Subgraph, Literal: "subgraph"},
				{Type: token.Node, Literal: "Node"},
				{Type: token.Node, Literal: "node"},
				{Type: token.Edge, Literal: "edge"},
				{Type: token.Edge, Literal: "Edge"},
			},
		},
		// TODO lex html string
		"Identifiers": { // https://graphviz.org/doc/info/lang.html#ids
			in: `"graph" "strict" "\"d" _A "_A" A_cZ A10 -.9 "-.9" -0.13 -92 -7.3 ÿ 100 200 47`,
			want: []token.Token{
				{Type: token.Identifier, Literal: `"graph"`},
				{Type: token.Identifier, Literal: `"strict"`},
				{Type: token.Identifier, Literal: `"\"d"`},
				{Type: token.Identifier, Literal: "_A"},
				{Type: token.Identifier, Literal: `"_A"`},
				{Type: token.Identifier, Literal: "A_cZ"},
				{Type: token.Identifier, Literal: "A10"},
				{Type: token.Identifier, Literal: "-.9"},
				{Type: token.Identifier, Literal: `"-.9"`},
				{Type: token.Identifier, Literal: "-0.13"},
				{Type: token.Identifier, Literal: "-92"},
				{Type: token.Identifier, Literal: "-7.3"},
				{Type: token.Identifier, Literal: `ÿ`},
				{Type: token.Identifier, Literal: `100 200`}, // non-breakig space \240
				{Type: token.Identifier, Literal: "47"},
			},
		},
		"Subgraphs": {
			in: `  A -> {B C}
				D -- E
			subgraph {
			  rank = same; A; B; C;
			}`,
			want: []token.Token{
				{Type: token.Identifier, Literal: "A"},
				{Type: token.DirectedEgde, Literal: "->"},
				{Type: token.LeftBrace, Literal: "{"},
				{Type: token.Identifier, Literal: "B"},
				{Type: token.Identifier, Literal: "C"},
				{Type: token.RightBrace, Literal: "}"},
				{Type: token.Identifier, Literal: "D"},
				{Type: token.UndirectedEgde, Literal: "--"},
				{Type: token.Identifier, Literal: "E"},
				{Type: token.Subgraph, Literal: "subgraph"},
				{Type: token.LeftBrace, Literal: "{"},
				{Type: token.Identifier, Literal: "rank"},
				{Type: token.Equal, Literal: "="},
				{Type: token.Identifier, Literal: "same"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Identifier, Literal: "A"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Identifier, Literal: "B"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Identifier, Literal: "C"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RightBrace, Literal: "}"},
			},
		},
		// TODO test invalid identifiers, how does any string not leading with a digit concern
		// lexing?
		// TODO test invalid edge operators
		// TODO handle EOF differently? I now have multiple places checking for io.EOF would be nice
		// to mark that in one place
		// TODO add hints to some errors like <- did you mean ->
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lexer := New(strings.NewReader(test.in))

			got := make([]token.Token, 0, len(tests))
			for token, err := range lexer.All() {
				t.Logf("%+v err %v\n", token, err)
				assert.NoError(t, err)
				got = append(got, token)
			}
			assert.EqualValuesf(t, got, test.want, "All(%q)", test.in)
		})
	}
}
