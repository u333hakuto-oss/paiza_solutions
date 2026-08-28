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
	l := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(buf, &l[i])
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if y > 0 && l[y-1][x] != '#' {
				continue
			}
			if y < h-1 && l[y+1][x] != '#' {
				continue
			}
			if x > 0 && l[y][x-1] != '#' {
				continue
			}
			if x < w-1 && l[y][x+1] != '#' {
				continue
			}
			fmt.Println(y, x)
		}
	}
}
