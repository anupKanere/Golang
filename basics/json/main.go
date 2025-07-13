package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
	Phone   string
}

func main() {

	p1 := Person{
		Name:    "Anup",
		Address: "Ashta",
		Phone:   "123456789",
	}

	barr, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("JSON FORMAT in string representation- ", string(barr))

	// Convert []byte to dynamic JSON object
	var mapObj map[string]interface{}
	json.Unmarshal(barr, &mapObj)

	// Now you can access like JSON
	fmt.Println("Name from JSON object: ", mapObj["Name"])

	var p2 Person
	json.Unmarshal(barr, &p2)
	fmt.Println("Name of p2 - ", p2.Name)
}
