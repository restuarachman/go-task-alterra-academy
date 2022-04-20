package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var result string
	for i := 0; i < len(input); i++ {
		result += string((input[i]+byte(offset)-'a')%26 + 'a')
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
