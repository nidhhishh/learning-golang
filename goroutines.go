package main

// goroutines are used to run the code faste in different segemnts cohrently

import (
	"fmt"
	"time"
)

func running(pace float64) {
	fmt.Println("your pace is: ", pace)
}

func main() {
	for i := 0.00; i <= 10.43; i++ {

		go running(i)
	}

	time.Sleep(time.Second)
}
