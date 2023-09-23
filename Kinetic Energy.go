package main

import (
	"fmt"
)

//Calculation of Kinetic Energy

func main() {

	var mass float32 = 8.0
	var velocity float32 = 4.0
	var half float32 = 0.5
	answer := half * mass * velocity * velocity

	fmt.Println("the amount of kinetic energy is", answer, "joule")
}
