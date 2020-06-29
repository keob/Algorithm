package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	cases := []struct {
		in       []int
		target   int
		expected int
	}{
		{[]int{2, 3, 4, 10, 40}, 10, 3},
		{[]int{2, 3, 4, 10, 40}, 1, -1},
		{[]int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}, 110, 6},
		{[]int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}, 100, 7},
	}

	for _, test := range cases {
		actual := LinearSearch(test.in, test.target)
		if actual != test.expected {
			t.Errorf("LinearSearch(%d, %d) = %d; expected %d\n", test.in, test.target, actual, test.expected)
		}
	}
}
