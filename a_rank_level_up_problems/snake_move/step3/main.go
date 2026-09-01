package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var y, x int
	var d1, d2 string
	fmt.Fscan(buf, &y, &x, &d1, &d2)
	leftMoveY := map[string]int{
		"N": 0,
		"S": 0,
		"W": 1,
		"E": -1,
	}
	leftMoveX := map[string]int{
		"N": -1,
		"S": 1,
		"W": 0,
		"E": 0,
	}
	switch d2 {
	case "L":
		fmt.Println(y+leftMoveY[d1], x+leftMoveX[d1])
	case "R":
		fmt.Println(y-leftMoveY[d1], x-leftMoveX[d1])
	}
}
