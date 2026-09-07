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
		fmt.Fscan(in, &d)
		if d == "L" {
			dir += 3
		} else {
			dir++
		}
		dir %= 4
		sy += moveY[dir]
		sx += moveX[dir]
		if sy < 0 || sy >= h || sx < 0 || sx >= w {
			fmt.Println("Stop")
			return
		}
		if grid[sy][sx] == '#' {
			fmt.Println("Stop")
			return
		}
		fmt.Println(sy, sx)
	}
}
