package gosnippet

import "fmt"

// Go 语言切片的追加、删除、插入
func SliceHandle(previous []int, value int, index int) []int {
	fmt.Println("原切片为：", previous)

	// 追加
	previous = append(previous, value)

	fmt.Println("追加元素：", value)
	fmt.Println("追加后的切片为", previous)

	// 插入
	temp := append([]int{}, previous[index:]...)
	previous = append(previous[:index], value)
	previous = append(previous, temp...)

	fmt.Println("在", index, "位置插入元素：", value)
	fmt.Println("插入后的切片为:", previous)

	// 删除
	previous = append(previous[:index], previous[index+1:]...)

	fmt.Println("删除", index, "位置上的元素:", value)

	return previous
}
