package main

import (
	"sync"
	"testing"
)

func TestUnbufferedChannels(t *testing.T) {
	ch := make(chan int)

	go func() { // Unbuffered channels need to run in separate coroutines as you need to push and read simultaneously
		ch <- 1
	}()

	val := <-ch

	if val != 1 {
		t.Fatalf("Invalid value read from channel %v", val)
	}
}

func TestBufferedChannels(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1 // Can only push 1 to the channel as that is the buffer limit before getting deadlock
	// ch <- 1 // this will deadlock

	val := <-ch
	// val := <- ch // this will deadlock

	if val != 1 {
		t.Fatalf("Invalid value read from channel %v", val)
	}
}

func TestLoopOverChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // Channel needs to avoid deadlock and notify reader
	}()

	for val := range ch {
		t.Logf("value read %v", val)
	}
}

func TestSyncWithWaitGroups(t *testing.T) {
	ch := make(chan int)

	wg := sync.WaitGroup{} // We can use a wg to sync when coroutines are finished

	go func() { // We run on a coroutine to avoid locking the main thread and read data as goes in
		for i := 0; i < 100; i++ {
			wg.Add(1) // We track we started a new coroutine
			go func(val int) {
				defer wg.Done() // We indicate that this coroutine is finished with defer to execute this last, act as finally also in panic
				ch <- val
			}(i)
		}

		wg.Wait() // Wait until all routines are finished
		close(ch)
	}()

	for val := range ch {
		t.Logf("value read %v", val)
	}
}
