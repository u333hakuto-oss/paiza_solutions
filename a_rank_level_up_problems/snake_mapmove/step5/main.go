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
Loop:
	for i := 0; i < n; i++ {
		var d string
		var l int
		fmt.Fscan(in, &d, &l)
		if d == "L" {
			dir += 3
		} else {
			dir++
		}
		dir %= 4
		for j := 0; j < l; j++ {
			ny := sy + moveY[dir]
			nx := sx + moveX[dir]
			if ny < 0 || ny >= h || nx < 0 || nx >= w || grid[ny][nx] == '#' {
				break Loop
			}
			grid[ny][nx] = '*'
			sy, sx = ny, nx
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
