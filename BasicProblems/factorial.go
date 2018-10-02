package main

import (
	"fmt"
)

func main() {

	//result int
	result := fact(0)
	fmt.Println(result)

}

func fact(n int) int {

	if n == 0 {
		return 1
	}
	return n * fact(n-1)

}
