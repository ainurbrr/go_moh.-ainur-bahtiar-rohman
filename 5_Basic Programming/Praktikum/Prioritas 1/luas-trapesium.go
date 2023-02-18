package main

import "fmt"

func main() {
	fmt.Println("===Menentukan Luas Trapesium===")
	var luas_trapesium, bawah, atas, tinggi float64

	fmt.Print("Masukkan panjang sisi sejajar atas : ")
	fmt.Scanln(&atas)
	fmt.Print("masukkan panjang sisi sejajar bawah : ")
	fmt.Scanln(&bawah)
	fmt.Print("Masukkan tinggi trapesium : ")
	fmt.Scanln(&tinggi)

	luas_trapesium = (atas + bawah) * tinggi * 0.5

	fmt.Println("Luas Trapesium adalah ", luas_trapesium)

}
