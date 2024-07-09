package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var width, height float64
	fmt.Println("Введите ширину прямоугольника:")
	fmt.Scan(&width)
	fmt.Println("Введите высоту прямоугольника:")
	fmt.Scan(&height)
	rect := Rectangle{width, height}
	fmt.Println("Площадь прямоугольника:", rect.Area())
}
