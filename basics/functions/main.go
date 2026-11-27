package main

import "fmt"

// main function
func main() {

	defer fmt.Println("Calling defer function - 1")
	defer fmt.Println("Calling defer function - 2")

	// defer greet()
	greet()

	a, b := 10, 20
	addition := Add(a, b)
	fmt.Println("RESULT = ", addition)

	x, y := 10, 0

	res, err := Division(x, y)
	if err != nil {
		fmt.Println("Error Occured : ", err)
	} else {
		fmt.Printf("Result = %d \n ", res)
	}

	// Anonymous function
	minus := func(a, b int) int {
		return a - b
	}

	fmt.Println(minus(5, 2))

	// Recursion
	fact := Factorial(5)
	fmt.Println("Factorial - ", fact)

	// Closure
	double := Multiplier(5)
	fmt.Println("Closure example ", double(2))

	counter1 := Counter(5)

	fmt.Println(counter1(2))
	fmt.Println(counter1(2))
	fmt.Println(counter1(2))
	fmt.Println(counter1(2))
	fmt.Println(counter1(1))
	fmt.Println(counter1(1))
	fmt.Println(counter1(1))

}
