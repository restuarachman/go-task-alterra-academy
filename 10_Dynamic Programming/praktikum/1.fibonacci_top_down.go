package main

import "fmt"

var mapp map[int]int = map[int]int{}

func fibo(n int) int {

	if n <= 1 {
		mapp[n] = n
	} else if _, exist := mapp[n]; !exist {
		mapp[n] = fibo(n-1) + fibo(n-2)
	}
	return mapp[n]
}

func main() {

	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}
