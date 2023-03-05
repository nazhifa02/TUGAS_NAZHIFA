package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {

}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang ->1 ruby->2 js->4
	fmt.Println(MostAppearItem([]string{"A", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// c->1 D->2 B->3 A->4

	fmt.Println(MostAppearItem([]string{" football", "basketball", "tenis"}))
	// footbal ->1 basketball->1 tenis->1
}
