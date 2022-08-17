package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestChannelsTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	ch := make(chan int)
	defer cancel()

	go func() {
		t.Log("Starting goroutine")

		time.Sleep(time.Millisecond * 200)

		t.Log("Ending goroutine")
		ch <- 1
	}()

	select {
	case <-ctx.Done():
		t.Logf("Context cancelled before reaching the end as expected")
	case <-ch:
		t.Fatal("Reached end of function but should have timed out")
	}
}

func TestEarlyCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)

	_, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	_, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	ch := make(chan int, 2)

	go func() {
		ch <- 1
	}()

	go func() {
		time.Sleep(300)
		ch <- 2
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case val := <-ch:
		if val != 1 {
			log.Fatal("Only expecting 1")
		}
		t.Log("Sub1 done before parent context")
	case <-ctx.Done():
		t.Fatal("Sub1 wasn't done before parent context")
	}

	cancel()

	select {
	case val := <-ch:
		if val != 2 {
			log.Fatal("Only expecting 2")
		}
		t.Fatal("Sub2 also should have cancelled before")
	case <-ctx.Done():
		t.Log("Parent cancelled all childs context before sub2 had a chance")
	}
}
