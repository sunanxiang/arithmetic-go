package sorts

// 检测是否有序
func IsSorted(previous []int) bool {
	for i := 1; i < len(previous); i++ {
		if previous[i] < previous[i - 1] {
			return false
		}
	}

	return true
}
