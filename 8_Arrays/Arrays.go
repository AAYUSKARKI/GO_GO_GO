package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(len(arr))

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	var arr4 = [...]string{1: "Hello", 4: "World"}
	fmt.Println(arr4)

	var arr5 = [...]string{"Hello", "World"}
	fmt.Println(arr5)

	var arr6 = [5]string{"Hello", "World"}
	fmt.Println(arr6)

	var arr7 = [5][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr7)
	fmt.Println(len(arr7))
	for i := 0; i < len(arr7); i++ {
		fmt.Println(arr7[i])
	}
}
