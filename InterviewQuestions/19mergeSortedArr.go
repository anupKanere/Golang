package main

func mergeSortedArr(arr1, arr2 []int) []int {
	res := make([]int, 0, (len(arr1) + len(arr2)))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	// Append remaining elements

	if i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	if j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}

	return res
}
