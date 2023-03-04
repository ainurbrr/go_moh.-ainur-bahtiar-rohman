package main

import "fmt"

func caesar(offset int, input string) string {

	outputCaesarCipher := []rune{}

	for _, character := range input {
		character = 97 + (character-97+rune(offset))%26
		outputCaesarCipher = append(outputCaesarCipher, character)
	}

	return string(outputCaesarCipher)

}

func main() {
	fmt.Println(caesar(3, "abc"))                           //def
	fmt.Println(caesar(2, "alta"))                          //cnvc
	fmt.Println(caesar(10, "alterraacademy"))               //kvdobbkkmknow
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    //bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl
}
