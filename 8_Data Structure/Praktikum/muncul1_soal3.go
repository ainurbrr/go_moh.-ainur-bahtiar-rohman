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
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
