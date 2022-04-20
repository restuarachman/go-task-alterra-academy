package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var res interface{}
	exist := func(card, deck []int) bool {
		return card[0] == deck[0] || card[0] == deck[1] || card[1] == deck[0] || card[1] == deck[1]
	}
	max := 0
	for _, val := range cards {
		if exist(val, deck) {
			if val[0]+val[1] > max {
				res = val
				max = val[0] + val[1]
			}
		}
	}
	switch res.(type) {
	case nil:
		return "tutup kartu"
	default:
		return res
	}

}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	//// [6 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	// “tutup kartu”
}
