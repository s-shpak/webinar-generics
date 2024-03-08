package intro

import (
	"fmt"
	"testing"
)

func TestPrintNumber(t *testing.T) {
	cases := []struct {
		N        int
		Expected string
	}{
		{
			N:        16,
			Expected: "16",
		},
		{
			N:        -16,
			Expected: "-16",
		},
	}

	for i, tc := range cases {
		i, tc := i, tc
		t.Run(fmt.Sprintf("test case #%d", i), func(t *testing.T) {
			actual := PrintNumber(tc.N)
			if actual != tc.Expected {
				t.Errorf("expected %s, got %s", tc.Expected, actual)
				return
			}
		})
	}
}
