package main

import (
	"fmt"
)

func main() {

	fmt.Println(euclidGCD(10, 2))

}

func euclidGCD(a int, b int) int {

	divident := 0
	divisor := 0
	rem := 0

	if a >= b {
		divident = a
		divisor = b
	} else {
		divident = b
		divisor = a
	}

	for divisor != 0 {
		rem = divident % divisor
		divident = divisor
		divisor = rem

	}

	return divident

}
