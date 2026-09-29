package main

import "fmt"

func counterAttack() func() int {
	count := 0

	return func() int {
		count += 2
		return count
	}
}

func main() {

	increaseBytwo := counterAttack()

	fmt.Println(increaseBytwo())
	fmt.Println(increaseBytwo())
}
