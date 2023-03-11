package main

import (
	"fmt"
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
	ch := make(chan int, 5)
	go kelipatan3(ch)

	for c := range ch {
		fmt.Println(c)
	}
}
