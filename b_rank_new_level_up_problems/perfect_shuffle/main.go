package main

import (
	"fmt"
	"strconv"
)

func main() {
	const Deck = 52
	suits := []string{"S", "H", "D", "C"}
	cards := make([]string, Deck)
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			num := strconv.Itoa(j + 1)
			cards[i*13+j] = suits[i] + "_" + num
		}
	}
	var k int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		next := make([]string, Deck)
		for i := 0; i < Deck/2; i++ {
			next[2*i] = cards[i]
			next[2*i+1] = cards[i+26]
		}
		cards = next
	}
	for _, card := range cards {
		fmt.Println(card)
	}
}
