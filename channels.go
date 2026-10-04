package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func goprocess(numchannel chan int) {

	for num := range numchannel {
		fmt.Println("processing no.", num)
		time.Sleep(time.Second)
	}

}

func main() {

	numchannel := make(chan int)

	go goprocess(numchannel)

	for {
		numchannel <- rand.IntN(20)
	}

	// channelmessage := make(chan string)

	// channelmessage <- "heyo famm"

	// msg := <-channelmessage

	// fmt.Println(msg)
}
