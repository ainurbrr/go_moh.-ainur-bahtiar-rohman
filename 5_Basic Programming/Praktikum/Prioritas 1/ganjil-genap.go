package main

import "fmt"

func main()  {
	fmt.Println("===Menentukan Ganjil/Genap===")
	var angka int

	fmt.Print("Masukkan angka yang akan ditentukan : ")
	fmt.Scanln(&angka)

	if angka%2 == 0 {
		fmt.Println("Angka yang diinput adalah angka Genap")
	}else {
		fmt.Println("Angka yang diinput adalah angka Ganjil")
	}
}