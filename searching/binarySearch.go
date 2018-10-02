package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 4, 5, 7, 92}
	val, index := binarySearch(arr, 1)

	if val == true {
		fmt.Println("Element found at index: ", index)
	} else {
		fmt.Println("Element NOT found")
	}

}

func binarySearch(list []int, x int) (bool, int) {

	low := 0
	high := len(list) - 1

	for low <= high {

		mid := (low + (high-low)/2)

		if list[mid] == x {
			return true, mid
		} else if list[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false, 0
}
