package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//{"king", "devil jin", "akuma", "eddie", "steve", "geese}

	fmt.Println(ArrayMerge([]string{"sergie", "jin"}, []string{"jin", "steve", "bryan"}))
	//{"sergie", "jin", "steve", "bryan"}

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu"}))
	//{"alisa", "yoshimitsu", "devil jin", "law"}

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	//{"devil jin", "sergei"}

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	//{"hwoarang"}

	fmt.Println(ArrayMerge([]string{}, []string{}))
	//{}
}