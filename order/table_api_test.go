package order_test

import (
	"strconv"
	"testing"

	"github.com/teleivo/assertive/require"
	"github.com/teleivo/skeleton/order"
)

func TestPutAndGet(t *testing.T) {
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

	st := order.Table{}

	for _, test := range tests {
		t.Run(test.operation+"/"+strconv.Itoa(test.key), func(t *testing.T) {
			switch test.operation {
			case "Put":
				st.Put(test.key, 1)

				gotValue, gotOk := st.Get(test.key)
				require.Truef(t, gotOk, "Get(%d)", test.key)
				require.Equalsf(t, gotValue, 1, "Get(%d)", test.key, 1)

				got := st.Contains(test.key)
				require.Truef(t, got, "Contains(%d)", test.key)
			}
		})
	}
}
