package calculation

// Penjumlahan.
func Add(number int, numberTwo int) int { // nama function "Add" menggunakan awalan huruf kapital agar bisa diakses secara public (bisa diakses dari package lain).
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}

// Perkalian.
func Multiply(number int, numberTwo int) int {
	return multiply(number, numberTwo)
}

func multiply(number int, numberTwo int) int {
	return number * numberTwo
}

// Pengurangan.
func Pengurangan(number int, numberTwo int) int {
	return pengurangan(number, numberTwo)
}

func pengurangan(number int, numberTwo int) int {
	return number - numberTwo
}

func Pembagian(number int, numberTwo int) int {
	return pembagian(number, numberTwo)
}

func pembagian(number int, numberTwo int) int {
	return number / numberTwo
}
