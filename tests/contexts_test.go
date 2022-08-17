package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestChannelsTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	inFunc := false

	go func() {
		t.Log("Starting goroutine")

		select {
		case <-time.After(200 * time.Millisecond):
			t.Log("Ending goroutine")
			inFunc = true
			cancel()
		case <-ctx.Done():
			cancel()
		}
	}()

	<-ctx.Done()

	if inFunc {
		log.Fatal("Shouldn't reach in fatal")
	}
}

func TestEarlyCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	ctx1, cancel1 := context.WithCancel(ctx)
	defer cancel1()
	ctx2, cancel2 := context.WithCancel(ctx)
	defer cancel2()

	reached1, reached2 := false, false

	go func() {
		reached1 = true
		cancel1()
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		reached2 = true
		cancel2()
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case <-ctx.Done():
		log.Fatal("Ctx1 should finish first")
	case <-ctx1.Done():
		if !reached1 {
			log.Fatal("it should have reached 1")
		}
		log.Print("Ctx1 is done early as expected")
	}

	select {
	case <-ctx.Done():
		if reached2 {
			log.Fatal("Ctx: It shouldn't have reached 2")
		}
		log.Print("Ctx finished first as expected")
	case <-ctx2.Done():
		if reached2 {
			log.Fatal("Ctx2: It shouldn't have reached 2")
		}
		log.Print("Ctx2 is finished due to cascade effect")
	}
}
