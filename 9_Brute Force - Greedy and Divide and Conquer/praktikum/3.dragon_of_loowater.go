package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	var (
		idxDragon  = 0
		max        = 0
		knightKill = 0
	)
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("knight fall")
		return
	}
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	for _, val := range knightHeight {
		if dragonHead[idxDragon] <= val {
			max += val

			knightKill += 1
			idxDragon += 1
		}
		if idxDragon == len(dragonHead) {
			fmt.Println(max)
			return
		}
	}
	fmt.Println("knight fall")
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11
	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
