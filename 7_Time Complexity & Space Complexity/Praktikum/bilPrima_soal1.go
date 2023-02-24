package main

import (
	"fmt"
)


func primeNumber(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Print(n, " Bukan Bilangan Prima maka ")
			return false
		}
	}
	fmt.Print(n, " Bilangan Prima maka ")
	return true
}
//O(sqrt(n)) lebih cepat daripada O(n) atau O(n/2)
//Kompleksitas waktu dari fungsi diatas adalah O(sqrt(n)) karena perulangan hanya akan dilakukan sebanyak akar kuadrat dari n
//n adalah angka yang akan ditentukan bilangan prima atau bukan
//fungsi akan mencari faktor dari bilangan n dengan melakukan perulangan dari 2 sampai akar kuadrat dari n
// Pada setiap iterasi perulangan, kita memeriksa apakah bilangan n habis dibagi dengan i yang sedang diperiksa. Jika ditemukan suatu i yang dapat membagi bilangan n, maka bilangan n bukan bilangan prima. Sebaliknya, jika tidak ditemukan faktor yang dapat membagi bilangan n, maka bilangan n adalah bilangan prima.


func main() {
	fmt.Println(primeNumber(1000000007)) //true
	fmt.Println(primeNumber(13))         //true
	fmt.Println(primeNumber(17))         //true
	fmt.Println(primeNumber(20))         //false
	fmt.Println(primeNumber(35))         //false
}
