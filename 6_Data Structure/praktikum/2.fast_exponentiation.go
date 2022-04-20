package main

import "fmt"

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * x
		}

		n /= 2
		x *= x
	}
	return res
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
