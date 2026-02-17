package main

func rotateArray(nums []int, k int) []int {
	k = k % len(nums)
	n := len(nums)

	reverseArr(nums, 0, n-1)
	reverseArr(nums, 0, k-1)
	reverseArr(nums, k, n-1)

	return nums
}

func reverseArr(arr []int, start, end int) []int {

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
