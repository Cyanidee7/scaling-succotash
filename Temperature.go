package main

import (
	"fmt"
)

//Convertion of Temperature

func main() {

	var celcius float32 = 50.0
	var kelvin float32 = celcius + 273.15
	var fahrenheit float32 = celcius/1.8 + 32.0

	fmt.Println(celcius, "C°", kelvin, "K°", fahrenheit, "F°")
}
