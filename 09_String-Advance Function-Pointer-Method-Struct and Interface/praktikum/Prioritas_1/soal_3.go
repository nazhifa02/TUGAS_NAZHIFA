package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	res := strings.Compare(a, b)

	if res == 0 {
		return a
	} else if res < 0 {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
