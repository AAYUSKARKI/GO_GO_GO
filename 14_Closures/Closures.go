package main

import (
	"fmt"
)

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}
