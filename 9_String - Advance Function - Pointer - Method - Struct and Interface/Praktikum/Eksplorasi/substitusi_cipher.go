package main

import (
	"fmt"
	"os"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	nameEncode = CipherSubstitute(s.name)
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	if s.nameEncode == "" {
		s.Encode()
	}
	nameDecode = CipherSubstitute(s.nameEncode)
	return nameDecode
}

func CipherSubstitute(s string) string {
	result := []rune{}
	for _, val := range s {
		result = append(result, 'z'-((val)-'a'))
	}
	return string(result)
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a
	for {
		fmt.Print("\n[1] Encrypt \n[2] Decrypt\n[3] Exit \nChoose your menu? ")
		fmt.Scan(&menu)

		if menu == 1 {
			fmt.Print("\nInput Student Name: ")
			fmt.Scan(&a.name)
			fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
		} else if menu == 2 {
			fmt.Print("\nInput Student Name: ")
			fmt.Scan(&a.name)
			fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
		} else {
			os.Exit(0)
		}
	}
}
