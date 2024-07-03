package main

import (
	"fmt"
)

func fibonacciSequence(n int) []int {
	fibonacci := []int{0, 1}

	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	return fibonacci
}

func main() {
	var n int
	fmt.Print("Введите элемет: ")
	fmt.Scanln(&n)

	fmt.Print("Последовательность Фибоначчи: ")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacciSequence(n)[i], " ")
	}

	fmt.Println()
}
