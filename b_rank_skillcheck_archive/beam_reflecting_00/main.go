// 　ビームが無限ループしない想定

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var h, w int
	fmt.Fscan(buf, &h, &w)
	grid := make([][]byte, h)
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(buf, &s)
		grid[i] = []byte(s)
	}
	const (
		Right = iota
		Down
		Left
		Up
	)
	y, x, dir := 0, -1, Right
	count := 0
	dy := []int{0, 1, 0, -1}
	dx := []int{1, 0, -1, 0}
	for {
		y += dy[dir]
		x += dx[dir]
		if y < 0 || y >= h || x < 0 || x >= w {
			break
		}
		switch grid[y][x] {
		case '/':
			dir = 3 - dir
		case '\\':
			if dir%2 == 0 {
				dir++
			} else {
				dir--
			}
		}
		count++
	}
	fmt.Println(count)
}
