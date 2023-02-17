package sort

func HeapSort(s Sortable) {
	n := s.Len()
	buildMaxHeap(s, n)
	for n > 1 {
		s.Swap(0, n-1)
		n--
		maxHeapify(s, 0, n)
	}
}

func maxHeapify(s Sortable, i, n int) {
	l, r := i<<1+1, i<<1+2
	largest := i
	if l < n && s.Less(largest, l) {
		largest = l
	}
	if r < n && s.Less(largest, r) {
		largest = r
	}
	if largest != i {
		s.Swap(i, largest)
		maxHeapify(s, largest, n)
	}
}

// buildMaxHeap converts s[:n] into a max-heap
// n - heap size.
func buildMaxHeap(s Sortable, n int) {
	for i := n>>1 - 1; i >= 0; i-- {
		maxHeapify(s, i, n)
	}
}
