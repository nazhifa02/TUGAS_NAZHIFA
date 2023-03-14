package main

import (
	"fmt"
)

func main() {
	m := make(chan int) //buat chanel untuk mengirim bilangan kelipatan 3

	go kelipatan3(m) // jalankan go nilai routine untuk mengirim nilai ke chanel

	for n := range m { // baca nilai dari chanel dan cetak satu per satu
		fmt.Println(n)
	}
}

func kelipatan3(m chan int) {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			m <- i // kiris bilangan kelipatan 3 ke chanel
		}
	}
	close(m) // tutup chanel setelah selesai mengirim nilai
}
