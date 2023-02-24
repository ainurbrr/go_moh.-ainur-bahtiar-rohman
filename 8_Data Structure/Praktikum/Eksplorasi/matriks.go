package main

import "fmt"

func hitungSelisihDiagonal(matriks [][]int) int {
	var dariKiri int
	var dariKanan int
	var selisih int
	for i := 0; i < len(matriks); i++ {
		dariKanan += matriks[i][len(matriks)-1-i]
		dariKiri += matriks[i][i]
	}
	fmt.Println(dariKiri, "-", dariKanan)

	selisih = dariKiri - dariKanan
	return selisih
}

func main() {
	matrikPersegi := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	fmt.Println("Selisih dari diagonal adalah : ", hitungSelisihDiagonal(matrikPersegi))
}
