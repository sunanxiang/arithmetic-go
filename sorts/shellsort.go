package sorts

// 希尔排序：插入排序的泛化
func ShellSort(previous []int) []int {
	var (
		n    = len(previous)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := pow(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}
	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if previous[j-gap] > previous[j] {
					previous[j-gap], previous[j] = previous[j], previous[j-gap]
				}
				j = j - gap
			}
		}
	}

	return previous
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
