package main

import "sort"

func TopKFrequentElements(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, val := range nums {
		counts[val]++
	}

	type kv struct {
		key int
		val int
	}

	var freqList []kv
	for key, val := range counts {
		freqList = append(freqList, kv{key, val})
	}

	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].val > freqList[j].val
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = freqList[i].key
	}

	return res
}
