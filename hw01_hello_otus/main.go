// Package main переворачивает строку Hello, OTUS!
package main


import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	// Place your code here.
	fmt.Print(reverse.String("Hello, OTUS!"))
}
