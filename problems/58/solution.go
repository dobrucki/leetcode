package main

func lengthOfLastWord(s string) int {
	var last int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			last++
		}
		if i == 0 || s[i-1] == ' ' && s[i] != ' ' {
			return last
		}
	}
	return last
}
