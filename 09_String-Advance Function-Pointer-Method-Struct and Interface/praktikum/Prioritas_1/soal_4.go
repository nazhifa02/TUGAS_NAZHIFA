package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	// input
	for i, e := range numbers {
		switch {
		case i == 0:
			max, min = *e, *e
		case *e > max:
			max = *e
		case *e < min:
			min = *e
		}
	}

	return min, max
}
func main() {
	var a1, a2, a3, a4, a5, a6, min int
	fmt.Println("inputkan 6 angka:")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max := getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
