package linknode

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	headA := GenListNode([]int{1})
	headB := GenListNode([]int{1})
	//headC := GenListNode([]int{2, 4})
	//headA.Next.Next = headC
	//headB.Next = headC
	PrintNode(headA)
	PrintNode(headB)

	node := getIntersectionNode(headA, headB)
	fmt.Println(node)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA.Next == nil && headB.Next == nil {
		return nil
	}

	for i := headA; i != nil; i = i.Next {
		for j := headB; j != nil; j = j.Next {
			if i.Next == j.Next {
				return i.Next
			}
		}
	}

	return nil
}
