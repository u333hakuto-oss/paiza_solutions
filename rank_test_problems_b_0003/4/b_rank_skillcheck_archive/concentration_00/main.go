package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var h, w, n, l int
	fmt.Fscan(buf, &h, &w, &n)
	cards := make([][]int, h)
	scores := make([]int, n)
	for i := range cards {
		cards[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(buf, &cards[i][j])
		}
	}
	player := 0
	fmt.Fscan(buf, &l)
	for i := 0; i < l; i++ {
		var a1, b1, a2, b2 int
		fmt.Fscan(buf, &a1, &b1, &a2, &b2)
		a1--
		b1--
		a2--
		b2--
		if cards[a1][b1] == cards[a2][b2] {
			scores[player] += 2
		} else {
			player = (player + 1) % n
		}
	}
	for _, score := range scores {
		fmt.Println(score)
	}
}
