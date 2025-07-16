package main

func containsDuplicate(arr []int) bool {
	seen := make(map[int]bool)
	for _, v := range arr {
		if _, found := seen[v]; found {
			return true
		}
		seen[v] = true
	}
	return false
}
