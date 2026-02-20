package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "Amir Xayatpour"
	fmt.Println(utf8.RuneCountInString(name))

}
