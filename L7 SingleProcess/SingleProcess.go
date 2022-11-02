package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

var locked bool = false

var n int = 1000000

var safe, unsafe int

func main() {
	fmt.Println("starting...")
	for k := 0; k < n; k++ {
		go setLeader()
		go unsafeLeader()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(safe)
	fmt.Println(unsafe)
}

func setLeader() {
	mu.Lock()         // Lock resource when available
	defer mu.Unlock() // Unlock resource once done

	// critical function, can only do one thing at a time
	safe++
}

func unsafeLeader() {
	unsafe++
}
