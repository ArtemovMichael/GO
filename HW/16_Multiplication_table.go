package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число для вывода таблицы умножения: ")
	fmt.Scanln(&number)

	fmt.Println("Таблица умножения для числа", number, ":")
	for i := 1; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}
}
