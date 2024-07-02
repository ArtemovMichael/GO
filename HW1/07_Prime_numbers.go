package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) []bool {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for j := 2; j*j <= n; j++ {
		if primes[j] == true {
			for i := j * j; i <= n; i += j {
				primes[i] = false
			}
		}
	}

	return primes
}

func main() {
	var n int
	fmt.Scan(&n)

	primes := sieveOfEratosthenes(n)
	for i := 2; i <= n; i++ {
		if primes[i] {
			fmt.Println(i)
		}
	}
}
