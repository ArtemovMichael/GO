package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := -1
	right := len(arr)

	for left < right-1 {
		mid := (left + right) / 2

		if arr[mid] < target {
			left = mid
		} else {
			right = mid
		}

	}
	return right
}

func main() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	fmt.Scanln(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива в возрастающем порядке: ")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}

	var elem int
	fmt.Print("Введите элемент для поиска: ")
	fmt.Scanln(&elem)

	index := binarySearch(arr, elem)
	if index == len(arr) || arr[index] != elem {
		fmt.Println("Элемент не найден.")
	} else {
		fmt.Println("Индекс элемента:", index+1)
	}
}
