package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {

	myStr := "Hello, OTUS!"
	fmt.Println(reverse.String(myStr))

}
