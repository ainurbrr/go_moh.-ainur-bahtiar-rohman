package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sentence string

	fmt.Printf("Apakah Palindrome?\nmasukkan kata : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence = scanner.Text()

	var i = 0
	var length int = len(sentence) - 1
	fmt.Println("captured : " + sentence)

	for i < length {
		i += 1
		length -= 1
		if sentence[i] != sentence[length] {
			fmt.Println("Bukan Palindrome!")
			break
		}
		fmt.Println("Palindrome")
		os.Exit(0)
	}

}
