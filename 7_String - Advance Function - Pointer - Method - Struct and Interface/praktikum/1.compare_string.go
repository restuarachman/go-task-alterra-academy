package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	if strings.Contains(b, a) {
		return a
	}
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
