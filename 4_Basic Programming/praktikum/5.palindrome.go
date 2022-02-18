package main

import "fmt"

func palindrome(input string) bool {
	n := len(input)
	for i := 0; i < n/2; i++ {
		if input[i] != input[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
