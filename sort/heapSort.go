package sort

// HeapSort => heap sort
func HeapSort(array []int) []int {
	heapify(array)
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		siftDown(array, 0, i)
	}

	return array
}

func heapify(array []int) {
	for i := (len(array) - 1) / 2; i >= 0; i-- {
		siftDown(array, i, len(array))
	}
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}
	}
}
