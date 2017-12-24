package main

import (
	"fmt"

	"github.com/devincheca/string"
)

func main() {
	for i := 0; i < 3; i++ {
		if i % 2 == 0 {
			fmt.Println(string.Reverse("Hello, world"))
		} else {
			fmt.Println("Hello, world")
		}
	}
}
