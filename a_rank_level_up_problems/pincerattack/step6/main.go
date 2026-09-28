package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w, y, x int
	fmt.Fscan(in, &h, &w)
	grid := make([][]byte, h)
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			if grid[i][j] == '!' {
				y, x = i, j
			}
		}
	}
	grid[y][x] = '*'
	dy := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dx := []int{0, 0, -1, 1, -1, 1, 1, -1}
	for d := 0; d < 8; d++ {
		ny := y
		nx := x
		found := false
		for {
			ny += dy[d]
			nx += dx[d]
			if ny < 0 || ny >= h || nx < 0 || nx >= w {
				break
			}
			if grid[ny][nx] == '#' {
				break
			}
			if grid[ny][nx] == '*' {
				found = true
				break
			}
		}
		if found {
			for {
				ny -= dy[d]
				nx -= dx[d]
				if ny == y && nx == x {
					break
				}
				grid[ny][nx] = '*'
			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
