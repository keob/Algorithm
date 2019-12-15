package sort

// BubbleSort => bubble sort
func BubbleSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i+1] < array[i] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}

	return array
}
