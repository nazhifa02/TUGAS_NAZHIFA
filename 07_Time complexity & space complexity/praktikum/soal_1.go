package main

import (
	"fmt"
)

func primeNumber(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d Bukan Bilangan Prima\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d Bilangan Prima\n", number)
		}
	}
	return isPrime
}
func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false

}
