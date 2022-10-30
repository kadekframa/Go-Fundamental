package main

import "fmt"

func main() {
	score := 92
	var grade string

	fmt.Println(score)

	if score > 95 {
		grade = "A"
	} else if score > 90 {
		grade = "B"
	} else if score > 80 {
		grade = "C"
	} else if score > 70 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Println(grade)
}
