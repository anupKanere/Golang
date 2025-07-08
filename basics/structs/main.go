package main

import "fmt"

type Person struct {
	Name string
	Age  int
	city string
}

func (p Person) printPersonDetails() {
	fmt.Printf("Name = %s \n Age = %d \n city = %s \n", p.Name, p.Age, p.city)
}

func main() {

	person1 := Person{
		Name: "Anup", Age: 28, city: "Pune",
	}

	person2 := Person{"Akshay", 30, "Pune"}

	person1.printPersonDetails()
	person2.printPersonDetails()

	fmt.Printf("Name of first Person is %s \n", person1.Name)

}
