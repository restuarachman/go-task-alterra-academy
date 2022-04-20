package main

import "fmt"

type countMap map[string]int

func Frequency(s string) countMap {
	m := countMap{}
	for _, r := range s {
		m[string(r)]++
		fmt.Println(string(r), ":", m[string(r)])
	}
	return m
}

func letterCount(l string) countMap {
	channel := make(chan countMap, 2)
	wordSplit := []string{l[:len(l)/2], l[len(l)/2:]}
	for _, val := range wordSplit {
		go func(w string) {
			channel <- Frequency(w)
		}(val)
	}
	res := countMap{}
	for range wordSplit {
		for key, value := range <-channel {
			res[key] += value
		}
	}
	return res
}

func main() {
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	fmt.Println(letterCount(data))
}
