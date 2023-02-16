package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type intSlice []int

func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

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

func generateRandomSlice() []int {
	s := make([]int, 1<<14)
	for i := range s {
		s[i] = rand.Intn(1 << 32)
	}
	return s
}
