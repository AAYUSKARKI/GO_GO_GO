package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Paying", amount, "using stripe")
}
func main() {
	stripePayment := payment{stripe{}}
	stripePayment.makePayment(100)
}
