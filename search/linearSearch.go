package search

// LinearSearch => linear search
func LinearSearch(array []int, target int) int {
	for i, n := range array {
		if n == target {
			return i
		}
	}
	return -1
}
