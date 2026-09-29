package main

import "fmt"

func changevariable(a *int) {
	*a = 30
}

func main() {

	a := 25

	changevariable(&a)
	fmt.Println(&a)
	fmt.Println("number is changed", a)

	// used to change the value in its address, which is already defined

}
