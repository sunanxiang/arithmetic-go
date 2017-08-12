package sorts

// 归并排序：划分子问题，分治算法
func MergeSort(previous []int) []int {
	var n = len(previous)

	if n == 1 {
		return previous
	}

	middle := int(n / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, n-middle)
	)
	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = previous[i]
		} else {
			right[i-middle] = previous[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
