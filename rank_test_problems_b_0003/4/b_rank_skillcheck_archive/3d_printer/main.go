package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var x, y, z int
	fmt.Fscan(buf, &x, &y, &z)
	grid := make([][]byte, z)
	for i := range grid {
		grid[i] = make([]byte, y)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for d := 0; d < z; d++ {
		for h := 0; h < x; h++ {
			var s string
			fmt.Fscan(buf, &s)
			for w := 0; w < y; w++ {
				if s[w] == '#' {
					grid[d][w] = '#'
				}
			}
		}
		var dummy string
		fmt.Fscan(buf, &dummy)
	}
	for i := z - 1; i >= 0; i-- {
		fmt.Println(string(grid[i]))
	}
}
