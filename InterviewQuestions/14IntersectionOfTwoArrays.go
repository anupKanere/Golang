package main

func intersectionOfTwoArrays(nums1, nums2 []int) []int {

	nums1Map := make(map[int]bool)

	for _, val := range nums1 {
		nums1Map[val] = true
	}

	seen := make(map[int]bool)
	var result []int
	for _, val := range nums2 {
		if nums1Map[val] {
			if !seen[val] {
				result = append(result, val)
			}
			seen[val] = true
		}
	}
	return result
}
