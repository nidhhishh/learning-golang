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

	result := multiply(2, 2, 2, 2)

	fmt.Println("your ans is", result)

}
