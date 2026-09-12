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
	dir := 0
	timeline := make([]string, 100)
	for i := 0; i < n; i++ {
		var t int
		var d string
		fmt.Fscan(in, &t, &d)
		timeline[t] = d
	}
	for _, d := range timeline {
		switch d {
		case "L":
			dir += 3
		case "R":
			dir++
		}
		dir %= 4
		sy += moveY[dir]
		sx += moveX[dir]
		if sy < 0 || sy >= h || sx < 0 || sx >= w || grid[sy][sx] == '#' {
			fmt.Println("Stop")
			return
		} else {
			fmt.Println(sy, sx)
		}
	}
}
