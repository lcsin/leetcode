package linknode

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := GenListNode([]int{1, 2, 2, 1})
	PrintNode(head)

	fmt.Println(isPalindrome(head))
}

func isPalindrome(head *ListNode) bool {
	//return isPalindromeByArray(head)
	//return isPalindromeByReverse(head)
	return isPalindromeByFastSlowPointer(head)
}

func isPalindromeByArray(head *ListNode) bool {
	array := make([]int, 0)
	for i := head; i != nil; i = i.Next {
		array = append(array, i.Val)
	}

	tail := len(array) - 1
	for i := 0; i < len(array)/2; i++ {
		if array[i] != array[tail] {
			return false
		} else {
			tail--
		}
	}

	return true
}

func isPalindromeByReverse(head *ListNode) bool {
	var newHead *ListNode
	curr := head
	b1 := strings.Builder{}

	// 反转链表并拼接反转前的值
	for curr != nil {
		b1.WriteString(fmt.Sprintf("%v", curr.Val))

		next := curr.Next
		curr.Next = newHead
		newHead = curr
		curr = next
	}

	// 遍历拼接反转后的值
	b2 := strings.Builder{}
	for i := newHead; i != nil; i = i.Next {
		b2.WriteString(fmt.Sprintf("%v", i.Val))
	}

	// 如果反转前和反转后的值一样，说明是回文链表
	if b1.String() == b2.String() {
		return true
	}
	return false
}

func isPalindromeByFastSlowPointer(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseMidNode := reverse(slow)

	curr := head
	for i := reverseMidNode; i != nil; i = i.Next {
		if curr.Val != i.Val {
			return false
		} else {
			curr = curr.Next
		}
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var newHead *ListNode
	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = newHead
		newHead = curr

		curr = next
	}

	return newHead
}
