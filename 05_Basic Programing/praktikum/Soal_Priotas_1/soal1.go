package main

import "fmt"

func main() {
	var luas, a, b, t float32
	fmt.Print("memasukkan tinggi segitiga:")
	ftm.Scanln(&t)
	ftm.Print("masukkan panjang sisi a:")
	ftm.Scanln(&a)
	ftm.Print(",masukkan panjang sisi b:")
	ftm.Scanln(&b)

	luas = 0.5 * t * (a + b)

	ftm.Println("luas trapesium adalah", luas)
}
