package main

import (
	"fmt"
	"slices"
)

func main() {
	// nums := []string{"i", "go", "park"}
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums := make([]int, 0, 5)
	// nums = append(nums, 2, 3)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))
	num := []string{"go", "golang", "c", "c++"}
	num1 := []string{"go", "golang", "c", "c++"}
	fmt.Println(slices.Equal(num, num1))

}
