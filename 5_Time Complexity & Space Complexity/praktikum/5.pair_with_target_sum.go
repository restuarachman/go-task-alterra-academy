package main

import "fmt"

func PairSum(arr []int, target int) []int {
	hash := map[int]int{}
	result := []int{}

	for i, value := range arr {
		if pos, exist := hash[target-value]; exist && i != pos {
			result = append(result, hash[target-value], i)
		} else {
			hash[value] = i
		}
	}
	return result
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
}
