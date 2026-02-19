package main

import "fmt"

// package scope variable declaration
var speed int

func main() {

	// group variable declaration
	var (
		name string = "amir"
		age  int    = 20
	)

	// short variable declaration

	lastName := "Xayatpour"
	fmt.Printf("%s %s age is %d", name, lastName, age)
}
