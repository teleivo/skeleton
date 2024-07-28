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

func TestTable(t *testing.T) {
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

	st := order.Table[int, string]{}

	t.Run("IsEmtpy/EmtpyTable", func(t *testing.T) {
		require.Truef(t, st.IsEmpty(), "IsEmpty()")
	})

	t.Run("Get/EmptyTable", func(t *testing.T) {
		testKey := 27

		gotValue, gotOk := st.Get(testKey)

		require.Equalsf(t, gotValue, "", "Get(%d)", testKey)
		require.Falsef(t, gotOk, "Get(%d)", testKey)
	})

	t.Run("Contains/EmptyTable", func(t *testing.T) {
		testKey := 27

		got := st.Contains(testKey)

		require.Falsef(t, got, "Contains(%d)", testKey)
	})

	t.Run("Min/EmptyTable", func(t *testing.T) {
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
			for gotKey, gotValue := range st.Iterate() {
				t.Logf("%v: %v", gotKey, gotValue)
				require.Equalsf(t, gotKey, wantOrder[j].key, "Iterate()")
				require.Equalsf(t, gotValue, wantOrder[j].value, "Iterate()")
				j++
			}
			require.Equalsf(t, j, i, "Iterate() did not return all key and value pairs")
		})
	}

	// for _, wantMin := range wantOrder {
	// 	t.Run("DeleteMin/"+strconv.Itoa(wantMin.key), func(t *testing.T) {
	// 		gotMinKey, gotOk := st.Min()
	// 		require.Equalsf(t, gotMinKey, wantMin.key, "Min()")
	// 		require.Truef(t, gotOk, "Min()")
	//
	// 		gotKey, gotValue, gotOk := st.DeleteMin()
	// 		require.Equalsf(t, gotKey, wantMin.key, "DeleteMin()")
	// 		require.Equalsf(t, gotValue, wantMin.value, "DeleteMin()")
	// 		require.Truef(t, gotOk, "DeleteMin()")
	// 	})
	// }
}
