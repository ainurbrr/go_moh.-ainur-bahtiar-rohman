package main

import "fmt"

func pow(x, n int) int {

	if n == 0 {
		return 1
	} else if n%2 == 0 {
		result := pow(x, n/2)
		return result * result
	} else {
		result := pow(x, (n-1)/2)
		return x * result * result
	}
}

//Kompleksitas waktu rogram menentukan perpangkatan diatas adalah O(log(n))
//int x merupakan base angka yang akan dipangkatkan sedangkan n adalah eksponen atau pangkatnya
//program ini menggunakan rekursif pemanggilan fungsi dirinya sendiri
//program ini jika pangkatnya = 0 maka akan me return 1
//pada dasarnya program ini membagi pangkatnya menjadi setengahnya karena sudah langsung dikalikan didalam rekursif
//jika pangkatnya ganjil maka dilakukan satu perkalian lagi dengan base angkanya

func main() {
	fmt.Println(pow(2, 3))  //8
	fmt.Println(pow(5, 3))  //125
	fmt.Println(pow(10, 2)) //100
	fmt.Println(pow(2, 5))  //32
	fmt.Println(pow(7, 3))  //343
}
