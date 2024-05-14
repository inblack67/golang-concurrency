package main

import (
	"fmt"
	"sync"
	"time"
)

func getUserName() string {
	time.Sleep(time.Millisecond * 100)
	return "inblack67"
}

func getUserLikes(username string, ch chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	ch <- 10
	wg.Done()
}

func getUserFollowers(username string, ch chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	ch <- 43
	wg.Done()
}

func main2() {
	start := time.Now()
	username := getUserName()

	ch := make(chan int, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go getUserLikes(username, ch, wg)
	go getUserFollowers(username, ch, wg)

	wg.Wait()
	close(ch)

	for r := range ch {
		fmt.Println(r)
	}

	end := time.Since(start)
	fmt.Println("took : ", end)
}
