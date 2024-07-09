package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Println("Введите строку:")
	fmt.Scanln(&str)

	if isPalindrome(str) {
		fmt.Println("Строка является палиндромом")
	} else {
		fmt.Println("Строка не является палиндромом")
	}
}
