package main

import "fmt"

// type MarksEvaluation int

// const (
// 	Exam_in_process MarksEvaluation = iota
// 	Exam_is_completed
// 	Exam_is_evaluated
// )

type OrderStatus string

const (
	R OrderStatus = "recieved"
	C             = "confirmed"
	P             = "processing"
	D             = "delivered"
)

// func MarksEvaluationStatus(status MarksEvaluation) {
// 	fmt.Println("your exam sheets are curently being:", status)
//}

func OrderStatusOn(status OrderStatus) {
	fmt.Println("your order status is:", status)
}

func main() {
	// MarksEvaluationStatus(Exam_is_completed)

	OrderStatusOn(C)

}
