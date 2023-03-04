package main

import (
	"fmt"
)



func primeX(number int) int {
	count := 0
	i := 2

	for {
		is_prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			count++
		}
		if count == number {
			return i
		}
		i++
	}
}

func main() {
	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29
}
