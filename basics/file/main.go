package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// CREATE AND WRITE INTO FILE
	barr := "Anup"
	f, _ := os.Create("test.txt")
	f.Write([]byte(barr))
	f.WriteString("\nThis is new line")
	f.Close()

	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File exists")
	}

	// READ FROM FILE
	barr1, _ := os.ReadFile("test.txt")
	fmt.Println(string(barr1))

	// DELETE A FILE
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted")
	}

	// Write json to a file
	p := Person{"Anup", 28}
	data, _ := json.MarshalIndent(p, "", " ")
	os.WriteFile("person.txt", data, 0644)

	// Read json from  a file
	fd, _ := os.ReadFile("person.txt")
	var checkPerson Person
	json.Unmarshal(fd, &checkPerson)
	fmt.Println("Data from a file - ", checkPerson)

}
