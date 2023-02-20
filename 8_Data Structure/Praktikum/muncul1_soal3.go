package main

import (
	"fmt"
	"strings"
	"strconv"
)
func munculSekali(angka string) []int{
	pisahStr := strings.SplitAfter(angka, "")

	var result []int
	for i, element := range pisahStr{
		jadiInt, _ := strconv.Atoi(element)
		if element != pisahStr[i]{
			result = append(result, jadiInt)
		}
	} 
	
	return result
}

func main() {
	fmt.Println(munculSekali("42312218"))
}	