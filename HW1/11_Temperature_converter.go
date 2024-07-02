package main

import "fmt"

func CelsiusToFahrenheit(celsius float64) float64 {
	fahrenheit := (celsius * 9 / 5) + 32
	return fahrenheit
}

func main() {
	var celsius float64
	fmt.Print("Введите температуру в градусах Цельсия: ")
	fmt.Scan(&celsius)
	fahrenheit := CelsiusToFahrenheit(celsius)
	fmt.Printf("Температура в градусах Фаренгейта: %.1f\n", fahrenheit)
}
