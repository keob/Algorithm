package sort

// SelectionSort => selection sort
func SelectionSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}

	return array
}
