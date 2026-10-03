package main

import "fmt"

func multiply(nums ...int) int {
	total := 1

	for _, nums := range nums {
		total = total * nums
	}

	return total
}

func main() {

	fmt.Println()

}
