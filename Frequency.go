package main

import (
	"fmt"
)

//Calculation of Frequency

func main() {

	var cycle float32 = 7.0
	var time float32 = 3.5
	answer := cycle / time

	fmt.Println("the frequency is", answer, "hertz")
}
