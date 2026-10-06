package main

import (
	"fmt"
	"sync"
)

type carsSale struct {
	sold int

	hack sync.Mutex
}

func (c *carsSale) count(wg *sync.WaitGroup) {

	defer wg.Done()

	c.hack.Lock()
	c.sold += 1
	c.hack.Unlock()
}

func main() {

	var wg sync.WaitGroup

	carsolds := carsSale{sold: 0}

	for i := 0; i < 199; i++ {
		wg.Add(1)
		go carsolds.count(&wg)
	}

	wg.Wait()

	fmt.Println(carsolds.sold)

}
