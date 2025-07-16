package main

import (
	"sort"
	"strings"
)

func groupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		val := sortString(word)
		groups[val] = append(groups[val], word)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
