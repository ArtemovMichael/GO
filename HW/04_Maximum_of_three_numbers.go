package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxOfThreeNumbers(a, b, c int) int {
	return Max(Max(a, b), c)
}

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(MaxOfThreeNumbers(a, b, c))
}
