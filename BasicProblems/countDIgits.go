package main

import (
	"fmt"
)

func main() {
	Digits := countDigits(1234)
	fmt.Println(Digits)
}

func countDigits(n int) int {

	count := 0

	for n > 0 {
		n = n / 10
		count++
	}

	return count
}
