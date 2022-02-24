package main

import "fmt"

func primeNumber(number int) string {
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return "Bukan Bilangan Prima"
		}
	}
	return "Bilangan Prima"
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
