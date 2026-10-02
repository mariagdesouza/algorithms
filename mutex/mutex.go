package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func main() {
	fmt.Println("Hello, playground")
	//useMutex()
	useChannel()
}

func useMutex() {
	x = 0
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 50000; i++ {
		w.Add(1)
		//go increment(&w)
		//It is important to pass the address of the mutex here.
		//If the mutex is passed by value instead of passing the address,
		//each Goroutine will have its own copy of the mutex and the race condition will still occur.
		go incrementWithMutex(&w, &m)
	}
	w.Wait()
	fmt.Println(x)
	if x != 50000 {
		fmt.Println("Flawed output")
	}

}

func incrementUsingChannel(wg *sync.WaitGroup, ch chan struct{}) {
	var a struct{}
	ch <- a
	x = x + 1
	<-ch
	wg.Done()
}

func useChannel() {
	x = 0
	var w sync.WaitGroup
	// this has to be a buffered channel or it wont proceed and deadlock after
	ch := make(chan struct{}, 1)
	for i := 0; i < 50000; i++ {
		w.Add(1)
		//go increment(&w)
		//It is important to pass the address of the mutex here.
		//If the mutex is passed by value instead of passing the address,
		//each Goroutine will have its own copy of the mutex and the race condition will still occur.
		go incrementUsingChannel(&w, ch)
	}
	w.Wait()
	fmt.Println(x)
	if x != 50000 {
		fmt.Println("Flawed output")
	}

}
