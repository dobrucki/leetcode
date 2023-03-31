package main

import "testing"
import "fmt"

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		desc      string
		giveList1 *ListNode
		giveList2 *ListNode
		wantList  *ListNode
	}{
		{
			desc:      "example 1",
			giveList1: makeListNode(1, 2, 4),
			giveList2: makeListNode(1, 3, 4),
			wantList:  makeListNode(1, 1, 2, 3, 4, 4),
		},
		{
			desc:      "example 2",
			giveList1: makeListNode(),
			giveList2: makeListNode(),
			wantList:  makeListNode(),
		},
		{
			desc:      "example 3",
			giveList1: makeListNode(),
			giveList2: makeListNode(0),
			wantList:  makeListNode(0),
		},
		{
			desc:      "example 4",
			giveList1: makeListNode(1),
			giveList2: makeListNode(),
			wantList:  makeListNode(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotList := mergeTwoLists(tt.giveList1, tt.giveList2)
			assertListsEqual(t, tt.wantList, gotList)
		})
	}
}

func assertListsEqual(t *testing.T, want, got *ListNode) {
	if fmt.Sprint(want) != fmt.Sprint(got) {
		t.Fatalf(`Lists not equal, want "%v", got "%v"`, want, got)
	}
}
