package main

import "fmt"

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func main() {
	items1 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	printSlice(items1)
	items2 := []int{1, 2, 3, 4, 5}
	printSlice(items2)
	items3 := []bool{true, false, true, false, true}
	printSlice(items3)
	items4 := []string{"a", "b", "c", "d", "e"}
	printSlice(items4)
}
