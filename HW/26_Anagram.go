package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	chars1 := strings.Split(str1, "")
	chars2 := strings.Split(str2, "")

	sort.Strings(chars1)
	sort.Strings(chars2)

	str1 = strings.Join(chars1, "")
	str2 = strings.Join(chars2, "")

	return str1 == str2
}

func main() {
	var str1, str2 string
	fmt.Print("Введите первую строку: ")
	fmt.Scanln(&str1)
	fmt.Print("Введите вторую строку: ")
	fmt.Scanln(&str2)

	if anagrams(str1, str2) {
		fmt.Println("Строки являются анаграммами.")
	} else {
		fmt.Println("Строки не являются анаграммами.")
	}
}
