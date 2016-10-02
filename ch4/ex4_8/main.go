package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	count := make(map[rune]int)
	kind := map[string]int{"Letter": 0, "Number": 0, "Chinese": 0, "other": 0}
	invalid := 0

	file, err := os.Open("read.txt")
	if err != nil {
		fmt.Errorf("Read file error: %s\n", err)
		return
	}
	defer file.Close()
	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		count[r]++
		if unicode.Is(unicode.Scripts["Han"], r) {
			kind["Chinese"]++
		} else if unicode.IsNumber(r) {
			kind["number"]++
		} else if unicode.IsLetter(r) {
			kind["Letter"]++
		} else {
			kind["other"]++
		}
	}
	fmt.Print("rune\tcount\n")
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("kind\tnumber")
	for k, v := range kind {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
