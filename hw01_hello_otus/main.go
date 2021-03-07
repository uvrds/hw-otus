package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	word := "Hello, OTUS!"
	fmt.Printf("%s", stringutil.Reverse(word))
}
 
