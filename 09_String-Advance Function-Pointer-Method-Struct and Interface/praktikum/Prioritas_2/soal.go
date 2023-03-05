package main

import "fmt"

func caesar(offset int, input string) string {
	var encrypted string
	for _, char := range input {
		if char != ' ' {
			charCode := int(char)
			shiftedCode := (charCode-97+offset)%26 + 97
			shiftedChar := string(shiftedCode)
			encrypted += shiftedChar
		} else {
			encrypted += " "
		}
	}

	return encrypted
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmnowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
