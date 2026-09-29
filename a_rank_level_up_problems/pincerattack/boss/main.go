package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w, n, turn int
	fmt.Fscan(in, &h, &w, &n, &turn)
	grid := make([][]byte, h)
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
	}
	dy := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dx := []int{0, 0, -1, 1, -1, 1, 1, -1}
	for i := 0; i < turn; i++ {
		var p, y, x int
		fmt.Fscan(in, &p, &y, &x)
		player := byte('0' + p)
		grid[y][x] = player
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
				if grid[ny][nx] == player {
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
					grid[ny][nx] = player
				}
			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
