package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}
func main() {
	result := multiply(2, 3)
	fmt.Println(result)

	result = divide(2, 3)
	fmt.Println(result)
}
