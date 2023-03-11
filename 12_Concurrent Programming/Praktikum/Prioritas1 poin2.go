package main

import (
	"fmt"
	"time"
)

func kelipatan3(ch chan int) {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
	close(ch)

}

func main() {
	ch := make(chan int)
	go kelipatan3(ch)

	time.Sleep(10 * time.Second)
	for c := range ch {
		fmt.Println(c)
	}
}
