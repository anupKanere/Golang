package main

import "fmt"

// Global Variables
var globalName = "Global Name"
var globalAge = 30

const Pi = 3.14

var defaultInt int
var defaultString string

func main() {

	// Local variable (typed)
	var name string = "Anup"
	var age int = 28

	// Local variable (untyped)
	var country = "India"

	// Short declaration (only inside functions)
	language := "Golang"

	// Reassigning value
	name = "Anup Kanere"

	fmt.Println("Global Name : ", globalName)
	fmt.Println("Global age : ", globalAge)
	fmt.Println("PI global : ", Pi)
	fmt.Println("Local name : ", name)
	fmt.Println("local age : ", age)
	fmt.Println("local country : ", country)
	fmt.Println("local language (short declaration) : ", language)

	fmt.Println("Default Int : ", defaultInt)
	fmt.Println("Default String : ", defaultString)

}
