package main

import "fmt"

func SimpleEquations(a, b, c int) {
	var i, j, k = 1, 1, 1
	for a-(i+j+k) >= 0 || b-(i*j*k) >= 0 || c-(i*i+j*j+k*k) >= 0 {
		for a-(i+j+k) >= 0 || b-(i*j*k) >= 0 || c-(i*i+j*j+k*k) >= 0 {
			for a-(i+j+k) >= 0 || b-(i*j*k) >= 0 || c-(i*i+j*j+k*k) >= 0 {
				if a-(i+j+k) == 0 && b-(i*j*k) == 0 && c-(i*i+j*j+k*k) == 0 {
					fmt.Println(i, j, k)
					return
				}
				k++
			}
			j++
			k = 1
		}
		i++
		j = 1
	}
	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
