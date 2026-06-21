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
	dy := []int{0, 1, -1, 0, 0}
	dx := []int{0, 0, 0, 1, -1}
	for i := 0; i < 5; i++ {
		ny, nx := y+dy[i], x+dx[i]
		if ny >= 0 && ny < h && nx >= 0 && nx < w {
			if grid[ny][nx] == '#' {
				grid[ny][nx] = '.'
			} else {
				grid[ny][nx] = '#'
			}
		}
	}

	for _, row := range grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}
