package main

import "fmt"

//for -> only construct in go for looping

func odd(n int) bool {
	return n%2 == 1
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "is odd:", odd(i))
	}
}
