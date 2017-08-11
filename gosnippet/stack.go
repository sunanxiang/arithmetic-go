package gosnippet

import (
	"fmt"
	"os"
)

type Node struct {
	val   interface{}
	pNode *Node
}

type Stack struct {
	pTop, pBottom *Node
}

func (pStack *Stack) InitStack() {
	pNew := new(Node)
	pNew.pNode = nil
	pStack.pTop = pNew
	pStack.pBottom = pNew
	if pStack.pTop == nil || pStack.pBottom == nil {
		fmt.Println("分配头节点内存失败，程序退出")
		os.Exit(-1)
	}
}

func (pStack *Stack) Push(val interface{}) {
	pNew := new(Node)
	if pNew == nil {
		fmt.Println("分配头节点内存失败，程序退出")
		os.Exit(-1)
	}
	pNew.val = val

	pNew.pNode = pStack.pTop
	pStack.pTop = pNew
}

func (pStack *Stack) Pop() interface{} {
	if pStack.Empty() {
		return nil
	}

	reval := pStack.pTop.val
	pStack.pTop = pStack.pTop.pNode

	return reval
}

func (pStack *Stack) Empty() bool {
	return pStack.pTop == pStack.pBottom
}

func (pStack *Stack) Top() interface{} {
	if pStack.Empty() {
		return nil
	}

	return pStack.pTop.val
}

func (pStack *Stack) Size() int {
	var count = 0
	if pStack.Empty() {
		return 0
	}

	for p := pStack.pTop; p != pStack.pBottom; p = p.pNode {
		count++
	}

	return count
}

func (pStack *Stack) Clear() {
	if pStack.Empty() {
		return
	}
	pStack.pTop = pStack.pBottom
}
