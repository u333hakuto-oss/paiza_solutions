package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	moveY := map[string]int{
		"N": -1,
		"S": 1,
		"E": 0,
		"W": 0,
	}
	moveX := map[string]int{
		"N": 0,
		"S": 0,
		"E": 1,
		"W": -1,
	}
	in := bufio.NewReader(os.Stdin)
	var h, w, sy, sx int
	var m string
	fmt.Fscan(in, &h, &w, &sy, &sx, &m)
	ny := sy + moveY[m]
	nx := sx + moveX[m]
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
