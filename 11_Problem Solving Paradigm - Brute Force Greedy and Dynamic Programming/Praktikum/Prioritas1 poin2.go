package main

import (
	"fmt"
)

func pascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func main() {
	fmt.Println(pascalTriangle(5)) //[[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
