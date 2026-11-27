package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func writeTestFile(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close %s: %w", path, closeErr)
		}
	}()

	if _, err := f.Write([]byte("Anup")); err != nil {
		return fmt.Errorf("write bytes to %s: %w", path, err)
	}
	if _, err := f.WriteString("\nThis is new line"); err != nil {
		return fmt.Errorf("write string to %s: %w", path, err)
	}
	return nil
}

func main() {
	// CREATE AND WRITE INTO FILE
	if err := writeTestFile("test.txt"); err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	if _, err := os.Stat("test.txt"); err == nil {
		fmt.Println("File exists")
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("Error checking file:", err)
	}

	// READ FROM FILE
	barr1, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(barr1))

	// DELETE A FILE
	if err := os.Remove("test.txt"); err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted")
	}

	// Write json to a file
	p := Person{Name: "Anup", Age: 28}
	data, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	if err := os.WriteFile("person.txt", data, 0644); err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	// Read json from a file
	fd, err := os.ReadFile("person.txt")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	var checkPerson Person
	if err := json.Unmarshal(fd, &checkPerson); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Println("Data from a file - ", checkPerson)
}
