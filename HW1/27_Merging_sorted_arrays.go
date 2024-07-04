package main

import (
	"fmt"
)

func mergeArrays(arr1, arr2 []int) []int {
	var answer []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			answer = append(answer, arr1[i])
			i++
		} else {
			answer = append(answer, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		answer = append(answer, arr1[i])
		i++
	}

	for j < len(arr2) {
		answer = append(answer, arr2[j])
		j++
	}

	return answer
}

func main() {
	var n int
	fmt.Print("Введите количество элементов первого массива: ")
	fmt.Scanln(&n)

	arr1 := make([]int, n)
	fmt.Println("Введите элементы первого массива: ")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr1[i])
	}

	var m int
	fmt.Print("Введите количество элементов второго массива: ")
	fmt.Scanln(&m)

	arr2 := make([]int, m)
	fmt.Println("Введите элементы второго массива: ")
	for i := 0; i < m; i++ {
		fmt.Scanln(&arr2[i])
	}

	fmt.Print("Объединение массивов: ")
	answer := mergeArrays(arr1, arr2)
	for i := 0; i < len(answer); i++ {
		fmt.Print(answer[i], " ")
	}

	fmt.Println()
}
