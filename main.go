package main

import (
	"fmt"

	"arithmetic-go/sorts"
)

var testSlice = []int{4, 2, 3, 6, 5, 1, 7}

func main() {
	// 冒泡排序
	fmt.Println(sorts.BubbleSortOne(testSlice))
	fmt.Println(sorts.BubbleSortTwo(testSlice))
}
