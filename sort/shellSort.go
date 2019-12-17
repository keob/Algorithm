package sort

// ShellSort => shell sort
func ShellSort(array []int) []int {
	h := 1
	for h < len(array) {
		h = h*3 + 1
	}

	for h >= 1 {
		for i := h; i < len(array); i++ {
			for j := i; j >= h && array[j] < array[j-h]; j -= h {
				array[j], array[j-h] = array[j-h], array[j]
			}
		}
		h /= 3
	}

	return array
}
