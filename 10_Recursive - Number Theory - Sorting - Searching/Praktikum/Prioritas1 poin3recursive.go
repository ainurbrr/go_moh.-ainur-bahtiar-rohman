package main

import "fmt"

func isPrime(n int, i int) bool {
	if i == 1 {
		return true
	}

	if n%i == 0 {
		return false
	}

	return isPrime(n, i-1)
}

func primeX(n int) int {
	for i := 2; n >= 1; i++ {
		if isPrime(i, i-1) {
			if n == 1 {
				return i
			}
			n--
		}
	}
	return 0
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
