package main

func validAnagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	charMap := make(map[rune]int)

	for _, char := range s1 {
		charMap[char]++
	}

	for _, char := range s2 {
		charMap[char]--
		if charMap[char] < 0 {
			return false
		}
	}

	for _, val := range charMap {
		if val != 0 {
			return false
		}
	}

	return true

}
