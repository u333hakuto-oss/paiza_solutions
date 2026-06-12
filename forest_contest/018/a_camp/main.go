package main

import (
	"bufio"
	"fmt"
	"os"
)

func expand(grid [][]rune, queue *[][]int, player rune, h, w int) bool {
	if len(*queue) == 0 {
		return false
	}
	moved := false
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	length := len(*queue)
	for i := 0; i < length; i++ {
		curr := (*queue)[0]
		*queue = (*queue)[1:]
		r, c := curr[0], curr[1]
		for d := 0; d < 4; d++ {
			nr := r + dx[d]
			nc := c + dy[d]
			if nr >= 0 && nr < h && nc >= 0 && nc < w && grid[nr][nc] == '.' {
				grid[nr][nc] = player
				*queue = append(*queue, []int{nr, nc})
				moved = true
			}
		}
	}
	return moved
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var h, w int
	var n string
	fmt.Fscan(buf, &h, &w, &n)
	var queueA [][]int
	var queueB [][]int
	grid := make([][]rune, h)
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(buf, &s)
		grid[i] = []rune(s)
		for j := 0; j < w; j++ {
			if grid[i][j] == 'A' {
				queueA = append(queueA, []int{i, j})
			} else if grid[i][j] == 'B' {
				queueB = append(queueB, []int{i, j})
			}
		}
	}
	for {
		movedA, movedB := false, false
		if n == "A" {
			movedA = expand(grid, &queueA, 'A', h, w)
			movedB = expand(grid, &queueB, 'B', h, w)
		} else {
			movedB = expand(grid, &queueB, 'B', h, w)
			movedA = expand(grid, &queueA, 'A', h, w)
		}
		if !movedA && !movedB {
			break
		}
	}
	var countA, countB int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 'A' {
				countA++
			} else if grid[i][j] == 'B' {
				countB++
			}
		}
	}
	fmt.Println(countA, countB)
	if countA > countB {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
