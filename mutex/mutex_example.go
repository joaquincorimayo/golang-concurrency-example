package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex   sync.Mutex
)

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	// critical section
	counter++
	fmt.Printf("Incremented counter to %d\n", counter)
}

func main() {

	for i := 0; i < 50; i++ {
		go increment()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("Final counter value:", counter)
}
