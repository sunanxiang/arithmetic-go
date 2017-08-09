package sorts

// 冒泡排序(每趟都选出一个相对较大值放在最后)，升序排序
// 方法1：利用isSorted 可以提前退出循环
func BubbleSortOne(previous []int) []int {
	isSorted := false
	length := len(previous)

	for {
		if !isSorted {
			isSorted = true
			for i := 0; i < length-1; i++ {
				if previous[i] > previous[i+1] {
					previous[i], previous[i+1] = previous[i+1], previous[i]
					isSorted = false
				}
			}
			continue
		}
		length--
		break
	}

	return previous
}

// 方法2：无脑循环排序
func BubbleSortTwo(previous []int) []int {
	for i := 0; i < len(previous)-1; i++ {
		for j := i + 1; j < len(previous); j++ {
			if previous[i] > previous[j] {
				previous[i], previous[j] = previous[j], previous[i]
			}
		}
	}

	return previous
}
