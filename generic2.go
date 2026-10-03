package main

import "fmt"

// generic for structs

type game[T any] struct {
	elements []T
}

func main() {

	// game1 := game[string]{

	// 	elements: []string{"free fire",  "pubg"},
	// }

	// fmt.Println(game1)

	gameNo := game[int]{
		elements: []int{231, 834},
	}

	fmt.Println(gameNo)
}
