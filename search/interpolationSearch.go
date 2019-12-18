package search

// InterpolationSearch => interpolation search.go
func InterpolationSearch(array []int, target int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex && target >= array[minIndex] && target <= array[maxIndex] {
		midIndex := minIndex + (target-array[minIndex])*(maxIndex-minIndex)/(array[maxIndex]-array[minIndex])
		midItem := array[midIndex]
		if midItem == target {
			return midIndex
		} else if midItem < target {
			minIndex = midIndex + 1
		} else if midItem > target {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
