package main

import "fmt"

func intersection(arr1, arr2 []int) []int {
	used := make(map[int]int)
	var answer []int

	for i := 0; i < len(arr1); i++ {
		used[arr1[i]]++
	}

	for j := 0; j < len(arr2); j++ {
		if used[arr2[j]] > 0 {
			answer = append(answer, arr2[j])
			used[arr2[j]]--
		}
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

	fmt.Print("Пересечение массивов: ")
	answer := intersection(arr1, arr2)
	for i := 0; i < len(answer); i++ {
		fmt.Print(answer[i], " ")
	}

	fmt.Println()
}
