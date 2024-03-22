package easy

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveElements(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	printNode(head)

	removeElements(head, 6)
	printNode(head)
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	for temp := head; temp.Next != nil; {
		if temp.Next.Val == val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	if head.Val == val {
		return head.Next
	}
	return head
}

func printNode(head *ListNode) {
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Printf("%v ", temp.Val)
	}
	fmt.Println()
}
