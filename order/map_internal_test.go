package order

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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

	st := Map[int, int]{}

	for _, test := range tests {
		st.Put(test.key, 0)

		if diff := cmp.Diff(test.want, st.root, cmp.AllowUnexported(node[int, int]{})); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}

	}
}
