package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	value   = 0
	rwMutex sync.RWMutex
)

func readValue() {
	// RWMutex lock for reading (1 to N goroutines / reading)
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// Critical section for reading
	fmt.Printf("Read value: %d\n", value)
}

func writeValue(val int) {
	// RWMutex lock for writing (exclusive)
	rwMutex.Lock()
	defer rwMutex.Unlock()

	// Critical section / writing
	value = val
	fmt.Printf("Wrote value: %d\n", value)
}

func main() {

	go writeValue(100)

	for i := 0; i < 50; i++ {
		go readValue()
	}

	time.Sleep(1 * time.Second)
}
