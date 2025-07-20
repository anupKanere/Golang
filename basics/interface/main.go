package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

type Human struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof"
}

func (h Human) Speak() string {
	return "my Name is " + h.Name
}

func saySomething(s Speaker) {
	res := s.Speak()
	fmt.Println(res)
}

func main() {
	dog := Dog{}
	human := Human{Name: "Anup"}

	//Method 1
	var s1 Speaker
	var s2 Speaker
	s1 = dog
	s2 = human
	res1 := s1.Speak()
	res2 := s2.Speak()

	fmt.Println(res1)
	fmt.Println(res2)

	// method 2
	saySomething(dog)
	saySomething(human)
}
