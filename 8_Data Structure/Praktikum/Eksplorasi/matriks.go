package main

import (
	"fmt"
	"math"
)

func hitungSelisihDiagonal(matriks [][]int) int {
	var dariKiri int
	var dariKanan int
	var selisih int
	for i := 0; i < len(matriks); i++ {
		dariKiri += matriks[i][i]
		dariKanan += matriks[i][len(matriks)-1-i]
	}
	fmt.Println(dariKiri, "-", dariKanan)

	selisih = dariKiri - dariKanan
	return int(math.Abs(float64(selisih)))
}

func main() {
	matriksPersegi := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	fmt.Println("Selisih dari diagonal adalah : ", hitungSelisihDiagonal(matriksPersegi))
}
