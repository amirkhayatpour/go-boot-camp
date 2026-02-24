package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: [feet]")
		return
	}

	arg := os.Args[1]
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n", arg)
		return
	}
	result := convFeetToMetre(feet)
	fmt.Printf("%g feet is %g meters.\n", feet, result)
}

func convFeetToMetre(n float64) float64 {
	return n * 0.3048
}
