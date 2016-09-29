package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		r, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Printf("%s%v", r, err)
	}
}