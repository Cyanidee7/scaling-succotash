package main

import (
	"fmt"
)

//Calculation of Potential Energy

func main() {

	var mass float32 = 8.0
	var g float32 = 9.8
	var height float32 = 2.0
	answer := mass * g * height

	fmt.Println("the amount of potential energy is", answer, "joule")
}
