package main

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		giveList *ListNode
		giveN    int
		wantList *ListNode
	}{
		{
			name:     "example 1",
			giveList: makeListNode(1, 2, 3, 4, 5),
			giveN:    2,
			wantList: makeListNode(1, 2, 3, 5),
		},
		{
			name:     "example 2",
			giveList: makeListNode(1),
			giveN:    1,
			wantList: makeListNode(),
		},
		{
			name:     "example 3",
			giveList: makeListNode(1, 2),
			giveN:    1,
			wantList: makeListNode(1),
		},
		{
			name:     "example 4",
			giveList: makeListNode(1, 2),
			giveN:    2,
			wantList: makeListNode(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList := removeNthFromEnd(tt.giveList, tt.giveN)
			assertListsEqual(t, tt.wantList, gotList)
		})
	}
}

func assertListsEqual(t *testing.T, want, got *ListNode) {
	if fmt.Sprint(want) != fmt.Sprint(got) {
		t.Fatalf(`Lists not equal, want "%v", got "%v"`, want, got)
	}
}
