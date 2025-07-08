package main

import "fmt"

// Function with zero parameter and without return
func greet() {
	fmt.Println("Hello Everyone!!!")
}

// Capitalized name makes it exported (optional here, since same package)
// Function with multiple parameters and single return type
func Add(num1 int, num2 int) int {
	fmt.Printf("Num 1 = %d Num 2 = %d \n", num1, num2)
	return num1 + num2
}

// Function returning multiple values
func Division(a, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return a / b, nil
}

// Recursion

func Factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * Factorial(num-1)
}

// Closure function
func Multiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
