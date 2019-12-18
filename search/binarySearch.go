package search

// BinarySearch => binary search
func BinarySearch(array []int, target int) int {
	startIndex := 0
	endIndex := len(array) - 1
	midIndex := int((startIndex + endIndex) / 2)

	for startIndex <= endIndex {
		mid := array[midIndex]

		if mid == target {
			return midIndex
		}

		if mid > target {
			endIndex = midIndex - 1
			midIndex = (startIndex + endIndex) / 2
			continue
		}

		startIndex = midIndex + 1
		midIndex = (startIndex + endIndex) / 2
	}
	return -1
}
