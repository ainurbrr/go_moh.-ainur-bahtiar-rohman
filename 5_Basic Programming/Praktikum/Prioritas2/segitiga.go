package main

import "fmt"

func main() {
	fmt.Println("===Segitiga Asterik===")

	var ukuran int
	fmt.Printf("Masukkan Ukuran segitiga asterik : ")
	fmt.Scanln(&ukuran)

	for i := 1; i <= ukuran; i++ {
		for j := ukuran; j >= i; j-- {
			fmt.Printf(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println("")
	}
}
