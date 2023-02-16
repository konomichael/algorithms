package sort

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func IsSorted(s Sortable) bool {
	n := s.Len()
	for i := n - 1; i > 0; i-- {
		if s.Less(i, i-1) {
			return false
		}
	}
	return true
}
