package list

// singly linked list
type SinglyNode struct {
	Next  *SinglyNode
	Value int
}

// two-way linked list
type TwoWayNode struct {
	Next  *TwoWayNode
	Pre   *TwoWayNode
	Value int
}

// reverse singly linked list use non-recursive way
func ReverseSinglyLinkedListNonRecursive(root *SinglyNode) *SinglyNode {
	if root == nil || root.Next == nil {
		return root
	}

	var preNode, nextNode *SinglyNode

	for root.Next != nil {
		nextNode = root.Next
		root.Next = preNode
		preNode = root
		root = nextNode
	}
	return preNode
}

// reverse singly linked list use recursive way
func ReverseSinglyLinkedListRecursive(root *SinglyNode) *SinglyNode {
	if root == nil || root.Next == nil {
		return root
	}

	var nextNode *SinglyNode
	nextNode = root.Next
	root.Next = nil
	reverse := ReverseSinglyLinkedListRecursive(nextNode)
	nextNode.Next = root
	return reverse
}

// reverse singly linked list use easy way that in go.
func ReverseSinglyLinkedListSimple(root *SinglyNode) *SinglyNode {
	var tail *SinglyNode
	for root != nil {
		root.Next, tail, root = tail, root, root.Next
	}
	return tail
}

// todo:
// reverse two-way linked list
func ReverseTwoWayLinkedList(root *TwoWayNode) *TwoWayNode {
	return nil
}
