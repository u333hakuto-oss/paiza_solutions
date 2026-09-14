package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Y, X int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w, n int
	fmt.Fscan(in, &h, &w, &n)
	grid := make([][]byte, h)
	var queue []Point
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			if grid[i][j] == '*' {
				queue = append(queue, Point{i, j})
			}
		}
	}
	isQuestion := make([]bool, 101)
	for i := 0; i < n; i++ {
		var l int
		fmt.Fscan(in, &l)
		isQuestion[l] = true
	}
	if isQuestion[0] {
		grid[queue[0].Y][queue[0].X] = '?'
	}
	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}
	moveCount := 1
	start := 0
	end := 1
	for start < end {
		for i := start; i < end; i++ {
			for d := 0; d < 4; d++ {
				ny := queue[i].Y + dy[d]
				nx := queue[i].X + dx[d]
				if ny < 0 || ny >= h || nx < 0 || nx >= w {
					continue
				}
				if grid[ny][nx] == '.' {
					if isQuestion[moveCount] {
						grid[ny][nx] = '?'
					} else {
						grid[ny][nx] = '*'
					}
					queue = append(queue, Point{ny, nx})
				}
			}
		}
		start = end
		end = len(queue)
		moveCount++
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
