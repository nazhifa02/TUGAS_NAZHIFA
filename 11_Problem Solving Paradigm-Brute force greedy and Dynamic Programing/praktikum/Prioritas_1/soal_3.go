package main

import "fmt"

func fibo(n int) int {
	if n < 0 {
		return -1
	}
	if n <= 1 {
		return n
	}
	numA, numB := 0, 1
	for i := 2; i <= n; i++ {
		numAB := numA + numB
		numA = numB
		numB = numAB
	}
	return numB
}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}
