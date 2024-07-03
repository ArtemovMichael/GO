package main

import "fmt"

func findNuberOfOccurrences(slice []int, elem int) int {
	cnt := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == elem {
			cnt++
		}
	}
	return cnt
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
	fmt.Print("Введите элемент, для которого нужно найти количество вхождений: ")
	fmt.Scanln(&elem)

	fmt.Println("Количество вхождений элемента:", findNuberOfOccurrences(arr, elem))
}
