package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	name := "amir"
	password := "1234"

	inputName := os.Args[1]
	inputPassword := os.Args[2]

	if inputName != name {
		fmt.Printf("Access denied for \"%s\"", inputName)
	} else if inputName == name && inputPassword != password {
		fmt.Println("Invalid password for \"amir\"")
	} else if inputName == name && inputPassword == password {
		fmt.Println("Access granted to \"amir\"")
	}
}
