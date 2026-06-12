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
	forest := make([][]rune, h)
	var ac, ar int
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(buf, &s)
		forest[i] = []rune(s)
		for j, r := range forest[i] {
			if r == 'A' {
				ac = i
				ar = j
			}
		}
	}
	for i := ac - 1; i >= 0; i-- {
		if forest[i][ar] == '.' {
			forest[i][ar] = 'V'
		} else {
			break
		}
	}
	for i := ac + 1; i < h; i++ {
		if forest[i][ar] == '.' {
			forest[i][ar] = 'V'
		} else {
			break
		}
	}
	for i := ar - 1; i >= 0; i-- {
		if forest[ac][i] == '.' {
			forest[ac][i] = 'V'
		} else {
			break
		}
	}
	for i := ar + 1; i < w; i++ {
		if forest[ac][i] == '.' {
			forest[ac][i] = 'V'
		} else {
			break
		}
	}
	for _, row := range forest {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}
