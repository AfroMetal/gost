package main

import (
	"fmt"
	"os"
)

func main() {
	var s, delimiter string

	for _, arg := range os.Args[1:] {
		s += delimiter + arg
		delimiter = " "
	}

	fmt.Println(s)
}