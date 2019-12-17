package sort

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	radixTests := []struct {
		in       []int
		expected []int
	}{
		{[]int{43, 123, 51, 5, 1, 4, 7, 234, 6123, 5123, 11, 23, 5123}, []int{1, 4, 5, 7, 11, 23, 43, 51, 123, 234, 5123, 5123, 6123}},
		{[]int{23, 44, 54, 67, 45, 12, 10}, []int{10, 12, 23, 44, 45, 54, 67}},
		{[]int{11, 10, 13, 12, 15, 14}, []int{10, 11, 12, 13, 14, 15}},
	}

	for _, test := range radixTests {
		actual := RadixSort(test.in)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("RadixSort(%d) = %d; expected %d\n", test.in, actual, test.expected)
		}
	}
}
