package main

import (
	"fmt"
)

func hej() {
	fmt.Println("hej")
	go func() {
		fmt.Println("hej")
	}()
	fmt.Println("hej")
}
