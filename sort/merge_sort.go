package sort

func MergeSort(s Sortable) {
	mergeSort(s, 0, s.Len())
}

// merge_sort sort s[start, end) in ascending order
func mergeSort(s Sortable, start, end int) {
	if end-start <= 1 {
		return
	}

	size := end - start
	mid := start + size>>1

	mergeSort(s, start, mid)
	mergeSort(s, mid, end)

	merge(s, start, mid, end)
}

func merge(s Sortable, a, m, b int) {
	ln, rn := m-a, b-m
	if ln == 0 || rn == 0 {
		return
	}

	if ln == 1 && rn == 1 {
		if s.Less(m, a) {
			s.Swap(a, m)
		}
		return
	}

	var p, q int

	if ln <= rn {
		q = m + rn>>1
		i, j := a, m
		// find lowest idx such that s[idx] >= s[q]
		// if not found, then idx = m
		for i < j {
			mid := i + (j-i)>>1
			if s.Less(mid, q) {
				i = mid + 1
			} else {
				j = mid
			}
		}
		p = i
	} else {
		p = a + ln>>1
		i, j := m, b
		// find highest idx such that s[idx-1] < s[p]
		// if not found. then idx = m
		for i < j {
			mid := i + (j-i)>>1
			if s.Less(mid, p) {
				i = mid + 1
			} else {
				j = mid
			}
		}
		q = i
	}

	// sort s[p, q) by rotatation
	// if p == m or q == m then rotate twice, which means nothing change
	if p != m && q != m {
		rotate(s, p, m, q)
	}
	m = p + q - m

	merge(s, a, p, m)
	merge(s, m, q, b)
}

func rotate(s Sortable, p, m, q int) {
	reverse(s, p, m)
	reverse(s, m, q)
	reverse(s, p, q)
}

// reverse [a, b)
func reverse(s Sortable, a, b int) {
	for a < b {
		s.Swap(a, b-1)
		a++
		b--
	}
}
