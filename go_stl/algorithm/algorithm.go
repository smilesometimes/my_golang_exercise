package algorithm

// Merge two sorted int slice.
// copy code fromn stl merge.
func Merge(s1 []int, s2 []int, d []int) {
	i, j, k := 0, 0, 0

	for i < len(s1) && j < len(s2) {
		if s2[j] < s1[i] {
			d[k] = s2[j]
			j++
		} else {
			d[k] = s1[i]
			i++
		}
		k++
	}

	//if s1 empty copy from s2
	//else copy from s1
	if i == len(s1) {
		copy(d[k:], s2[j:])
	} else {
		copy(d[k:], s1[i:])
	}
}
