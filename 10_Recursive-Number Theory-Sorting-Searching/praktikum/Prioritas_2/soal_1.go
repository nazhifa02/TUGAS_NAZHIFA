package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		if card[0] == deck[0] || card[0] == deck[1] {
			return card
		}
		if card[1] == deck[0] || card[1] == deck[1] {
			return card
		}
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{3, 3}, []int{4, 3}}, []int{3, 4}))
	// [3 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [3 6]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}, []int{5, 1}}, []int{4, 6}))
	// tutup kartu
}
