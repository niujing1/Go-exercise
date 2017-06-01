package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() { // start 50 goroutines
			for {
				time.Sleep(time.Millisecond)
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
