//print SHA256 default, and SHA384 SHA512, did not use flag

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Nothing typed.")
	} else if len(os.Args) == 2 {
		s1 := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("SHA256 of %s: %x\n", os.Args[1], s1)
	} else if os.Args[1] == "sha384" {
		s2 := sha512.Sum384([]byte(os.Args[2]))
		fmt.Printf("SHA384 of %s: %x\n", os.Args[2], s2)
	} else if os.Args[1] == "sha512" { //can't print sha512
		s3 := sha512.Sum512([]byte(os.Args[2]))
		fmt.Printf("SHA512 of %s: ", os.Args[3])
		fmt.Println(s3)
	} else if os.Args[1] == "sha256" {
		s1 := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("SHA256 of %s: %x\n", os.Args[1], s1)
	}
}
