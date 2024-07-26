package order_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
	"github.com/teleivo/skeleton/order"
)

func TestTable(t *testing.T) {
	tests := []struct {
		key       int
		operation string
	}{
		{
			key:       10,
			operation: "Put",
		},
		{
			key:       5,
			operation: "Put",
		},
		{
			key:       7,
			operation: "Put",
		},
		{
			key:       15,
			operation: "Put",
		},
		{
			key:       9,
			operation: "Put",
		},
		{
			key:       20,
			operation: "Put",
		},
		{
			key:       6,
			operation: "Put",
		},
		{
			key:       23,
			operation: "Put",
		},
		{
			key:       8,
			operation: "Put",
		},
		{
			key:       2,
			operation: "Put",
		},
		{
			key:       3,
			operation: "Put",
		},
		{
			key:       4,
			operation: "Put",
		},
	}

	st := order.Table[int, int]{}

	t.Run("IsEmtpy/EmtpyTable", func(t *testing.T) {
		require.Truef(t, st.IsEmpty(), "IsEmpty()")
	})

	t.Run("Get/EmptyTable", func(t *testing.T) {
		testKey := 27

		gotValue, gotOk := st.Get(testKey)

		require.Equalsf(t, gotValue, 0, "Get(%d)", testKey)
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

	wantMin := math.MaxInt
	for _, test := range tests {
		t.Run(test.operation+"/"+strconv.Itoa(test.key), func(t *testing.T) {
			switch test.operation {
			case "Put":
				testValue := 1

				st.Put(test.key, testValue)

				assert.Falsef(t, st.IsEmpty(), "IsEmpty()")

				gotValue, gotOk := st.Get(test.key)
				require.Equalsf(t, gotValue, testValue, "Get(%d)", test.key)
				require.Truef(t, gotOk, "Get(%d)", test.key)

				gotOk = st.Contains(test.key)
				require.Truef(t, gotOk, "Contains(%d)", test.key)

				gotKey, gotOk := st.Min()
				wantMin = min(wantMin, test.key)
				require.Equalsf(t, gotKey, wantMin, "Min()")
				require.Truef(t, gotOk, "Min()")
			}
		})
	}
}
