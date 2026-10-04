package main

import (
	"fmt"
	"sync"
)

//   wait groups in goroutine
// wait grp should be equal to number of go tasks

func running1(pace float64, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Println("your pace is: ", pace)
}

func main() {

	var wait sync.WaitGroup
	for i := 0.00; i <= 10.43; i++ {
		wait.Add(1)
		go running1(i, &wait)
	}

	wait.Wait()

}
