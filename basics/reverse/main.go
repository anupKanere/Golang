package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)
	// words := strings.Split(input, " ") // this is also valid
	fmt.Println("words - ", words)
	fmt.Println("Length of words = ", len(words))

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	words := "Pune Chennai Bengalore"
	rev := reverseWords(words)
	fmt.Println("Reversed words - ", rev)
}
