package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0
	for _, num := range s {
		fmt.Println(num)
		sum += num
	}

	fmt.Println("Sum:", sum)
}
