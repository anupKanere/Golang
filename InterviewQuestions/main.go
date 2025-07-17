package main

import (
	"fmt"
)

func main() {
	// LinkedList
	values1 := []int{1, 3, 5, 7, 9}
	values2 := []int{2, 4, 6, 8, 10}
	head1 := createLinkedList(values1)
	head2 := createLinkedList(values2)
	display(head1)
	display(head2)
	// reverseList := reverse(head1)
	// display(reverseList)
	mergedList := mergeSortedList(head1, head2)
	display(mergedList)

	fmt.Println("------------------------------------------------------------------")

	// 02 Two sun
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	target := 9
	res := twoSum(nums, target)
	fmt.Printf("Indexes for result %d is %v \n", target, res)

	fmt.Println("------------------------------------------------------------------")

	// 03 contains Duplicate
	fmt.Printf("Array has duplicate value = %t \n", containsDuplicate(nums))

	fmt.Println("------------------------------------------------------------------")

	// 04 product of array except self
	arr := []int{1, 2, 3, 4}
	fmt.Printf("Product of Array Excep Self is \n %v \n", productExceptSelf(arr))

	fmt.Println("------------------------------------------------------------------")
	// 05 Max Sub array
	maxSubarr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Max sub array = ", maxSubArr(maxSubarr))

	fmt.Println("------------------------------------------------------------------")
	// 06 Move Zeroes
	zerosArr := []int{1, 2, 0, 5, 0, 7, 6, 0}
	fmt.Println("Move zeros = ", moveZeros(zerosArr))

	fmt.Println("------------------------------------------------------------------")
	// 07 Rotate Array
	rotateArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 3
	fmt.Println("Rotated Array := ", rotateArray(rotateArr, k))

	fmt.Println("------------------------------------------------------------------")
	// 08 Longest Substring Without Repeating Characters
	str := "abcabcdefabcde"
	fmt.Println("Length of longest Substring : ", lengthOfLongestSubstring(str))
	fmt.Println("Longest Substring : ", LongestSubstring(str))

	fmt.Println("------------------------------------------------------------------")
	// 09 Valid Anagram
	fmt.Println("Is valid Anagram : ", validAnagram("listen", "iselnt"))

	fmt.Println("------------------------------------------------------------------")
	//10 Group Anagrams
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Group Anagrams : ", groupAnagrams(words))

	fmt.Println("------------------------------------------------------------------")
	// 11 TopKFrequentElements
	topNums := []int{1, 1, 1, 2, 2, 2, 2, 3, 4, 9, 6, 5, 5, 5, 5, 5, 5, 5}
	fmt.Println("Top k elements : ", TopKFrequentElements(topNums, 3))

	fmt.Println("------------------------------------------------------------------")
	// 12 Happy Number
	fmt.Println("Is Happy number : ", isHappyNumber(10))

	fmt.Println("------------------------------------------------------------------")
	// 13 Isomorphic Strings
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
	fmt.Println(isIsomorphic("ab", "aa"))       // false

	fmt.Println("------------------------------------------------------------------")
	// 14 Intersection of Two Arrays
	fmt.Println(intersectionOfTwoArrays([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println("------------------------------------------------------------------")

	// 15 Basic Binary Seach
	fmt.Println("------------------------------------------------------------------")
	targetSearch := 9
	searchArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Basic Binary Search : \n Number ", targetSearch, "is available at index ", basicBinarySearch(searchArr, targetSearch))
	fmt.Println("------------------------------------------------------------------")

	// 16 Search in Rotated Sorted Array
	numsCheck := []int{4, 5, 6, 7, 0, 1, 2}
	targetS := 10
	fmt.Println("Target is Available at index : ", seachInRotatedArray(numsCheck, targetS))

	// 17 Find Minimum in Rotated Sorted Array
	fmt.Println("Min in sorted arr : ", FindMin(numsCheck))

	// 18 Valid Palindrome
	palindomeStr := "ABCDDCBA"
	fmt.Println("Valid Palindrome : ", ispalindrome(palindomeStr))

	// 19 Merge Sorted Array
	num1 := []int{2, 4, 6, 8}
	num2 := []int{1, 3, 5, 7, 9}
	fmt.Println("Merge Sorted Array : ", mergeSortedArr(num1, num2))
}
