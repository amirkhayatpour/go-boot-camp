package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	fmt.Println(os.Args)
	dir, file := path.Split(os.Args[0])
	fmt.Println(dir)
	fmt.Println(file)
}
