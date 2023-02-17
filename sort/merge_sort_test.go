package sort

import (
	"sort"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	data := ints
	a := intSlice(data[0:])
	MergeSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestMergeSortLarge(t *testing.T) {
	data := generateRandomSlice()
	a := intSlice(data[0:])
	start := time.Now()
	MergeSort(a)
	elapsed := time.Since(start)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
	t.Logf("Elapsed time: %s\n", elapsed)
}

func BenchmarkMergeSortLarge(b *testing.B) {
	data := generateRandomSlice()
	for i := 0; i < b.N; i++ {
		a := intSlice(data[0:])
		MergeSort(a)
	}
}

func BenchmarkLibSortLarge(b *testing.B) {
	data := generateRandomSlice()
	for i := 0; i < b.N; i++ {
		a := sort.IntSlice(data[0:])
		sort.Sort(a)
	}
}

func TestLibSortLarge(t *testing.T) {
	data := generateRandomSlice()
	a := sort.IntSlice(data[0:])
	start := time.Now()
	sort.Sort(a)
	elapsed := time.Since(start)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
	t.Logf("Elapsed time: %s\n", elapsed)
}
