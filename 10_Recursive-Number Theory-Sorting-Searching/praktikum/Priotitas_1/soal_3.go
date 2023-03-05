package main

import "fmt"

func getPrime(n int) int {
	var primes []int
	i := 2
	for len(primes) < n {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
		i++
	}
	return primes[n-1]
}

func main() {
	fmt.Println(getPrime(1))
	fmt.Println(getPrime(5))
	fmt.Println(getPrime(8))
	fmt.Println(getPrime(9))
	fmt.Println(getPrime(10))
}
