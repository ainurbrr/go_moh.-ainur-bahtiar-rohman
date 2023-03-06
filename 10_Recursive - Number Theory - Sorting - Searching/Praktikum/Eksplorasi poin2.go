package main

import (
	"fmt"
	"sort"
)

// urutkan barang telebih dahulu lalu hitung banyak barang yang bisa dibeli
func MaximumBuyProduct(money int, productPrice []int) {
	countStuff := 0
	moneyLeft := money
	sort.Ints(productPrice)
	for _, price := range productPrice {
		if moneyLeft >= price {
			countStuff += 1
			moneyLeft -= price
		}
	}
	fmt.Println(countStuff)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}
