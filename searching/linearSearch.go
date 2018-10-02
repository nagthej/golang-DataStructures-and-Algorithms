package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 4, 5, 7, 92, 4, 5}
	val, index := linearSearch(arr, 92)

	if val == true {
		fmt.Println("Element found at index: ", index)
	} else {
		fmt.Println("Element NOT found")
	}

}

func linearSearch(list []int, x int) (bool, int) {

	for i := 0; i < len(list); i++ {
		if list[i] == x {
			return true, i
		}
	}

	return false, 0

}
