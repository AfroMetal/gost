package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dupa2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, where := range counts {
		if len(where) > 1 {
			n := 0
			for key, i := range where {
				fmt.Printf("%s:%d\t", key, i)
				n += i
			}
			fmt.Printf("total:%d\t%s\n", n, line)
		}
	}
}
func countLines(file *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		mm, ok := counts[input.Text()]
		if !ok {
			mm = make(map[string]int)
			counts[input.Text()] = mm
		}
		mm[file.Name()]++
	}
}
