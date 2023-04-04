package main

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		desc      string
		giveLists []*ListNode
		wantList  *ListNode
	}{
		{
			desc: "example 1",
			giveLists: []*ListNode{
				makeListNode(1, 4, 5),
				makeListNode(1, 3, 4),
				makeListNode(2, 6),
			},
			wantList: makeListNode(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			desc: "example 2",
			giveLists: []*ListNode{},
			wantList: makeListNode(),
		},
		{
			desc: "example 3",
			giveLists: []*ListNode{
				makeListNode(),
			},
			wantList: makeListNode(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotList := mergeKLists(tt.giveLists)
			assertListsEqual(t, tt.wantList, gotList)
		})
	}
}

func assertListsEqual(t *testing.T, want, got *ListNode) {
	if fmt.Sprint(want) != fmt.Sprint(got) {
		t.Fatalf(`Lists not equal, want "%v", got "%v"`, want, got)
	}
}
