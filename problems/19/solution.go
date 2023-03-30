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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	fast := dummy
	slow := dummy
	for i := 1; fast.Next != nil; i++ {
		fast = fast.Next
		if i <= n {
			continue
		}
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
