package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func testRaces(t *testing.T) {
	var state int
	var mu sync.RWMutex
	for i := 0; i < 10; i++ {
		go func(i int) {
			// mu.RLock() // read lock but we are writng here
			mu.Lock()
			state += i
			mu.Unlock()
		}(i)
	}
}

func testRaces2(t *testing.T) {
	var state int32
	for i := 0; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
