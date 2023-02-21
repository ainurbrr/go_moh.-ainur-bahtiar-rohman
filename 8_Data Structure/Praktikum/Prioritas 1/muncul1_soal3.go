package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {
	pisahStr := strings.SplitAfter(angka, "")

	checker := make(map[string]int)

	for _, element := range pisahStr {
		checker[element]++
	}

	var result []int
	for _, element := range pisahStr {
		if checker[element] == 1 {
			jadiInt, _ := strconv.Atoi(element)
			result = append(result, jadiInt)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123")) //[4]
	fmt.Println(munculSekali("76523752")) //[6 3]
	fmt.Println(munculSekali("12345")) //[1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) //[] 
	fmt.Println(munculSekali("0872504")) //[8 7 2 5 4]
}
