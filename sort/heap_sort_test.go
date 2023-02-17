package sort

import "testing"

func TestHeapSort(t *testing.T) {
	data := ints
	a := intSlice(data[0:])
	HeapSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}
