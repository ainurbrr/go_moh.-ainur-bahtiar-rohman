package main

import "fmt"

func main() {
	fmt.Println("== Fizz Buzz ===")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf(" FizzBuzz")
		} else if i%5 == 0 {
			fmt.Printf(" Buzz")
		} else if i%3 == 0 {
			fmt.Printf(" Fizz")
		} else {
			fmt.Print(" ", i)
		}
	}
}

// buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”. Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….
