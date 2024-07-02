package main

import "fmt"

func EvenOrOdd(num int) string {
	if num % 2 == 0 {
		return "even"
	} 
	return "odd"
	
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(EvenOrOdd(num))
}
