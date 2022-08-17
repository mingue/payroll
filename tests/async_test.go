package main

import (
	"log"
	"testing"
	"time"
)

func TestAsyncAwaitWithCompletion(t *testing.T) {
	reached := false

	task := async(func() {
		time.Sleep(100 * time.Millisecond)
		reached = true
	})

	task.Wait()

	if !reached {
		log.Fatal("Should have waited to reach value")
	}
}

func TestAsyncAwaitWithReturn(t *testing.T) {

	task := asyncR(func() string {
		time.Sleep(100 * time.Millisecond)
		return "String"
	})

	var s = task.Get()
	
	if s!= "String"{
		log.Fatal("Not working")
	}
}

// Run a func as a courutine returning a task that can be awaited for completion
func async(f func()) Task {
	return Task{f}
}

type Task struct {
	f func()
}

// Blocks the current Goroutine to obtain the value from func
func (a Task) Wait() {
	ch := make(chan int)

	go func() {
		a.f()
		ch <- 0
		close(ch)
	}()

	<- ch
}

// Run a func as a courutine returning a task that can be awaited to obtain the result
func asyncR[T any](f func() T) TaskReturn[T] {
	return TaskReturn[T]{f}
}

type TaskReturn[T any] struct {
	f func() T
}

// Blocks the current Goroutine to obtain the value from func
func (a TaskReturn[T]) Get() T{
	ch := make(chan T)

	go func() {
		ch <- a.f()
		close(ch)
	}()

	value := <- ch
	return value
}