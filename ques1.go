package main

// Student Marks Analyzer

import "fmt"

func calculateAvarage(a, b, c, d float64) float64 {
	total := a + b + c + d
	avg := (total) / 4.00
	return avg

}

func totalMarks(a, b, c, d int) int {
	total := a + b + c + d
	return total
}

func main() {
	total1 := totalMarks(79, 95, 93, 90)
	fmt.Println("your total marks out of 400", total1)
	result := calculateAvarage(79, 95, 93, 90)
	fmt.Println("your percentage is:", result)
	if result >= 80 {
		fmt.Println("Excellent")
	} else if result >= 60 {
		fmt.Println("Good")
	} else {
		fmt.Println("Needs improvement")
	}

}
