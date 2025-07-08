package main

import "fmt"

func main() {
	fmt.Println("For loop in Go : ")

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("Second Way : ")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Third Way : ")

	// Infinite loop
	// j := 0
	// for {
	// 	fmt.Println("Printing Infinite Loop j : " , j)
	// 	j++
	// }

	// Looping over array
	nums := []int{0, 1, 2, 3, 4, 5}

	for key, val := range nums {
		fmt.Println("Key - ", key, "Value - ", val*2)
	}

	for _, val := range nums {
		fmt.Println("Value - ", val)
	}

	// looping over map
	userMap := map[string]string{
		"name": "anup",
		"city": "pune",
		"age":  "28",
	}

	for key, val := range userMap {
		fmt.Println("key : ", key, "val : ", val)
		if userMap[key] == "anup" {
			fmt.Println("name Found!!!")
		} else if val == "28" {
			fmt.Println("Age found")
		}

	}

	// loop over channel
	fmt.Println("LOOP OVER CHANNEL!!!")
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	// loop over string

	name := "Anup"

	for i, char := range name {
		// fmt.Println("index - ", i, "character - ", char)
		fmt.Printf("Char at index %d is %c \n", i, char)
	}
}
