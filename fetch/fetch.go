package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		start := time.Now()
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: odczytywanie %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n%.2fs\t%s\t%s\t\n"+
			"=======================================================================================\n",
			time.Since(start).Seconds(), url, response.Status)
	}
	fmt.Printf("Total time: %.2fs\n", time.Since(start).Seconds())
}
