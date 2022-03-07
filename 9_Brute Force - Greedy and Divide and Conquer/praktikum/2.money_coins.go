package main

import "fmt"

func moneyCoins(money int) []int {
	coins := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	i := 0
	result := []int{}
	for money > 0 {
		if money-coins[i] >= 0 {
			result = append(result, coins[i])
			money -= coins[i]
		} else {
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(moneyCoins(123))
	fmt.Println(moneyCoins(432))
	fmt.Println(moneyCoins(543))
	fmt.Println(moneyCoins(7752))
	fmt.Println(moneyCoins(15321))
}
