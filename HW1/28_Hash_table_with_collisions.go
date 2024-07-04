package main

import (
	"fmt"
)

type Node struct {
	key   string
	value string
}

type HashTable struct {
	table [][]*Node
	size  int
}

func MakeHashTable(size int) *HashTable {
	return &HashTable{
		table: make([][]*Node, size),
		size:  size,
	}
}

func (h *HashTable) Insert(key, value string) {
	index := hash(key, h.size)
	if h.table[index] == nil {
		h.table[index] = make([]*Node, 0)
	}
	h.table[index] = append(h.table[index], &Node{key, value})
}

func (h *HashTable) Get(key string) (string, bool) {
	index := hash(key, h.size)

	for _, node := range h.table[index] {
		if node.key == key {
			return node.value, true
		}
	}
	return "", false
}

func hash(key string, size int) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (31*hash + int(key[i])) % size
	}
	return hash
}

func main() {
	var size int
	fmt.Println("Введите размер хеш-таблицы: ")
	fmt.Scan(&size)

	hashTable := MakeHashTable(size)
	var n int
	fmt.Println("Введите количество элементов: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var key, value string
		fmt.Scan(&key)
		fmt.Scan(&value)
		hashTable.Insert(key, value)
	}

	var key string
	fmt.Println("Введите ключ для поиска: ")
	fmt.Scan(&key)

	value, ok := hashTable.Get(key)
	if ok {
		fmt.Println("Значение:", value)
	} else {
		fmt.Println("Элемент не найден")
	}
}
