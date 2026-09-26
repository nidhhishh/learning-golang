package main

import (
	"fmt"
)

func main() {

	// a := 19
	// switch a {
	// case 18:
	// 	fmt.Println("eighteen day")
	// case 19:
	// 	fmt.Println("nineteen day")
	// default:
	// 	fmt.Println("its wrong day")

	// }

	// multiple swtich

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its weeknd")
	// default:
	// 	fmt.Println("its weekdy")

	// }

	// type swtich

	// whoAreU := func(i interface{}) {
	// 	switch i.(type) {
	// 	case int:
	// 		fmt.Println("its an integer")
	// 	case bool:
	// 		fmt.Println("its a boolean")
	// 	case float32:
	// 		fmt.Println("its a float ")

	// 	}

	// }
	// whoAreU(34.5)

	// vowel checker

	letter := 'e'
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("its a vowel")
	default:
		fmt.Println("constant")
	}
}
