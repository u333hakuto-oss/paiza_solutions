package main

import "fmt"

func main() {
	var h, w, d int
	fmt.Scan(&h, &w, &d)
	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}
	var y, x int
	switch d {
	case 1:
		num := 1
		for diag := 0; diag < h+w-1; diag++ {
			firstY := diag
			if h-1 < firstY {
				firstY = h - 1
			}
			for y := firstY; y >= 0; y-- {
				x := diag - y
				if x < 0 || x >= w {
					continue
				}
				grid[y][x] = num
				num++
			}
		}
	case 2:
		r := 0
		for i := 1; i <= h*w; i++ {
			grid[y][x] = i
			if x == w-1 {
				r++
				y = r
				x = 0
			} else {
				x++
			}
		}
	case 3:
		c := 0
		for i := 1; i <= h*w; i++ {
			grid[y][x] = i
			if y == h-1 {
				c++
				y = 0
				x = c
			} else {
				y++
			}
		}
	case 4:
		num := 1
		for diag := 0; diag < h+w-1; diag++ {
			lastY := diag
			if h-1 < lastY {
				lastY = h - 1
			}
			for y := 0; y <= lastY; y++ {
				x := diag - y
				if x < 0 || x >= w {
					continue
				}
				grid[y][x] = num
				num++
			}
		}
	}
	for _, row := range grid {
		for i, num := range row {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}
		fmt.Println()
	}
}
