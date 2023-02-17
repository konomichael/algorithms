package sort

import "math/rand"

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

func generateRandomSlice() []int {
	s := make([]int, 1<<14)
	for i := range s {
		s[i] = rand.Intn(1 << 32)
	}
	return s
}
