package main

import "fmt"

type Queue struct {
	first  []int
	second []int
}

func (q *Queue) update() {
	if len(q.second) == 0 {
		for len(q.first) > 0 {
			q.second = append(q.second, q.first[len(q.first)-1])
			q.first = q.first[:len(q.first)-1]
		}
	}
}

func (q *Queue) push(elem int) {
	q.first = append(q.first, elem)
}

func (q *Queue) front() int {
	q.update()

	if len(q.second) == 0 {
		return -1
	}

	return q.second[len(q.second)-1]
}

func (q *Queue) back() int {
	q.update()

	if len(q.first) == 0 && len(q.second) == 0 {
		return -1
	}

	if len(q.first) > 0 {
		return q.first[len(q.first)-1]
	}

	return q.second[0]
}

func (q *Queue) pop() int {
	q.update()

	if len(q.second) == 0 {
		return -1
	}

	value := q.second[len(q.second)-1]
	q.second = q.second[:len(q.second)-1]
	return value
}

func main() {
	q := Queue{}

	var n int
	fmt.Print("Введите количество запросов: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var query string
		fmt.Scanln(&query)

		switch query {
		case "push":
			var elem int
			fmt.Scanln(&elem)
			q.push(elem)
		case "pop":
			fmt.Println(q.pop())
		case "front":
			fmt.Println(q.front())
		case "back":
			fmt.Println(q.back())

		}
	}
}
