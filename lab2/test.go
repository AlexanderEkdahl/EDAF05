package main

import (
	"fmt"
	"strings"
)

func main() {
	// An artificial input source.
	const input = "other there\nother their"
	scanner := strings.NewReader(input)

	var a, b string

	for _, err := fmt.Fscanln(scanner, &a, &b); err != nil; {
		fmt.Println(a)
		fmt.Println(b)
	}
}
