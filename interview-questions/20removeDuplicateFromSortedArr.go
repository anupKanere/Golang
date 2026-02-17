package main

import "fmt"

func removeDuplicate(arr []int) []int {
	idx := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[idx] {
			idx++
			arr[idx] = arr[i]
		}
	}
	fmt.Println(idx)
	fmt.Println(arr[:idx+1])
	return arr[:idx+1]
}
