// Find all primes from 2 to given n
package main

import (
	"fmt"
)

func main() {
	sieve(20)
	return
}

func sieve(n int) {

	primeList := []bool{}

	for i := 0; i <= n; i++ {
		primeList = append(primeList, true)
	}

	//fmt.Println(primeList)
	p := 2

	for p*p <= n {
		if primeList[p] == true {
			for i := p * 2; i <= n+1; i = i + p {
				primeList[i] = false
			}
			p = p + 1
		}
	}

	// Print all primes
	for i := 2; i < n; i++ {
		if primeList[i] == true {
			fmt.Println(i)
		}
	}

}
