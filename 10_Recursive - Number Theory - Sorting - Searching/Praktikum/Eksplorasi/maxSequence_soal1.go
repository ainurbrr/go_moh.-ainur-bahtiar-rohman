package main

import "fmt"

func MaxSequence(arr []int) int {
	var maxTemp int
	maximumSequence := arr[0]
	for _, value := range arr {
		if value > maxTemp+value {
			maxTemp = value
		} else {
			maxTemp = maxTemp + value
		}

		if maximumSequence < maxTemp {
			maximumSequence = maxTemp
		}
	}
	return maximumSequence
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) //6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   //7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   //7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
