package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool) // a set of strings
	file, err := os.Open("read.txt")
	if err != nil {
		fmt.Errorf("Read file error: %s\n", err)
		return
	}
	defer file.Close()
	input := bufio.NewScanner(file)
	//input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
