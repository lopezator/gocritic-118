package slice

// Difference compares the first slice against the second and returns the values
// that are not present in the second.
func Difference[V comparable](a, b []V) []V {
	m := make(map[V]bool)

	for _, item := range b {
		m[item] = true
	}

	var diff []V
	for _, v := range a {
		if _, ok := m[v]; !ok {
			diff = append(diff, v)
		}
	}

	return diff
}

// Intersect computes the intersection of two slices.
func Intersect[V comparable](a, b []V) []V {
	am := make(map[V]bool)
	for _, v := range a {
		am[v] = true
	}

	im := make(map[V]bool)
	for _, v := range b {
		if _, ok := am[v]; ok {
			im[v] = true
		}
	}

	index := 0
	intersect := make([]V, len(im))
	for k := range im {
		intersect[index] = k
		index++
	}

	return intersect
}
