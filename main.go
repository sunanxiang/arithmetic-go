package main

import (
	"fmt"

	"github.com/sunanxiang/arithmetic-go/general"
	"github.com/sunanxiang/arithmetic-go/gosnippet"
	"github.com/sunanxiang/arithmetic-go/sorts"
)

func main() {
	// 冒泡排序
	testBubbleOne := general.RandTestSlice()
	testBubbleTwo := general.RandTestSlice()
	fmt.Println(sorts.BubbleSortOne(testBubbleOne))
	fmt.Println(sorts.BubbleSortTwo(testBubbleTwo))

	// 选择排序
	testSelect := general.RandTestSlice()
	fmt.Println(sorts.SelectSort(testSelect))

	// 插入排序
	testInsert := general.RandTestSlice()
	fmt.Println(sorts.InsertSort(testInsert))

	// 希尔排序
	testShell := general.RandTestSlice()
	fmt.Println(sorts.ShellSort(testShell))

	// 梳排序
	testComb := general.RandTestSlice()
	fmt.Println(sorts.CombSort(testComb))

	// 归并排序
	testMerge := general.RandTestSlice()
	fmt.Println(sorts.MergeSort(testMerge))

	// 快速排序
	testQuick := general.RandTestSlice()
	fmt.Println(sorts.QuickSort(testQuick, 0, len(testQuick)-1))

	// 切片追加、插入、删除操作
	fmt.Println(gosnippet.SliceHandle(testQuick, 8, 1))
}
