package main

import "fmt"

func main() {
	var name string = "Ruby on Rails" // Golang merupakan static language. Harus menerapkan tipe data pada deklarasi variable nya.

	age := 17          // Namun pada golang juga dapat mendeklarasikan variable nya langsung tanpa menerapkan tipe data pada variable baru.
	age = 20           // Apabila ingin mengubah value dari variable hanya menggunakan tanda sama dengan. (Tanda := hanya digunakan untuk membuat variable baru).
	var address string // Default value dari tipe data string adalah string kosong.

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(address) // Default value println.
}
