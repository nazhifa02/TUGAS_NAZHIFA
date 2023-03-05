package main

import "fmt"

func main() {
	fmt.Print("masukkan bilangan")
	var num int
	fmt.Scanf("%d", &num)

	if (num % 2) == 0 {
		fmt.Printf("%d bilangan genap", num)
	} else {
		fmt.Printf("%d bilangan ganjil", num)
	}
}
