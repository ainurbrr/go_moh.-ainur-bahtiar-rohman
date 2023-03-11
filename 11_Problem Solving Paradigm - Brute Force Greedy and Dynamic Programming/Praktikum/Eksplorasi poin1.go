package main

import (
	"fmt"
)

func convertRomawi(angka int) string {
	desimal := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var hasil string

	for i := 0; i < len(desimal); i++ {
		for desimal[i] <= angka {
			hasil += romawi[i]
			angka -= desimal[i]
		}
	}
	return hasil
}

func main() {
	fmt.Println(convertRomawi(4))    //IV
	fmt.Println(convertRomawi(9))    //IX
	fmt.Println(convertRomawi(23))   //XXIII
	fmt.Println(convertRomawi(2021)) //MMXXI
	fmt.Println(convertRomawi(1646)) //MDCXLVI
}
