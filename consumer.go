package main

import (
	"fmt"
	"sync"
)

func eamilWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	// signal done when the function exits
	defer wg.Done()
	// consume from the channel
	for recipient := range ch {
		fmt.Println(id, recipient)
	}
}
