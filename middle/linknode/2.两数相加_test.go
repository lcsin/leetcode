package linknode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//l1 := GenListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	//l2 := GenListNode([]int{5, 6, 4})
	l1 := GenListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := GenListNode([]int{9, 9, 9, 9})
	PrintNode(addTwoNumbersByArray(l1, l2))
	PrintNode(addTwoNumbersByLinkNode(l1, l2))
}

func addTwoNumbersByLinkNode(l1 *ListNode, l2 *ListNode) *ListNode {
	num := l1.Val + l2.Val
	head := &ListNode{}

	var count int
	if num >= 10 {
		head.Val = num % 10
		count = num / 10
	} else {
		head.Val = num
	}
	curr := head

	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil || l2 != nil {
		num = 0
		var next *ListNode
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		num += count

		if num < 10 {
			count = 0
			next = &ListNode{Val: num}
		} else {
			next = &ListNode{Val: num % 10}
			count = num / 10
		}
		curr.Next = next
		curr = next
	}

	if count > 0 {
		curr.Next = &ListNode{Val: count}
	}

	return head
}

func addTwoNumbersByArray(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1, arr2, arr3 := make([]int, 0), make([]int, 0), make([]int, 0)

	for i := l1; i != nil; i = i.Next {
		arr1 = append(arr1, i.Val)
	}
	for i := l2; i != nil; i = i.Next {
		arr2 = append(arr2, i.Val)
	}

	var count, idx int
	for idx < len(arr1) || idx < len(arr2) {
		var num int
		if idx < len(arr1) {
			num += arr1[idx]
		}
		if idx < len(arr2) {
			num += arr2[idx]
		}
		num += count

		if num < 10 {
			arr3 = append(arr3, num)
			count = 0
		} else {
			count = num / 10
			arr3 = append(arr3, num%10)
		}

		idx++
	}
	if count > 0 {
		arr3 = append(arr3, count)
	}

	fmt.Println(arr3)
	return array2linkList(arr3)
}

func array2linkList(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}

	head := &ListNode{Val: array[0]}
	curr := head
	for i := 1; i < len(array); i++ {
		next := &ListNode{Val: array[i]}
		curr.Next = next
		curr = next
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
