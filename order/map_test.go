package order_test

import (
	"cmp"
	"slices"
	"strconv"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
	"github.com/teleivo/skeleton/order"
)

func TestMap(t *testing.T) {
	type input struct {
		Key   int
		Value string
	}
	tests := []input{
		{
			Key:   10,
			Value: "a",
		},
		{
			Key:   5,
			Value: "b",
		},
		{
			Key:   7,
			Value: "c",
		},
		{
			Key:   15,
			Value: "d",
		},
		{
			Key:   9,
			Value: "e",
		},
		{
			Key:   20,
			Value: "f",
		},
		{
			Key:   6,
			Value: "g",
		},
		{
			Key:   23,
			Value: "h",
		},
		{
			Key:   8,
			Value: "i",
		},
		{
			Key:   2,
			Value: "j",
		},
		{
			Key:   3,
			Value: "k",
		},
		{
			Key:   4,
			Value: "l",
		},
	}

	st := order.Map[int, string]{}

	t.Run("IsEmpty/EmptyMap", func(t *testing.T) {
		require.Truef(t, st.IsEmpty(), "IsEmpty()")
	})

	t.Run("Get/EmptyMap", func(t *testing.T) {
		testKey := 27

		gotValue, gotOk := st.Get(testKey)

		require.Equalsf(t, gotValue, "", "Get(%d)", testKey)
		require.Falsef(t, gotOk, "Get(%d)", testKey)
	})

	t.Run("Contains/EmptyMap", func(t *testing.T) {
		testKey := 27

		got := st.Contains(testKey)

		require.Falsef(t, got, "Contains(%d)", testKey)
	})

	t.Run("Min/EmptyMap", func(t *testing.T) {
		gotKey, gotOk := st.Min()

		require.Equalsf(t, gotKey, 0, "Min()")
		require.Falsef(t, gotOk, "Min()")
	})

	for i, test := range tests {
		t.Run("Put/"+strconv.Itoa(test.Key), func(t *testing.T) {
			st.Put(test.Key, test.Value)

			assert.Falsef(t, st.IsEmpty(), "IsEmpty()")

			gotValue, gotOk := st.Get(test.Key)
			require.Equalsf(t, gotValue, test.Value, "Get(%d)", test.Key)
			require.Truef(t, gotOk, "Get(%d)", test.Key)

			gotOk = st.Contains(test.Key)
			require.Truef(t, gotOk, "Contains(%d)", test.Key)

			wantOrder := slices.Clone(tests[:i+1])
			slices.SortFunc(wantOrder, func(a, b input) int {
				return cmp.Compare(a.Key, b.Key)
			})

			gotKey, gotOk := st.Min()
			require.Equalsf(t, gotKey, wantOrder[0].Key, "Min()")
			require.Truef(t, gotOk, "Min()")

			got := make([]input, 0, len(tests))
			for key, value := range st.All() {
				got = append(got, input{Key: key, Value: value})
			}
			require.EqualValuesf(t, got, wantOrder, "All()")
		})
	}
}
