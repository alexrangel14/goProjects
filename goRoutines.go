package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Go Summary:

// Definition:
// They are lightweight concurrent execution units that allow you to run
// functions concurrently.They enable concurrent programming and help
// you achieve parallelism.

// Key features:

// Lightweight:

// Concurrent Execution:
// Enables multiple functions can run at the same time and independently.
// They can do this without explicitly managing threads or blocking the
// main program.

// Concurrency vs Parallelism:

// sync.WaitGroup ensures all go routines finish before exiting
var wg = sync.WaitGroup{}

// counter is 2nd concurrent variable
var counter = 0

// RWMutex allows for any amount of things to be read in, but only
// one thing to be written to at a time.
var m = sync.RWMutex{}

func main() {
	// GOMAXPROCS changes the amount of threads used
	// Too many threads causing memory problems
	// Not enough slows performance down
	// While testing change its value to see how many
	// threads is optimal amount
	runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		// 'wg.Add(2) Allow multiple go routines to run
		wg.Add(2)
		// 'm.RLock()' Allow multiple go routines to read. Provides a
		// read lock for all subsequent go routines. It is a read
		// lock since it is taking in the shared variable 'counter'
		m.RLock()
		// Start a go rountine with 'sayHello'
		go sayHello()
		// 'm.Lock()' Ensures exclusive access to the writer operation.
		// Each go rountine needs its own lock/unlock.
		// Its a write lock since you're writing (manipulating) one
		// of the shared variable.
		m.Lock()
		// Start a go rountine with 'increment'
		go increment()
	}
	// 'wg.Wait()' Ensures that all go routines finished before moving on
	wg.Wait()
}

func sayHello() {
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	fmt.Printf("Hello #%v\n", counter)
	// Ends the lock allowing other go routines to access the resource
	m.RUnlock()
	//Signals that the go rountine is inished
	wg.Done()
}

func increment() {
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	counter++
	// Ends the lock allowing other go rountines to access the resource
	m.Unlock()
	// Signals the completion on the go routine
	wg.Done()
}
