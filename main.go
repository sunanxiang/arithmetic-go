package main

import (
	"fmt"

	"arithmetic-go/gosnippet"
	"arithmetic-go/sorts"
)

var testSlice = []int{4, 2, 3, 6, 5, 1, 7}

func main() {
	// 冒泡排序
	fmt.Println(sorts.BubbleSortOne(testSlice))
	fmt.Println(sorts.BubbleSortTwo(testSlice))

	// 选择排序
	fmt.Println(sorts.SelectSort(testSlice))

	// 插入排序
	fmt.Println(sorts.InsertSort(testSlice))

	// 切片追加、插入、删除操作
	fmt.Println(gosnippet.SliceHandle(testSlice, 8, 1))
}
