package general

import "math/rand"

// 检测是否有序
func IsSorted(previous []int) bool {
	for i := 1; i < len(previous); i++ {
		if previous[i] < previous[i-1] {
			return false
		}
	}

	return true
}

// 随机生成排序测试数据
func RandTestSlice() []int {
	var result []int

	for i := 0; i < 100; i++ {
		result = append(result, rand.Intn(100))
	}

	return result
}
