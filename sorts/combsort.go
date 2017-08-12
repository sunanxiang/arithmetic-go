package sorts

// 梳排序：冒泡排序的改进，间距更大，shrink 常量通常定为1.3
func CombSort(previous []int) []int {
	var (
		gap     = len(previous)
		shrink  = 1.3
		swapped = true
	)

	for swapped || gap > 1 {
		if gap > 1 {
			gap = int(float64(gap) / shrink)
		}
		swapped = false

		for i := 0; i + gap < len(previous); i++ {
			if previous[i] > previous[i+gap] {
				previous[i+gap], previous[i] = previous[i], previous[i+gap]
				swapped = true
			}
		}
	}

	return previous
}
