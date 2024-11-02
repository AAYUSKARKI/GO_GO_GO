package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	age       int
	email     string
	createdAt time.Time // nanoseconds precision
}

func newPerson(name string, age int, email string) person {
	return person{
		name:      name,
		age:       age,
		email:     email,
		createdAt: time.Now(),
	}
}

func (p *person) changeName(name string) {
	p.name = name
}
func main() {

	//create a new person
	p := person{
		name:      "John",
		age:       20,
		email:     "j@j.com",
		createdAt: time.Now(),
	}

	p2 := newPerson("Janeykc", 21, "j@j.com")
	fmt.Println(p, p.name, p.age, p.email, p.createdAt)
	fmt.Println(p2, p2.name, p2.age, p2.email, p2.createdAt)

	//change name
	p.changeName("Jane")
	fmt.Println(p, p.name, p.age, p.email, p.createdAt)

	//inline struct
	p3 := struct {
		name      string
		age       int
		email     string
		createdAt time.Time
	}{
		name:      "Jane",
		age:       21,
		email:     "j@j.com",
		createdAt: time.Now(),
	}
	fmt.Println(p3, p3.name, p3.age, p3.email, p3.createdAt)
}
