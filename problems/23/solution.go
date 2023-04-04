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
			Val: v}
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
		const delimiter = "->"
		fmt.Fprint(&sb, delimiter)
		node = node.Next
	}
	sb.WriteString("]")
	return sb.String()
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	next := dummy
	for {
		idx := -1
		for i, node := range lists {
			if node == nil {
				continue
			}
			if idx < 0 {
				idx = i
				continue
			}
			if node.Val < lists[idx].Val {
				idx = i
			}
		}
		if idx < 0 {
			break
		}
		next.Next = lists[idx]
		lists[idx] = lists[idx].Next
		next = next.Next
	}
	return dummy.Next
}
