package main

import (
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	// create channel for recipients
	recipentChannel := make(chan Recipient)
	// start producer goroutine
	go func() {
		loadRecipients("./emails.csv", recipentChannel)
	}()
	// wait group to wait for all workers to finish
	var wg sync.WaitGroup
	// number of workers
	workerCount := 5
	// start workers
	for i := 1; i <= workerCount; i++ {
		// increment wait group counter
		wg.Add(1)
		go eamilWorker(i, recipentChannel, &wg)
	}
	// wait for all workers to finish
	wg.Wait()
}
