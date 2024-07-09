package main

import "fmt"

func find(arr []int, elem int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
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

	var elem int
	fmt.Println("Введите элемент, который хотите найти:")
	fmt.Scan(&elem)

	if find(arr, elem) {
		fmt.Println("Элемент найден")
	} else {
		fmt.Println("Элемент не найден")
	}
}
