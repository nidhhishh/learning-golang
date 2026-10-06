package main

import "fmt"

func product(result chan int, a, b float64) {
	productof := a * b
	result <- int(productof)
}

func main() {
	result := make(chan int)
	go product(result, 12, 1)
	ans := <-result

	fmt.Println(ans)

}
