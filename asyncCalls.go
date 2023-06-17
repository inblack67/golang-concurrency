package main

import (
	"sync"
	"time"
)

func asyncCalls() {
	now := time.Now()
	responseChannel := make(chan string, 128)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go task1(responseChannel, wg)
	go task2(responseChannel, wg)
	go task3(responseChannel, wg)

	wg.Wait()

	close(responseChannel)

	for v := range responseChannel {
		println(v)
	}

	elapsed := time.Since(now)
	println(elapsed.Milliseconds())
}

func task1(respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 80)
	respCh <- "task 1 done"
	wg.Done()
}

func task2(respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 120)
	respCh <- "task 2 done"
	wg.Done()
}

func task3(respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	respCh <- "task 3 done"
	wg.Done()
}
