package main

import "fmt"

func main() {
	//uninitialized slice is nil
	var slice []int
	fmt.Println(slice == nil)

	//make slice
	slice = make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//append slice
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
