package linknode

import (
	"fmt"
	"testing"
)

func GenListNode(nums []int) *ListNode {
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	curr := head

	for i := 1; i < len(nums); i++ {
		next := &ListNode{Val: nums[i], Next: nil}
		curr.Next = next
		curr = next
	}

	return head
}

func PrintNode(head *ListNode) {
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Printf("%v ", temp.Val)
	}
	fmt.Println()
}

func TestGenListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	head := GenListNode(nums)
	PrintNode(head)
}
