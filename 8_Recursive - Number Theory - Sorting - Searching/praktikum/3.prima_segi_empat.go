package main

import (
	"fmt"
)

func primaSegiEmpat(high, wide, start int) {
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
	start++
	sum := 0
	for i := 0; i < wide; i++ {
		j := 0
		for j < high {
			if isPrime(start) {
				sum += start
				fmt.Print(start, " ")
				j++
			}
			start++
		}
		fmt.Println()
	}
	fmt.Print(sum, "\n")
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
