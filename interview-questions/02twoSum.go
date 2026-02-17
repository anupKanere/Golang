package main

func twoSum(arr []int, target int) []int {
	seen := make(map[int]int) // map [value] = index

	for i, num := range arr {
		compliment := target - num
		val, found := seen[compliment]
		if found {
			return []int{val, i}
		}
		seen[num] = i
	}
	return nil
}
