package main

import "fmt"

// func add(a, b int) int {
// 	return a + b

func carBrands() (string, string, string) {
	return "bmw", "audi", "porsche"
}

func main() {

	// ans := add(5, 6)
	// fmt.Println(ans)
	car1, car2, car3 := carBrands()
	fmt.Println(car1, car2, car3)
}
