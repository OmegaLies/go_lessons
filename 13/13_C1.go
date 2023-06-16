package main

import (
	"fmt"
	"sync"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	count := 2
	var wg sync.WaitGroup
	wg.Add(count)
	go ponger(&wg, pings, pongs)
	go pinger(&wg, pings, pongs, numPings)
	wg.Wait()
}

// don't touch below this line

func pinger(wg *sync.WaitGroup, pings, pongs chan struct{}, numPings int) {
	defer wg.Done()
	go func() {
		sleepTime := 50 * time.Millisecond
		for i := 0; i < numPings; i++ {
			fmt.Println("ping", i, "sent")
			pings <- struct{}{}
			time.Sleep(sleepTime)
			sleepTime *= 2
		}
		close(pings)
	}()
	i := 0
	for range pongs {
		fmt.Println("pong", i, "got")
		i++
	}
	fmt.Println("pongs done")
}

func ponger(wg *sync.WaitGroup, pings, pongs chan struct{}) {
	defer wg.Done()
	i := 0
	for range pings {
		fmt.Println("ping", i, "got", "pong", i, "sent")
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}
