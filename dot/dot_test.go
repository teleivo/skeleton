package dot_test

import (
	"strings"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/skeleton/dot"
)

func TestParse(t *testing.T) {
	t.Skip()

	t.Run("Graph", func(t *testing.T) {
		tests := map[string]struct {
			in   string
			want *dot.Graph
		}{
			"EmptyGraph": {
				in:   "graph{}",
				want: &dot.Graph{},
			},
			"CaseInsensitiveKeywordGraph": {
				in:   "Graph{}",
				want: &dot.Graph{},
			},
			"EmptyGraphWithWhitespace": {
				in:   "graph {}",
				want: &dot.Graph{},
			},
			"StrictGraph": {
				in:   "strict graph{}",
				want: &dot.Graph{Strict: true},
			},
		}

		for _, test := range tests {
			p := dot.Parser{}

			got := p.Parse(strings.NewReader(test.in))

			assert.EqualValuesf(t, got, test.want, "Parse(%q)", test.in)
		}
	})
}
