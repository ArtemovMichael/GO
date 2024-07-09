package main

import (
	"fmt"
	"strings"
)

func isVowel(char string) string {
	vowels := "aeiouyAEIOUYаеёиоуыэюяАЕЁИОУЫЭЮЯ"

	if strings.ContainsAny(vowels, char) {
		return "Vowel"
	}

	return "Consonant"
}

func main() {
	var char string
	fmt.Scan(&char)
	fmt.Println(isVowel(char))
}
