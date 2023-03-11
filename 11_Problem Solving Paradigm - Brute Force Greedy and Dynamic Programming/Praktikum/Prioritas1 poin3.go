package main

import "fmt"

func fibo(n int, listFib map[int]int) int {
	if n < 2 {
		return n
	}
	if val, ok := listFib[n]; ok {
		return val
	}
	listFib[n] = fibo(n-1, listFib) + fibo(n-2, listFib)
	return listFib[n]

}

func main() {
	list := make(map[int]int)
	fmt.Println(fibo(0, list))  //0
	fmt.Println(fibo(1, list))  //1
	fmt.Println(fibo(2, list))  //1
	fmt.Println(fibo(3, list))  //2
	fmt.Println(fibo(5, list))  //5
	fmt.Println(fibo(6, list))  //8
	fmt.Println(fibo(7, list))  //13
	fmt.Println(fibo(9, list))  //13
	fmt.Println(fibo(10, list)) //55
}
