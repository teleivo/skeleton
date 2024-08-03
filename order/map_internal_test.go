package order

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/teleivo/assertive/require"
)

func TestRenderDot(t *testing.T) {
	m := Map[int, int]{
		root: &node[int, int]{
			key: 10,
			left: &node[int, int]{
				red: true,
				key: 5,
				left: &node[int, int]{
					key: 3,
					left: &node[int, int]{
						key: 2,
					},
					right: &node[int, int]{
						key: 4,
					},
				},
				right: &node[int, int]{
					key: 7,
					left: &node[int, int]{
						key: 6,
					},
					right: &node[int, int]{
						key: 9,
						left: &node[int, int]{
							red: true,
							key: 8,
						},
					},
				},
			},
			right: &node[int, int]{
				key: 20,
				left: &node[int, int]{
					key: 15,
				},
				right: &node[int, int]{
					key: 23,
				},
			},
		},
	}

	var b bytes.Buffer
	m.RenderDot(&b)

	want := `strict digraph {
	3 -> 2 [label="L"]
	5 -> 3 [label="L"]
	3 -> 4 [label="R"]
	10 -> 5 [label="L", color = red]
	7 -> 6 [label="L"]
	5 -> 7 [label="R"]
	9 -> 8 [label="L", color = red]
	7 -> 9 [label="R"]
	20 -> 15 [label="L"]
	10 -> 20 [label="R"]
	20 -> 23 [label="R"]
}`
	require.EqualValues(t, b.String(), want)
}

func TestRotationsAndFlip(t *testing.T) {
	// testing tree transformations on put via diffing on the nodes with its
	// internals as the node color is important in the LLRB
	tests := []struct {
		key  int
		want *node[int, int]
	}{
		{
			key: 10,
			want: &node[int, int]{
				key: 10,
			},
		},
		{
			key: 5,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					red: true,
					key: 5,
				},
			},
		},
		{
			key: 7,
			want: &node[int, int]{
				key: 7,
				left: &node[int, int]{
					key: 5,
				},
				right: &node[int, int]{
					key: 10,
				},
			},
		},
		{
			key: 15,
			want: &node[int, int]{
				key: 7,
				left: &node[int, int]{
					key: 5,
				},
				right: &node[int, int]{
					key: 15,
					left: &node[int, int]{
						red: true,
						key: 10,
					},
				},
			},
		},
		{
			key: 9,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					red: true,
					key: 7,
					left: &node[int, int]{
						key: 5,
					},
					right: &node[int, int]{
						key: 9,
					},
				},
				right: &node[int, int]{
					key: 15,
				},
			},
		},
		{
			key: 20,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					red: true,
					key: 7,
					left: &node[int, int]{
						key: 5,
					},
					right: &node[int, int]{
						key: 9,
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						red: true,
						key: 15,
					},
				},
			},
		},
		{
			key: 6,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					red: true,
					key: 7,
					left: &node[int, int]{
						key: 6,
						left: &node[int, int]{
							red: true,
							key: 5,
						},
					},
					right: &node[int, int]{
						key: 9,
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						red: true,
						key: 15,
					},
				},
			},
		},
		{
			key: 23,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					key: 7,
					left: &node[int, int]{
						key: 6,
						left: &node[int, int]{
							red: true,
							key: 5,
						},
					},
					right: &node[int, int]{
						key: 9,
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						key: 15,
					},
					right: &node[int, int]{
						key: 23,
					},
				},
			},
		},
		{
			key: 8,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					key: 7,
					left: &node[int, int]{
						key: 6,
						left: &node[int, int]{
							red: true,
							key: 5,
						},
					},
					right: &node[int, int]{
						key: 9,
						left: &node[int, int]{
							red: true,
							key: 8,
						},
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						key: 15,
					},
					right: &node[int, int]{
						key: 23,
					},
				},
			},
		},
		{
			key: 2,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					key: 7,
					left: &node[int, int]{
						red: true,
						key: 5,
						left: &node[int, int]{
							key: 2,
						},
						right: &node[int, int]{
							key: 6,
						},
					},
					right: &node[int, int]{
						key: 9,
						left: &node[int, int]{
							red: true,
							key: 8,
						},
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						key: 15,
					},
					right: &node[int, int]{
						key: 23,
					},
				},
			},
		},
		{
			key: 3,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					key: 7,
					left: &node[int, int]{
						red: true,
						key: 5,
						left: &node[int, int]{
							key: 3,
							left: &node[int, int]{
								red: true,
								key: 2,
							},
						},
						right: &node[int, int]{
							key: 6,
						},
					},
					right: &node[int, int]{
						key: 9,
						left: &node[int, int]{
							red: true,
							key: 8,
						},
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						key: 15,
					},
					right: &node[int, int]{
						key: 23,
					},
				},
			},
		},
		{
			key: 4,
			want: &node[int, int]{
				key: 10,
				left: &node[int, int]{
					red: true,
					key: 5,
					left: &node[int, int]{
						key: 3,
						left: &node[int, int]{
							key: 2,
						},
						right: &node[int, int]{
							key: 4,
						},
					},
					right: &node[int, int]{
						key: 7,
						left: &node[int, int]{
							key: 6,
						},
						right: &node[int, int]{
							key: 9,
							left: &node[int, int]{
								red: true,
								key: 8,
							},
						},
					},
				},
				right: &node[int, int]{
					key: 20,
					left: &node[int, int]{
						key: 15,
					},
					right: &node[int, int]{
						key: 23,
					},
				},
			},
		},
	}

	m := Map[int, int]{}

	for _, test := range tests {
		t.Logf("Put(%v)", test.key)

		m.Put(test.key, 0)

		if testing.Verbose() {
			var b bytes.Buffer
			m.RenderDot(&b)
			t.Logf("maps' internal tree in dot representation\n%s", b.String())
		}

		if diff := cmp.Diff(test.want, m.root, cmp.AllowUnexported(node[int, int]{})); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}
	}
}
