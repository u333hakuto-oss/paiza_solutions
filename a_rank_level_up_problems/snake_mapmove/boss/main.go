package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w, sy, sx, n int
	fmt.Fscan(in, &h, &w, &sy, &sx, &n)
	moveY := []int{-1, 0, 1, 0}
	moveX := []int{0, 1, 0, -1}
	grid := make([][]byte, h)
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
	}
	grid[sy][sx] = '*'
	dir := 0
	timeline := make([]int, 100)
	for i := 0; i < n; i++ {
		var t int
		var d string
		fmt.Fscan(in, &t, &d)
		if d == "L" {
			timeline[t] += 3
		} else {
			timeline[t]++
		}
	}
	for _, d := range timeline {
		dir = (dir + d) % 4
		ny := sy + moveY[dir]
		nx := sx + moveX[dir]
		if ny < 0 || ny >= h || nx < 0 || nx >= w {
			break
		}
		if grid[ny][nx] != '.' {
			break
		}
		grid[ny][nx] = '*'
		sy, sx = ny, nx
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
