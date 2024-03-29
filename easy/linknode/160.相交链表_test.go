package linknode

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	headA := GenListNode([]int{4, 1})
	headB := GenListNode([]int{5, 6, 1})
	headC := GenListNode([]int{8, 4, 5})
	headA.Next.Next = headC
	headB.Next.Next.Next = headC
	PrintNode(headA)
	PrintNode(headB)

	node := getIntersectionNode(headA, headB)
	fmt.Println(node)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//return getIntersectionNodeByIteration(headA, headB)
	//return getIntersectionNodeByMap(headA, headB)
	return getIntersectionNodeByDoublePointer(headA, headB)
}

func getIntersectionNodeByIteration(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	for i := headA; i != nil; i = i.Next {
		for j := headB; j != nil; j = j.Next {
			if i == j {
				return i
			}
		}
	}

	return nil
}

func getIntersectionNodeByMap(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})

	for i := headA; i != nil; i = i.Next {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
		}
	}
	for i := headB; i != nil; i = i.Next {
		if _, ok := m[i]; ok {
			return i
		}
	}

	return nil
}

func getIntersectionNodeByDoublePointer(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for {
		if pa == pb {
			return pa
		}
		if pa == nil && pb == nil {
			return nil
		}

		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
}
