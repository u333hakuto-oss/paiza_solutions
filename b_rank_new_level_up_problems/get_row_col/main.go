package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(buf, &h, &w)
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(buf, &s)
		grid[i] = []rune(s)
	}
	var y, x int
	fmt.Fscan(buf, &y, &x)
	// 縦横
	for i, row := range grid {
		for j := range row {
			if i == y || j == x {
				if grid[i][j] == '#' {
					grid[i][j] = '.'
				} else {
					grid[i][j] = '#'
				}
			}
		}
	}
	// 斜め
	limit := h
	if w > limit {
		limit = w
	}
	for i := 1; i < limit; i++ {
		out := 0
		nys := []int{y + i, y + i, y - i, y - i}
		nxs := []int{x + i, x - i, x + i, x - i}
		for j := 0; j < 4; j++ {
			ny := nys[j]
			nx := nxs[j]
			if ny >= 0 && ny < h && nx >= 0 && nx < w {
				if grid[ny][nx] == '#' {
					grid[ny][nx] = '.'
				} else {
					grid[ny][nx] = '#'
				}
			} else {
				out++
			}
		}
		if out == 4 {
			break
		}
	}
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}
