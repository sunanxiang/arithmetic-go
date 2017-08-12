package sorts

// 插入排序：从第二个开始，每次都和前面的比较排序，总保证前面的是有序的
func InsertSort(previous []int) []int {
	for i := 1; i < len(previous); i++ {
		for j := i; j > 0 && previous[j] < previous[j-1]; j-- {
			previous[j], previous[j-1] = previous[j-1], previous[j]
		}
	}

	return previous
}
