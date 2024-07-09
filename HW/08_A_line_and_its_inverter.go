package main

import (
	"fmt"
)

func reverseString(input string) string {
	reverseString := ""
	for i := len(input) - 1; i >= 0; i-- {
		reverseString += string(input[i])
	}
	return reverseString

}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(reverseString(str))
}
