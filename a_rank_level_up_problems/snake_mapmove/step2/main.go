package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w, sy, sx int
	var d, m string
	fmt.Fscan(in, &h, &w, &sy, &sx, &d, &m)
	dirIdx := map[string]int{
		"N": 0,
		"E": 1,
		"S": 2,
		"W": 3,
	}
	moveY := []int{-1, 0, 1, 0}
	moveX := []int{0, 1, 0, -1}
	dir := dirIdx[d]
	if m == "L" {
		dir += 3
	} else {
		dir++
	}
	dir %= 4
	ny := sy + moveY[dir]
	nx := sx + moveX[dir]
	if ny < 0 || ny >= h || nx < 0 || nx >= w {
		fmt.Println("No")
		return
	}
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(in, &s)
		if i == ny {
			if s[nx] == '#' {
				fmt.Println("No")
			} else {
				fmt.Println("Yes")
			}
			return
		}
	}
}
