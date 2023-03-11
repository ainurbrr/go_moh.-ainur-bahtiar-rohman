package main

import (
	"fmt"
	"strconv"
)

func representasiBiner(n int) []string {
	var ans []string

	for i := 0; i < n+1; i++ {
		ans = append(ans, strconv.FormatInt(int64(i), 2))
	}

	return ans
}

func main() {
	fmt.Println(representasiBiner(2)) //[0 1 10]
	fmt.Println(representasiBiner(3)) //[0 1 10 11]
}