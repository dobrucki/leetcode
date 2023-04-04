package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		desc string
		give string
		want int
	}{
		{
			desc: "example 1",
			give: "",
			want: 0,
		},
		{
			desc: "example 2",
			give: "Hello World",
			want: 5,
		},
		{
			desc: "example 3",
			give: "   fly me   to   the moon  ",
			want: 4,
		},
		{
			desc: "example 4",
			give: "luffy is still joyboy",
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := lengthOfLastWord(tt.give)
			if tt.want != got {
				t.Fatalf(`wanted "%v", got "%v"`, tt.want, got)
			}
		})
	}
}
