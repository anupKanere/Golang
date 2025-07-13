package main

import (
	"fmt"
	"sync"
	"time"
)

func increment2(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Go routine Thread!!!")
	for i := 1; i <= 10; i++ {
		num += 1
	}
	fmt.Println("Final Num with Waitgroup: ", num)
}

func increment(num int) {
	fmt.Println("Go routine Thread!!!")
	for i := 1; i <= 10; i++ {
		num += 1
	}
	fmt.Println("Final Num : ", num)
}

func main() {
	x := 10

	// First method
	go increment(x)
	time.Sleep(1 * time.Second)
	fmt.Println("Main Thread")

	// Second method using sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go increment2(x, &wg)
	wg.Wait()
	fmt.Println("Main thread Complete!!!")

}
