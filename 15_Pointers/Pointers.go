package main

import "fmt"

func changeByPointer(a *int) {
	*a = 233
}

func main() {
	a := 20

	changeByPointer(&a)

	fmt.Println(a)
}
