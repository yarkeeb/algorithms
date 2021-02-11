package linked_list

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func (l *ListNode) ToString() string {
	var s string
	iter := l
	for iter != nil {
		s += fmt.Sprint(iter.Val)
		iter = iter.Next
	}
	return s
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: 0,
		Next: nil,
	}
	curr := head
	for l1 != nil && l2 != nil {
		var next *ListNode
		if l1.Val < l2.Val {
			next = &ListNode{
				Val: l1.Val,
				Next: nil,
			}
			l1 = l1.Next
		} else {
			next = &ListNode{
				Val: l2.Val,
				Next: nil,
			}
			l2 = l2.Next
		}
		curr.Next = next
		curr = next
	}
	for l1 != nil {
		curr.Next = l1
		curr = curr.Next
		l1 = l1.Next
	}
	for l2 != nil {
		curr.Next = l2
		curr = curr.Next
		l2 = l2.Next
	}
	return head.Next
}
