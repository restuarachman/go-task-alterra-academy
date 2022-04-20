package main

import "fmt"

func main() {
	// input
	var studentScore int = 80

	var score string

	switch {
	case studentScore <= 100 && studentScore >= 80:
		score = "A"
	case studentScore >= 65:
		score = "B"
	case studentScore >= 50:
		score = "C"
	case studentScore >= 35:
		score = "D"
	case studentScore >= 0:
		score = "E"
	default:
		score = "Nilai"
	}
	fmt.Println(score)
}
