package main

import (
	"fmt"
	"maps"
)

func main() {
	// m := make(map[string]int)
	// m["go"] = 23
	// m["games"] = 21
	// fmt.Println(m["go"])

	m := map[int]string{2: "cars", 3: "bikes", 188: "cycles"}

	m2 := map[int]string{2: "cars", 3: "bikes", 188: "cycles"}

	fmt.Println(maps.Equal(m, m2))
	delete(m, 2)
	fmt.Println(m, m2)

}
