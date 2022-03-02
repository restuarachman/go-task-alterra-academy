package main

import (
	"fmt"
)

func primeX(number int) int {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {

				return false
			}
		}

		return true
	}

	i := 2
	res := 1
	count := 0
	for count < number {
		if isPrime(i) {
			res = i
			count++
		}
		i++
	}
	return res
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
