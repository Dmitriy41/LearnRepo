// Package main переворачивает строку Hello, OTUS!
package main

// nolint:depguard
import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	// Place your code here.
	myStr := "Hello, OTUS!"
	fmt.Print(reverse.String(myStr))
}
