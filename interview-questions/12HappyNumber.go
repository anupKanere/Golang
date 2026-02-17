package main

func isHappyNumber(num int) bool {

	seen := make(map[int]bool)

	for num != 1 {
		if seen[num] {
			return false
		}
		seen[num] = true
		num = findSumOfSquare(num)
	}

	return true
}

func findSumOfSquare(n int) int {

	sum := 0
	for n > 0 {
		digit := n % 10
		sq := digit * digit
		sum = sum + sq
		n = n / 10
	}

	return sum
}
