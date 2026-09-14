package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(in, &h, &w)
	grid := make([][]byte, h)
	var queueY []int
	var queueX []int
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			if grid[i][j] == '*' {
				grid[i][j] = '0'
				queueY = append(queueY, i)
				queueX = append(queueX, j)
			}
		}
	}
	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}
	moveCount := 1
	start := 0
	end := 1
	for {
		change := 0
		for i := start; i < end; i++ {
			for d := 0; d < 4; d++ {
				ny := queueY[i] + dy[d]
				nx := queueX[i] + dx[d]
				if ny < 0 || ny >= h || nx < 0 || nx >= w {
					continue
				}
				if grid[ny][nx] == '.' {
					grid[ny][nx] = byte('0' + moveCount)
					queueY = append(queueY, ny)
					queueX = append(queueX, nx)
					change++
				}
			}
		}
		if change == 0 {
			break
		}
		start = end
		end += change
		moveCount++
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
