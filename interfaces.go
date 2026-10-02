package main

import "fmt"

type GlobalPayment interface {
	pay(amount float64)
}

type payment struct {
	gateway GlobalPayment
}

func (p payment) makepayment(amount float64) {
	// phonepeGw := phonepe{}
	// phonepeGw.pay(amount)

	// paytmGw := paytm{}
	// paytmGw.pay2(amount)
	p.gateway.pay(amount)
}

type phonepe struct{}

func (p phonepe) pay(amount float64) {
	fmt.Println("making your payment with phonepe", amount)
}

type paytm struct{}

func (p paytm) pay(amount float64) {
	fmt.Println("making your payment using paytm", amount)
}

func main() {
	phonepeGw := phonepe{}
	// paytmGw := paytm{}

	newPayment := payment{
		gateway: phonepeGw,
		// gateway: paytmGw,
	}
	newPayment.makepayment(230)

}
