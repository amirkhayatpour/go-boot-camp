package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	charNum := len(msg)
	repNum := strings.Repeat("!", charNum)
	final := repNum + strings.ToUpper(msg) + repNum
	fmt.Println(final)

}
