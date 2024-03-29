package linknode

import (
	"testing"
)

func TestMiddleNode(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 4, 5, 6})
	PrintNode(head)

	newHead := middleNode(head)
	PrintNode(newHead)
}

func middleNode(head *ListNode) *ListNode {
	//return middleNodeByArray(head)
	//return middleNodeBySinglePointer(head)
	return middleNodeByDoublePointer(head)
}

func middleNodeByArray(head *ListNode) *ListNode {
	array := make([]*ListNode, 0)
	for i := head; i != nil; i = i.Next {
		array = append(array, i)
	}

	return array[len(array)/2]
}

func middleNodeBySinglePointer(head *ListNode) *ListNode {
	var curr, num int

	for i := head; i != nil; i = i.Next {
		num++
	}
	for i := head; i != nil; i = i.Next {
		if curr == num/2 {
			return i
		}
		curr++
	}

	return nil
}

func middleNodeByDoublePointer(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
