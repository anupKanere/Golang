package main

import (
	"fmt"
	"math"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func ispalindrome(input string) bool {

	s := strings.ToLower(strings.ReplaceAll(input, " ", ""))

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func longestSubString(input string) string {
	charMap := make(map[rune]int)
	maxStart := 0
	maxLen := 0
	start := 0

	for i, char := range input {
		if lastIndex, exists := charMap[char]; exists && lastIndex >= start {
			start = lastIndex + 1
		}

		charMap[char] = i
		currentLength := i - start + 1

		if currentLength > maxLen {
			maxLen = currentLength
			maxStart = start
		}
	}

	return input[maxStart : maxStart+maxLen]
}

func test(s string) string {
	charMap := make(map[rune]int)
	maxLen := 0
	maxStart := 0
	start := 0

	for i, char := range s {
		if lastIndex, exists := charMap[char]; exists && lastIndex >= start {
			start = lastIndex + 1
		}

		charMap[char] = i
		currLen := i - start + 1

		if currLen > maxLen {
			maxLen = currLen
			maxStart = start
		}
	}
	return s[maxStart : maxStart+maxLen]
}

func findSecondLargest(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	first := math.MinInt
	second := math.MinInt

	for _, num := range nums {
		if num > first {
			second = first
			first = num
		} else if num > second && num != first {
			second = num
		}
	}
	return second

}

func lenLongestSubString(input string) int {
	charMap := make(map[rune]int)
	maxLen := 0
	start := 0
	// maxStart := 0

	for i, char := range input {
		if lastIndex, exists := charMap[char]; exists && lastIndex >= start {
			start = lastIndex + 1
		}
		charMap[char] = i
		currLen := i - start + 1

		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}

func longestIncreasingSubsequence(arr []int) int {
	n := len(arr)

	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := dp[0]

	for i := 1; i < n; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max

}

func findMissingNum(arr []int, givenLen int) int {

	expactedSum := givenLen * (givenLen + 1) / 2

	actualSum := 0

	for _, num := range arr {
		actualSum += num
	}

	return expactedSum - actualSum
}

func main() {
	// 1. Write a Go program to reverse a string
	input := "Anup Kanere"
	res := reverseString(input)

	fmt.Println("Result = ", res)
	// 2. String is a palindrome

	check := []string{"test", "hello", "AnupunA", "racecar", "121"}
	for _, val := range check {
		ispal := ispalindrome(val)
		fmt.Printf("'%s' is palindrome: %v\n", val, ispal)

	}
	// 3. Check disk usage of root filesystem and send email alert if usage exceeds threshold
	// 4. Function in Go to return a singly linked list
	// 5. Find the longest substring without repeating characters
	substringCheck := "abcabcdab"
	maxSubString := longestSubString(substringCheck)
	fmt.Println("Longest substring = ", maxSubString)
	// 6. Program to print second largest number
	numList := []int{1, 2, 3, 4, 25, 5, 6, 7, 8, 9}
	secondLargest := findSecondLargest(numList)
	fmt.Println("SecondLargest = ", secondLargest)
	// 7. Length of longest substring without repeating characters
	substringTest := "abcabcdab"
	lenSubString := lenLongestSubString(substringTest)
	fmt.Println("Length Longest substring = ", lenSubString)
	// 8. Merge two sorted linked lists into a single sorted list
	// 9. Largest increasing subsequence in an array
	arr := []int{10, 22, 9, 33, 21, 50, 41, 60}
	fmt.Println("Length of Longest Increasing Subsequence is:", longestIncreasingSubsequence(arr))

	// 10. Missing number in array
	incompleteArr := []int{1, 2, 4, 5, 6}
	totalLen := 6
	missNum := findMissingNum(incompleteArr, totalLen)
	fmt.Println("Missing num = ", missNum)

}
