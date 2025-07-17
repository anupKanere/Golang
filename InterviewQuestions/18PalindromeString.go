package main

import "strings"

func ispalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	low, high := 0, len(s)-1
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
