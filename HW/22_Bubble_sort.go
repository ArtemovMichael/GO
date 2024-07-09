package main

import (
	"fmt"
)

func bubleSorted(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
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

	fmt.Println("Отсортированный массив: ")
	for i := 0; i < len(bubleSorted(arr)); i++ {
		fmt.Print(bubleSorted(arr)[i], " ")
	}
}
