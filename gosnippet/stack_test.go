package gosnippet

import (
	"fmt"
	"testing"
)

func TestStack_InitStack(t *testing.T) {
	var stack Stack
	stack.InitStack()

	fmt.Println(stack.Empty())

	stack.Push(3)
	fmt.Println(stack.Top())
	stack.Push(5)
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())

	fmt.Println(stack.Top())
	fmt.Println(stack.Size())
}
