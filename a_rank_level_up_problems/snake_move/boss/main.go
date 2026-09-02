package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var x, y, n int
	fmt.Fscan(in, &x, &y, &n)
	moveX := []int{0, 1, 0, -1}
	moveY := []int{-1, 0, 1, 0}
	dir := 0
	for i := 0; i < n; i++ {
		var d string
		fmt.Fscan(in, &d)
		if d == "L" {
			dir += 3
		} else {
			dir += 1
		}
		dir %= 4
		x += moveX[dir]
		y += moveY[dir]
		fmt.Println(x, y)
	}
}
