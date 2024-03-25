package linknode

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	node1 := &ListNode{
		Val:  3,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node1.Next = node2
	node2.Next = node2
	node3.Next = node4
	node4.Next = node2

	//node1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}

	fmt.Println(hasCycle(node1))
	fmt.Println(hasCycleByDoublePointer(node1))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	m := make(map[*ListNode]bool)
	for node := head; node != nil; node = node.Next {
		if _, ok := m[node]; ok {
			return true
		} else {
			m[node] = true
		}
	}

	return false
}

func hasCycleByDoublePointer(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
