package sort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	bucketTests := []struct {
		in       []int
		expected []int
	}{
		{[]int{23, 44, 54, 67, 45, 12, 10}, []int{10, 12, 23, 44, 45, 54, 67}},
		{[]int{11, 10, 13, 12, 15, 14}, []int{10, 11, 12, 13, 14, 15}},
	}

	for _, test := range bucketTests {
		actual := BubbleSort(test.in)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("BucketSort(%d) = %d; expected %d\n", test.in, actual, test.expected)
		}
	}
}
