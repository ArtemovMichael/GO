package main

import (
	"fmt"
)

func removeDuplicates(nums []int) []int {
	visited := make(map[int]bool)
	unique := []int{}

	for i := 0; i < len(nums); i++ {
		if visited[nums[i]] == false {
			visited[nums[i]] = true
			unique = append(unique, nums[i])
		}
	}

	return unique
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

	fmt.Println("Массив без дубликатов: ")
	for i := 0; i < len(removeDuplicates(arr)); i++ {
		fmt.Print(removeDuplicates(arr)[i], " ")
	}
}
