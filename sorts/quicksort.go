package sorts

// 快速排序
func QuickSort(previous []int, start, end int) []int {
	if start >= end {
		return nil
	}
	// 以最后一个数据为基准点
	s := end
	// i为从头开始遍历时小于基准的游标
	i := start - 1
	// j为从头开始遍历时大于基准的游标
	j := start
	for j <= s {
		//添加=的意思是将最后的基准点调整到小于
		//和大于它的中间
		if previous[j] <= previous[s] {
			previous[i+1], previous[j] = previous[j], previous[i+1]
			i++
		}
		j++
	}
	QuickSort(previous, start, i-1)
	QuickSort(previous, i+1, end)

	return previous
}
