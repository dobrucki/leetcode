package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var _ fmt.Stringer = (*ListNode)(nil)

func makeListNode(values ...int) *ListNode {
	dummy := &ListNode{}
	next := dummy
	for _, v := range values {
		next.Next = &ListNode{
			Val: v,
		}
		next = next.Next
	}
	return dummy.Next
}

func (ln ListNode) String() string {
	var sb strings.Builder
	node := &ln
	sb.WriteString("[")
	for {
		fmt.Fprint(&sb, node.Val)
		if node.Next == nil {
			break
		}
		const delimiter = ", "
		fmt.Fprint(&sb, delimiter)
		node = node.Next
	}
	sb.WriteString("]")
	return sb.String()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	next := dummy
	for {
		if list1 == nil {
			next.Next = list2
			break
		}
		if list2 == nil {
			next.Next = list1
			break
		}
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next  =list2
			list2 = list2.Next
		}
		next = next.Next
	}
	return dummy.Next
}
