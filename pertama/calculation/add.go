package calculation

func Add(number int, numberTwo int) int { // nama function "Add" menggunakan awalan huruf kapital agar bisa diakses secara public (bisa diakses dari package lain).
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + numberTwo
}
