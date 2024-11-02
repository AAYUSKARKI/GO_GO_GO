package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main() {
	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}
