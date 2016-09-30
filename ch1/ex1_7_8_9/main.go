package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, `http://`) {
			url = `http://` + url
		}
		resp, err := http.Get(url)
		status := resp.StatusCode
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		r, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		fmt.Printf("Status Code: %d", status)
		fmt.Printf("%s%v", r, err)
	}
}