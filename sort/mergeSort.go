package sort

// MergeSort => merge sort
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	var middle = len(array) / 2

	var a = MergeSort(array[:middle])
	var b = MergeSort(array[middle:])

	return merge(a, b)
}

func merge(a, b []int) []int {
	arr := make([]int, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			arr[i+j] = a[i]
			i++
		} else {
			arr[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		arr[i+j] = a[i]
		i++
	}

	for j < len(b) {
		arr[i+j] = b[j]
		j++
	}

	return arr
}
