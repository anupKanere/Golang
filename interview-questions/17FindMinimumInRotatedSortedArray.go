package main

func FindMin(nums []int) int {
	// numsCheck := []int{4, 5, 6, 7, 0, 1, 2}
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}
