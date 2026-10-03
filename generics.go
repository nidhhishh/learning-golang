package main

import "fmt"

// generic is used to pass multiple types
// of datatypes in a single function

// --------------------------------------------------------
/* in this function we can only pass int type string */

// func printslice(slice []int) {
// 	for i, item := range slice {
// 		fmt.Println(i, item)
// 	}
// }

// --------------------------------------------------------
func globalslice[G any](slice []G) {
	for _, item := range slice {
		fmt.Println(item)
	}
}

func main() {

	// num := []int{23, 32, 32}
	nums := []string{"yey", "yay"}

	// printslice(int1)

	globalslice(nums)
}
