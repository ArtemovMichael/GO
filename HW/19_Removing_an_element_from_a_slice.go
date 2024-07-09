package main

import "fmt"

func removingByIndex(slice []int, index int) []int {
	index--
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
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

	var index int
	fmt.Print("Введите индекс элемента, который нужно удалить: ")
	fmt.Scanln(&index)

	newArr := removingByIndex(arr, index)
	fmt.Println("Массив после удаления элемента: ")
	for i := 0; i < len(newArr); i++ {
		fmt.Print(newArr[i], " ")
	}
}
