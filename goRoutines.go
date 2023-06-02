package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

//var mutex = sync.Mutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	//mutex.Lock()
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	fmt.Printf("Hello #%v\n", counter)
	//mutex.Unlock()
	wg.Done()
}

func increment() {
	//mutex.Lock()
	sum := 0
	for i := 0; i < 99999999; i++ {
		sum += i
	}
	counter++
	//	mutex.Unlock()
	wg.Done()
}
