/*
Given an integer array nums, return an array answer such that
answer[i] is equal to the product of all the elements of nums except nums[i],
without using division, and in O(n) time.

Input:
nums = [1, 2, 3, 4]
Output:
[24, 12, 8, 6]

Explanation:

Index 0: 2×3×4 = 24

Index 1: 1×3×4 = 12

Index 2: 1×2×4 = 8

Index 3: 1×2×3 = 6
*/
package main

func productExceptSelf(arr []int) []int {
	n := len(arr)
	res := make([]int, n)

	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * arr[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = rightProduct * res[i]
		rightProduct = rightProduct * arr[i]
	}

	return res
}
