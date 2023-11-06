package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	a := 1

	if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1")
	}

	a = 3

	if a == 1 {
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1", a)
	}

	if b := rand.Intn(3); b == 2 {
		fmt.Println("b is 2", b)
	} else {
		fmt.Println("b is not 2", b)
	}
}
