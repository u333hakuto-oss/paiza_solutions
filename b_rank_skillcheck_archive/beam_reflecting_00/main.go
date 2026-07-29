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
	y, x := 0, -1
	dy, dx := 0, 1
	count := 0
	for {
		y += dy
		x += dx
		if y < 0 || y >= h || x < 0 || x >= w {
			break
		}
		switch grid[y][x] {
		case '/':
			dy, dx = -1*dx, -1*dy
		case '\\':
			dy, dx = dx, dy
		}
		count++
	}
	fmt.Println(count)
}
