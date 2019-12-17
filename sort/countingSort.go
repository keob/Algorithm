package sort

const maxUint = ^uint(0)
const maxInt = int(MaxUint >> 1)
const minInt = -maxInt - 1

// CountingSort => counting sort
func CountingSort(array []int) []int {
	maxNumber := minInt
	minNumber := maxInt

	for _, v := range array {
		if v > maxNumber {
			maxNumber = v
		}
		if v < minNumber {
			minNumber = v
		}
	}

	count := make([]int, maxNumber-minNumber+1)

	for _, x := range array {
		count[x-minNumber]++
	}

	index := 0
	for i, c := range count {
		for ; c > 0; c-- {
			array[index] = i + minNumber
			index++
		}
	}

	return array
}
