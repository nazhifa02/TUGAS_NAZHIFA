package main

import "fmt"

func main() {
	segitiga()
}
func segitiga() {
	for x := 1; x <= 6; x++ {
		for y := 6; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
