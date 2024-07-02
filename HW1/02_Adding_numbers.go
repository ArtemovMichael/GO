package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	fmt.Println(add(a, b))
}
