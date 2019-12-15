package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	selectionTests := []struct {
		in       []int
		expected []int
	}{
		{[]int{23, 44, 54, 67, 45, 12, 10}, []int{10, 12, 23, 44, 45, 54, 67}},
		{[]int{11, 10, 13, 12, 15, 14}, []int{10, 11, 12, 13, 14, 15}},
	}

	for _, test := range selectionTests {
		actual := SelectionSort(test.in)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Selection(%d) = %d; expected %d", test.in, actual, test.expected)
		}
	}
}
