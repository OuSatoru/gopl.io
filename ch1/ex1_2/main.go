//to print No. and value

package main

import (
	"os"
	"fmt"
)

func main() {
	s := ""
	for k, v := range os.Args{
		s = fmt.Sprintf("No.: %d, arg: %s", k, v)
		fmt.Println(s)
	}
}