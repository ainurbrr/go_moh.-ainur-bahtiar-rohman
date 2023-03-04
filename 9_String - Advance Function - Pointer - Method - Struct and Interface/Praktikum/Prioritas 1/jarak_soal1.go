package main

import (
	"fmt"
)

type Car struct {
	typeCar string
	fuelIn  float64
}

func (c Car) hitungJarak() float64 {
	var jarak float64

	const konsumsi = 1.5

	jarak = c.fuelIn / konsumsi
	fmt.Println("Jenis fuel yang digunakan adalah :", c.typeCar)
	fmt.Println("Banyak Fuel yang terisi adalah :", c.fuelIn, "L")
	return jarak
}

func main() {
	var car = Car{
		typeCar: "Sport Car",
		fuelIn:  100,
	}

	fmt.Printf("Jarak yang dapat ditempuh oleh mobil tersebut adalah : %.2f mil", car.hitungJarak())
}
