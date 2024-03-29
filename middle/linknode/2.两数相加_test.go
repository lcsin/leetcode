package linknode

import (
	"testing"
)

func Test(t *testing.T) {
	l1 := GenListNode([]int{2, 4, 3})
	l2 := GenListNode([]int{5, 6, 4})
	PrintNode(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 int
	var N int
	for i := l1; i != nil; i = i.Next {
		num1 += i.Val * getN1(N)
		N++
	}

	N = 0
	for i := l2; i != nil; i = i.Next {
		num2 += i.Val * getN1(N)
		N++
	}

	result := num1 + num2
	head := &ListNode{
		Val: result % 10,
	}
	result = result / 10
	curr := head
	for result != 0 {
		next := &ListNode{Val: result % 10}
		curr.Next = next
		curr = next
		result = result / 10
	}

	return head
}

func getN1(n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num = num * 10
	}
	return num
}

func reverse(head *ListNode) *ListNode {
	var newHead, curr *ListNode = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = next
	}
	return newHead
}
