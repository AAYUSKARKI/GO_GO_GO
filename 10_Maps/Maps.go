package main

import (
	"fmt"
	"maps"
)

func main() {

	var person = map[string]string{
		"name": "John Doe",
		"age":  "20",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	var person2 = make(map[string]string)

	person2["name"] = "John Doe"

	fmt.Println(person2)

	delete(person, "age")

	fmt.Println(person)

	for key, value := range person {
		fmt.Println(key, value)
	}

	//check equality
	fmt.Println(maps.Equal(person, person2))
}
