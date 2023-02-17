package sort

import "testing"

func TestQuickSort(t *testing.T) {
	data := generateRandomSlice()
	a := intSlice(data[0:])
	QuickSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}
