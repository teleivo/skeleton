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

	t.Run("EmptyMap", func(t *testing.T) {
		m := order.Map[int, string]{}

		assert.True(t, m.IsEmpty(), "IsEmpty()")

		testKey := 27

		gotValue, gotOk := m.Get(testKey)

		assert.Equals(t, gotValue, "", "Get(%d)", testKey)
		assert.False(t, gotOk, "Get(%d)", testKey)

		got := m.Contains(testKey)

		assert.False(t, got, "Contains(%d)", testKey)

		gotKey, gotOk := m.Min()

		assert.Equals(t, gotKey, 0, "Min()")
		assert.False(t, gotOk, "Min()")
	})

	t.Run("Put", func(t *testing.T) {
		m := order.Map[int, string]{}
		for i, test := range tests {
			t.Run(strconv.Itoa(test.Key), func(t *testing.T) {
				m.Put(test.Key, test.Value)

				assert.False(t, m.IsEmpty(), "IsEmpty()")

				gotValue, gotOk := m.Get(test.Key)
				require.Equals(t, gotValue, test.Value, "Get(%d)", test.Key)
				require.True(t, gotOk, "Get(%d)", test.Key)

				gotOk = m.Contains(test.Key)
				require.True(t, gotOk, "Contains(%d)", test.Key)

				wantOrder := slices.Clone(tests[:i+1])
				slices.SortFunc(wantOrder, func(a, b input) int {
					return cmp.Compare(a.Key, b.Key)
				})

				gotKey, gotOk := m.Min()
				require.Equals(t, gotKey, wantOrder[0].Key, "Min()")
				require.True(t, gotOk, "Min()")

				got := make([]input, 0, len(tests))
				for key, value := range m.All() {
					got = append(got, input{Key: key, Value: value})
				}
				require.EqualValues(t, got, wantOrder, "All()")
			})
		}
	})
}
