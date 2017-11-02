package gosnippet

func UniqueSliceString(slice []string) []string {
	found := make(map[string]bool)
	total := 0
	for i, val := range slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			(slice)[total] = (slice)[i]
			total++
		}
	}
	slice = (slice)[:total]
	return slice
}
