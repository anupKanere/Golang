package main

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left := 0
	charMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if idx, found := charMap[s[right]]; found && idx >= left {
			left = idx + 1
		}

		charMap[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func LongestSubstring(s string) string {
	maxLen := 0
	left := 0
	maxStart := 0
	charMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if idx, found := charMap[s[right]]; found && idx >= left {
			left = right + 1
		}

		charMap[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
			maxStart = left
		}
		// maxLen = max(maxLen, right-left+1)
	}
	return s[maxStart : maxStart+maxLen]
}
