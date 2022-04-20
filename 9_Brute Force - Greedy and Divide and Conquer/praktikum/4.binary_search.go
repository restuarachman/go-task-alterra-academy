package main

import "fmt"

func BinarySearch(array []int, x int) {
	left := 0
	right := len(array) - 1
	mid := (left + right) / 2
	for left < right && array[mid] != x {

		if x < array[mid] {
			right = mid - 1
		} else if x > array[mid] {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	if array[mid] == x {
		fmt.Println(mid)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)
}
