package search

import "math"

// JumpSearch => jump search
func JumpSearch(array []int, target int) int {
	step := int(math.Floor(math.Sqrt(float64(len(array)))))
	minIndex := 0
	maxIndex := step

	for array[maxIndex] <= target {
		minIndex += step
		maxIndex = minIndex + step
		if maxIndex >= len(array) {
			maxIndex = len(array)
			minIndex = maxIndex - step
			break
		}
	}

	for i := minIndex; i < maxIndex; i++ {
		if array[i] == target {
			return i
		}
	}

	return -1
}
