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
	moveLX := []int{-1, 0, 1, 0}
	moveLY := []int{0, -1, 0, 1}
	dir := 0
	for i := 0; i < n; i++ {
		var d string
		fmt.Fscan(in, &d)
		if d == "L" {
			x += moveLX[dir]
			y += moveLY[dir]
			dir--
		} else {
			x -= moveLX[dir]
			y -= moveLY[dir]
			dir++
		}
		fmt.Println(x, y)
		dir = (dir + 4) % 4
	}
}
