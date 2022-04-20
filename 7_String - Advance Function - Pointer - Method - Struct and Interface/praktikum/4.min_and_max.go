package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[1]
	for i := 0; i < len(numbers); i++ {
		if min > *numbers[i] {
			min = *numbers[i]
		}
		if max < *numbers[i] {
			max = *numbers[i]
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println(max, "is the maximum number")
	fmt.Println(min, "is the minimum number")
}
