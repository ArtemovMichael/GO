package main

import "fmt"

func findElem(slice []int, elem int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == elem {
			return i + 1
		}
	}
	return -1
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

	var elem int
	fmt.Print("Введите элемент, который нужно найти: ")
	fmt.Scanln(&elem)

	fmt.Println("Индекс элемента:", findElem(arr, elem))
}
