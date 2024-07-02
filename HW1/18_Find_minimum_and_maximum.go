package main

import (
	"fmt"
)

func findMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

func main() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива: ")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	min, max := findMinMax(arr)
	fmt.Println("Минимальный элемент:", min)
	fmt.Println("Максимальный элемент:", max)
}
