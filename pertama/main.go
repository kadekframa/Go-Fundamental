package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hello World, Belajar Golang!")

	penjumlahan := calculation.Add(10, 24)
	perkalian := calculation.Multiply(4, 6)
	pengurangan := calculation.Pengurangan(20, 2)
	pembagian := calculation.Pembagian(24, 4)

	fmt.Println("Hasil penjumlahan : ", penjumlahan)
	fmt.Println("Hasil perkalian : ", perkalian)
	fmt.Println("Hasil pengurangan : ", pengurangan)
	fmt.Println("Hasil pembagian : ", pembagian)
}
