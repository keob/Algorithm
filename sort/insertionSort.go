package sort

// InsertionSort => insertion sort
func InsertionSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			array[j-1], array[j] = array[j], array[j-1]
			j--
		}
	}

	return array
}
