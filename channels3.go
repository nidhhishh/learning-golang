package main

import (
	"fmt"
	"time"
)

// import (
// 	"fmt"
// 	"math/rand/v2"
// 	"time"
// )

func emailsender(email chan string, done chan bool) {
	defer func() { done <- true }()

	for mails := range email {

		fmt.Println("sending mails to", mails)
		time.Sleep(time.Second)
	}
}

func main() {
	email := make(chan string, 100)
	done := make(chan bool)

	go emailsender(email, done)

	for i := 0; i < 20; i++ {
		email <- fmt.Sprintf("%dgmail.com", i)
	}
	fmt.Println("emails are sent successfully... ")

	close(email)

	<-done
}

// func goprocess(numchannel chan int) {

// 	for num := range numchannel {
// 		fmt.Println("processing no.", num)
// 		time.Sleep(time.Second)
// 	}

// }

// func main() {

// 	numchannel := make(chan int)

// 	go goprocess(numchannel)

// 	for {
// 		numchannel <- rand.IntN(20)
// 	}

// 	// channelmessage := make(chan string)

// 	// channelmessage <- "heyo famm"

// 	// msg := <-channelmessage

// 	// fmt.Println(msg)
// }
