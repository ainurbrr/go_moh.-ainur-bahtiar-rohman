package main

import "fmt"

func main()  {

	fmt.Println("===Menentukan Grade untuk Nilai===")
	var nilai int
	fmt.Print("Masukkan nilai untuk ditentukan gradenya : ")
	fmt.Scanln(&nilai)

	if nilai >= 80 && nilai <=100{
		fmt.Println("Anda mendapat Grade A")
	}else if nilai >= 65 && nilai <= 79{
		fmt.Println("Anda mendapat Grade B")
	}else if nilai >= 50 && nilai <= 64{
		fmt.Println("Anda mendapat Grade C")
	}else if nilai >= 35 && nilai <= 49{
		fmt.Println("Anda mendapat Grade D")
	}else if nilai >= 0 && nilai <= 34{
		fmt.Println("Anda mendapat Grade E")
	}else{
		fmt.Println("Nilai Ivalid")
	}
}

