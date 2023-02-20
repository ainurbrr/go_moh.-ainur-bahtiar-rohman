package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	result := []string{}
	
	checker := make(map[string]bool)

	for _, element := range arrayA {
		checker[element] = true
	}

	for _, element := range arrayB {
		checker[element] = true
	}

	for element := range checker {
		result = append(result,element)
	}

	return result
}

func main() {
	//test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	//{"king", "devil jin", "akuma", "eddie", "steve", "geese"}

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