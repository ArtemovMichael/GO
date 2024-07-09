package main

import "fmt"

func mean(arr []int) float64 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	var n int
	var arr []int
	fmt.Println("Введите количество элементов массива:")
	fmt.Scan(&n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		arr = append(arr, x)
	}

	fmt.Println("Среднее арифметическое элементов массива:", mean(arr))
}
