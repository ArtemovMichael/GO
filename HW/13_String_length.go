package main

import "fmt"

func stringLength(str string) int {
	cnt := 0
	for range str {
		cnt++
	}
	return cnt
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(stringLength(str))
}
