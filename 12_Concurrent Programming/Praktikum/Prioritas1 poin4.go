package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu sync.Mutex

	sum := 0

	for i := 1; i <= 50; i++ {
		go func(num int) {
			// melakukan locking mutex sebelum mengakses variabel sum
			mu.Lock()
			// melakukan unlocking mutex setelah mengakses variabel sum
			defer mu.Unlock()
			sum += num
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Sum of numbers 1 to 50 is:", sum)
}
