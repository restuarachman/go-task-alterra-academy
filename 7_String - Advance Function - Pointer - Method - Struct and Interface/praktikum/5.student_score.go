package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var sum int
	for _, val := range s.score {
		sum += val
	}
	return float64(sum) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min, name = s.score[0], s.name[0]
	for i := 0; i < len(s.name); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max, name = s.score[0], s.name[0]
	for i := 0; i < len(s.name); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
