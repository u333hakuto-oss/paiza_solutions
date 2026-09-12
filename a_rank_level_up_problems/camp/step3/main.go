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
	var queue [][]int
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			if grid[i][j] == '*' {
				queue = append(queue, []int{i, j})
			}
		}
	}
	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			ny := cur[0] + dy[i]
			nx := cur[1] + dx[i]
			if ny < 0 || ny >= h || nx < 0 || nx >= w {
				continue
			}
			if grid[ny][nx] == '.' {
				grid[ny][nx] = '*'
				queue = append(queue, []int{ny, nx})
			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
