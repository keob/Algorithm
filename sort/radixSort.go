package sort

// MaxUint => max uint
const MaxUint = ^uint(0)

// MaxInt => max int
const MaxInt = int(MaxUint >> 1)

// MinInt => min int
const MinInt = -MaxInt - 1

// RadixSort => radix sort
func RadixSort(array []int) []int {
	maxNumber := MinInt

	for _, v := range array {
		if v > maxNumber {
			maxNumber = v
		}
	}

	n := 1
	bucket := make([][]int, 10)
	for n <= maxNumber {
		for _, v := range array {
			bucket[(v/n)%10] = append(bucket[(v/n)%10], v)
		}
		n *= 10

		k := 0
		for i, v := range bucket {
			for _, d := range v {
				array[k] = d
				k++
			}
			bucket[i] = bucket[i][:0]
		}
	}

	return array
}
