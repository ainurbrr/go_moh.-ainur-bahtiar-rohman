package main

import "fmt"

func SimpleEquation(A, B, C int) {
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000; y++ {
			z := A - x - y
			if  x*y*z == B && x*x+y*y+z*z == C {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquation(1, 2, 3)  //no solution
	SimpleEquation(6, 6, 14) // 1 2 3
}
