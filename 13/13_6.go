package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	var summaries int64
	summaries = 0
	ch := make(chan int64, n)
	go func() {
		fibonacci(n, ch)
	}()
	for item := range ch {
		summaries += item
	}
	fmt.Println(summaries)
}

// don't touch below this line

func test(n int) {
	fmt.Println(time.Now())
	concurrrentFib(n)
	fmt.Println(time.Now())
}

func main() {
	test(1000)
}

func fibonacci(n int, ch chan int64) {
	var x, y int64
	x, y = 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
