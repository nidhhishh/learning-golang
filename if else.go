package main

import "fmt"

func main() {

	// 	age := 12

	// 	if age >= 18 {
	// 		fmt.Println("you are an adult")
	// 	} else if age >= 12 {
	// 		fmt.Println("you are a teen ")
	// 	} else {
	// 		fmt.Println("you are a kid")
	// 	}

	/* Write a Go program that takes a person's age and prints:

	0–12 → "Child"
	13–17 → "Teenager"
	18–59 → "Adult"
	60+ → "Senior Citizen" */

	age := 132434

	if age >= 0 && age <= 12 {
		fmt.Println("you are a child")
	} else if age >= 13 && age <= 17 {
		fmt.Println("you are a teenager")
	} else if age >= 18 && age <= 59 {
		fmt.Print("you are a adult")
	} else {
		fmt.Println("you are a senior citizen")
	}

}
