package sorts

// 选择排序：每次都选出最小放在前面
func SelectSort(previous []int) []int {
	// 每次都从上次选出的最小值后面开始
	for i := 0; i < len(previous); i++ {
		min := previous[i]

		// 选出该次循环中最小的值，放入最前面
		for j := i + 1; j < len(previous); j++ {
			if previous[j] < min {
				min = previous[j]
				previous[i], previous[j] = previous[j], previous[i]
			} else {
				previous[i], min = min, previous[i]
			}
		}
	}

	return previous
}
