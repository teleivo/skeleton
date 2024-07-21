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
		want *node
	}{
		{
			key: 10,
			want: &node{
				key: 10,
			},
		},
		{
			key: 5,
			want: &node{
				key: 10,
				left: &node{
					red: true,
					key: 5,
				},
			},
		},
		{
			key: 7,
			want: &node{
				key: 7,
				left: &node{
					key: 5,
				},
				right: &node{
					key: 10,
				},
			},
		},
		{
			key: 15,
			want: &node{
				key: 7,
				left: &node{
					key: 5,
				},
				right: &node{
					key: 15,
					left: &node{
						red: true,
						key: 10,
					},
				},
			},
		},
		{
			key: 9,
			want: &node{
				key: 10,
				left: &node{
					red: true,
					key: 7,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 9,
					},
				},
				right: &node{
					key: 15,
				},
			},
		},
		{
			key: 20,
			want: &node{
				key: 10,
				left: &node{
					red: true,
					key: 7,
					left: &node{
						key: 5,
					},
					right: &node{
						key: 9,
					},
				},
				right: &node{
					key: 20,
					left: &node{
						red: true,
						key: 15,
					},
				},
			},
		},
		{
			key: 6,
			want: &node{
				key: 10,
				left: &node{
					red: true,
					key: 7,
					left: &node{
						key: 6,
						left: &node{
							red: true,
							key: 5,
						},
					},
					right: &node{
						key: 9,
					},
				},
				right: &node{
					key: 20,
					left: &node{
						red: true,
						key: 15,
					},
				},
			},
		},
		{
			key: 23,
			want: &node{
				key: 10,
				left: &node{
					key: 7,
					left: &node{
						key: 6,
						left: &node{
							red: true,
							key: 5,
						},
					},
					right: &node{
						key: 9,
					},
				},
				right: &node{
					key: 20,
					left: &node{
						key: 15,
					},
					right: &node{
						key: 23,
					},
				},
			},
		},
		{
			key: 8,
			want: &node{
				key: 10,
				left: &node{
					key: 7,
					left: &node{
						key: 6,
						left: &node{
							red: true,
							key: 5,
						},
					},
					right: &node{
						key: 9,
						left: &node{
							red: true,
							key: 8,
						},
					},
				},
				right: &node{
					key: 20,
					left: &node{
						key: 15,
					},
					right: &node{
						key: 23,
					},
				},
			},
		},
		{
			key: 2,
			want: &node{
				key: 10,
				left: &node{
					key: 7,
					left: &node{
						red: true,
						key: 5,
						left: &node{
							key: 2,
						},
						right: &node{
							key: 6,
						},
					},
					right: &node{
						key: 9,
						left: &node{
							red: true,
							key: 8,
						},
					},
				},
				right: &node{
					key: 20,
					left: &node{
						key: 15,
					},
					right: &node{
						key: 23,
					},
				},
			},
		},
		{
			key: 3,
			want: &node{
				key: 10,
				left: &node{
					key: 7,
					left: &node{
						red: true,
						key: 5,
						left: &node{
							key: 3,
							left: &node{
								red: true,
								key: 2,
							},
						},
						right: &node{
							key: 6,
						},
					},
					right: &node{
						key: 9,
						left: &node{
							red: true,
							key: 8,
						},
					},
				},
				right: &node{
					key: 20,
					left: &node{
						key: 15,
					},
					right: &node{
						key: 23,
					},
				},
			},
		},
		{
			key: 4,
			want: &node{
				key: 10,
				left: &node{
					red: true,
					key: 5,
					left: &node{
						key: 3,
						left: &node{
							key: 2,
						},
						right: &node{
							key: 4,
						},
					},
					right: &node{
						key: 7,
						left: &node{
							key: 6,
						},
						right: &node{
							key: 9,
							left: &node{
								red: true,
								key: 8,
							},
						},
					},
				},
				right: &node{
					key: 20,
					left: &node{
						key: 15,
					},
					right: &node{
						key: 23,
					},
				},
			},
		},
	}

	st := Table{}

	for _, test := range tests {
		st.Put(test.key, 0)

		if diff := cmp.Diff(test.want, st.root, cmp.AllowUnexported(node{})); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}

	}
}
