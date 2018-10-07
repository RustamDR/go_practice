package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root, sum, residue := &ListNode{0, nil}, 0, 0
	current := root

	for {
		sum = current.Val
		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		current.Val = sum % 10
		residue = (sum - current.Val) / 10
		current.Next = &ListNode{residue, nil}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			if residue == 0 {
				current.Next = nil
			}
			break
		}

		current = current.Next
	}

	return root
}

func main() {
	l1, l2 := &ListNode{1, nil},
		&ListNode{9, &ListNode{9, nil}}
	//l1, l2 := &ListNode{1, &ListNode{8, nil}},
	//	&ListNode{0, nil}

	res := addTwoNumbers(l1, l2)
	for next := res; next != nil; next = next.Next {
		fmt.Print(next.Val)
	}
}
