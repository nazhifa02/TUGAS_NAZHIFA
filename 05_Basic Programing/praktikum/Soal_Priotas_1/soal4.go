package main

import "fmt"

func main() {
	var nilai int = 100
	for i := 1; i <= nilai; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Print("fizz")
		} else if i%5 == 0 {
			fmt.Print("buzz")
		} else {
			fmt.Print(i)
		}
	}
}
