package main

import (
	"fmt"
	"time"
)

func kelipatanX(x int) {
	for i := 1; i <= x; i++ {
		fmt.Println(i * x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go kelipatanX(5)
	time.Sleep(20 * time.Second)
}
