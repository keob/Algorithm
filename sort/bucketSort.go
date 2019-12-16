package sort

func insertionSort(array []float64) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

// BucketSort => bucket sort
func BucketSort(array []float64) []float64 {
	var max, min float64

	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	nBuckets := int(max-min)/len(array) + 1
	buckets := make([][]float64, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, n := range array {
		idx := int(n-min) / len(array)
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}
