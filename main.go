package main

import (
	"fmt"
	"sync"
	"time"
)

type remoteUpdate struct {
	ticker bool
	freq   time.Duration
}

func main() {
	quit := make(chan bool)

	update := remoteUpdate{}

	barrier := sync.WaitGroup{}
	go func(barrier *sync.WaitGroup) {
		defer barrier.Done()
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				if update.ticker {
					ticker.Stop()
					fmt.Println("re-scheduling")
					update.ticker = false
					ticker = time.NewTicker(update.freq)
				} else {
					fmt.Println("ticking")

				}
			case <-quit:
				ticker.Stop()
				fmt.Println("stopped!")
				return
			}
		}
	}(&barrier)

	time.Sleep(3 * time.Second) // simulate runtime

	fmt.Println("accelerate")

	update.ticker = true
	update.freq = 250 * time.Millisecond

	time.Sleep(4 * time.Second) // simulate runtime

	quit <- true
}
