package main

func validParentheses(s string) bool {

	stack := []rune{}

	charMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		stackLen := len(stack)
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if stackLen == 0 || charMap[char] != stack[stackLen-1] {
				return false
			}
			stack = stack[:stackLen-1]
		}
	}

	return len(stack) == 0
}
