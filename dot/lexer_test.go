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
		"LiteralCharacterTokens": {
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
		// "Keywords": {
		// 	in: "graph Graph strict  Strict\ndigraph\tDigraph",
		// 	want: []token.Token{
		// 		{Type: token.Graph, Literal: "graph"},
		// 		{Type: token.Graph, Literal: "Graph"},
		// 		{Type: token.Strict, Literal: "strict"},
		// 		{Type: token.Strict, Literal: "Strict"},
		// 		{Type: token.Digraph, Literal: "digraph"},
		// 		{Type: token.Digraph, Literal: "Digraph"},
		// 	},
		// },
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lexer := New(strings.NewReader(test.in))

			got := make([]token.Token, 0, len(tests))
			for token, err := range lexer.All() {
				assert.NoError(t, err)
				got = append(got, token)
			}
			assert.EqualValuesf(t, got, test.want, "All(%q)", test.in)
		})
	}
}
