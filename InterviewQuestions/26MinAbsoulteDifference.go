package main

import "sort"

/*
Given an array of distinct integers arr, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair [a, b] follows

a, b are from arr
a < b
b - a equals to the minimum absolute difference of any two elements in arr
*/

func minAbsoluteDifference(arr []int) [][]int {
	if len(arr) < 2 {
		return [][]int{}
	}

	sort.Ints(arr)

	minDiff := arr[1] - arr[0]
	var res [][]int

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		// if diff found is less then reset the result else if diff == minDiff append it to the res
		if diff < minDiff {
			minDiff = diff
			res = [][]int{{arr[i], arr[i+1]}} 
		} else if diff == minDiff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}

