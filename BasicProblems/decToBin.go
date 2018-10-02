package main

import (
	"fmt"
)

func main() {
	fmt.Println(decToBin(20))
}

func decToBin(n int) []int {

	var result []int

	for n > 0 {
		rem := n % 2
		result = append(result, rem)
		n = n / 2
	}

	// Reverses result slice
	for i := len(result)/2 - 1; i >= 0; i-- {
		j := len(result) - 1 - i
		result[i], result[j] = result[j], result[i]
	}

	return result

}
