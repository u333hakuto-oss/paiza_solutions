package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bomb struct {
	Y, X, Power int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(buf, &h, &w)
	var bombs []Bomb
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(buf, &s)
		for j := 0; j < len(s); j++ {
			if s[j] != '.' {
				p := int(s[j] - '0')
				bombs = append(bombs, Bomb{i, j, p})
			}
		}
	}
	board := make([][]byte, h)
	for i := range board {
		board[i] = make([]byte, w)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	for _, b := range bombs {
		p := b.Power
		sy := b.Y - p
		if sy < 0 {
			sy = 0
		}
		sx := b.X - p
		if sx < 0 {
			sx = 0
		}
		ey := b.Y + p
		if ey >= h {
			ey = h - 1
		}
		ex := b.X + p
		if ex >= w {
			ex = w - 1
		}
		for r := sy; r <= ey; r++ {
			for c := sx; c <= ex; c++ {
				diffY := b.Y - r
				if diffY < 0 {
					diffY *= -1
				}
				diffX := b.X - c
				if diffX < 0 {
					diffX *= -1
				}
				if diffY+diffX <= p {
					board[r][c] = '#'
				}
			}
		}
	}
	for _, row := range board {
		fmt.Println(string(row))
	}
}
