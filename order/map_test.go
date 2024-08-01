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
		key   int
		value string
	}
	tests := []input{
		{
			key:   10,
			value: "a",
		},
		{
			key:   5,
			value: "b",
		},
		{
			key:   7,
			value: "c",
		},
		{
			key:   15,
			value: "d",
		},
		{
			key:   9,
			value: "e",
		},
		{
			key:   20,
			value: "f",
		},
		{
			key:   6,
			value: "g",
		},
		{
			key:   23,
			value: "h",
		},
		{
			key:   8,
			value: "i",
		},
		{
			key:   2,
			value: "j",
		},
		{
			key:   3,
			value: "k",
		},
		{
			key:   4,
			value: "l",
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
		t.Run("Put/"+strconv.Itoa(test.key), func(t *testing.T) {
			st.Put(test.key, test.value)

			assert.Falsef(t, st.IsEmpty(), "IsEmpty()")

			gotValue, gotOk := st.Get(test.key)
			require.Equalsf(t, gotValue, test.value, "Get(%d)", test.key)
			require.Truef(t, gotOk, "Get(%d)", test.key)

			gotOk = st.Contains(test.key)
			require.Truef(t, gotOk, "Contains(%d)", test.key)

			wantOrder := slices.Clone(tests[:i+1])
			slices.SortFunc(wantOrder, func(a, b input) int {
				return cmp.Compare(a.key, b.key)
			})

			gotKey, gotOk := st.Min()
			require.Equalsf(t, gotKey, wantOrder[0].key, "Min()")
			require.Truef(t, gotOk, "Min()")

			t.Logf("%v", wantOrder)
			var j int
			for gotKey, gotValue := range st.All() {
				t.Logf("%v: %v", gotKey, gotValue)
				require.Equalsf(t, gotKey, wantOrder[j].key, "All()")
				require.Equalsf(t, gotValue, wantOrder[j].value, "All()")
				j++
			}
			require.Equalsf(t, j, i+1, "All() did not return all key and value pairs")
		})
	}
}
