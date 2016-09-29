// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.

//ex1.3: timeit with echo1
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	t := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(t))
}

//!-
