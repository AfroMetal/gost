package main

import (
	"fmt"
	"os"
)

func main() {
	var s, delimiter string

	for i := 1; i < len(os.Args); i++ {
		s += delimiter + os.Args[i]
		if i == 1 {
			delimiter = " "
		}
	}

	fmt.Println(s)
}
