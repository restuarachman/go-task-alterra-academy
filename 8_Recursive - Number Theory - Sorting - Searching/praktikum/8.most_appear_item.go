package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	maping := make(map[string]int)

	insertion_sort := func(data []pair) []pair {
		for i := 0; i < len(data)-1; i++ {
			idxmin := i
			for j := i + 1; j < len(data); j++ {
				if data[j].count < data[idxmin].count {
					idxmin = j
				}
			}

			data[idxmin], data[i] = data[i], data[idxmin]
		}
		return data
	}

	for _, val := range items {
		maping[val]++
	}

	var p []pair
	var pairtemp pair
	for key, value := range maping {
		pairtemp.name = key
		pairtemp.count = value
		p = append(p, pairtemp)
	}
	p = insertion_sort(p)
	return p
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
