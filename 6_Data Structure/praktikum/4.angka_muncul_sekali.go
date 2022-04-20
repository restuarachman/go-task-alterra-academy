package main

import (
	"fmt"
	"strconv"
)

func angkaMunculSekali(angka string) []int {
	count := map[string]int{}
	for _, value := range angka {
		count[string(value)] += 1
	}

	result := []int{}
	for key, val := range count {
		if val == 1 {
			num, _ := strconv.Atoi(key)
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(angkaMunculSekali("1234123"))
	fmt.Println(angkaMunculSekali("76523752"))
	fmt.Println(angkaMunculSekali("1122"))
}
