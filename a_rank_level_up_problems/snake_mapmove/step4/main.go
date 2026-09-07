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
	grid := make([][]rune, h)
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []rune(s)
	}
	dir := 0
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
			if ny < 0 || ny >= h || nx < 0 || nx >= w {
				fmt.Println(sy, sx)
				fmt.Println("Stop")
				return
			}
			if grid[ny][nx] == '#' {
				fmt.Println(sy, sx)
				fmt.Println("Stop")
				return
			}
			sy, sx = ny, nx
		}
		fmt.Println(sy, sx)
	}
}
