package main

import "fmt"

type GlobalPayment interface {
	pay(amount float64)
}

// used to access multiple gateways at a single time

/*An interface is a type that defines a set of method signatures
 that a type must implement to satisfy that interface. It tells
 us what a type can do, rather than how it does it. Any struct or
 type that implements all the methods defined in an interface
 automatically satisfies it. Go does not require an explicit
 implements keyword. Interfaces help us write flexible, reusable,
 and maintainable code. For example, different payment methods
 like PhonePe and Google Pay can implement the same Payment
 interface in their own ways. */ 

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
