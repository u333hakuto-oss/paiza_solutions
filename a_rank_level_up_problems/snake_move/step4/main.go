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
	moveY := []int{0, 1, 0, -1}
	moveX := []int{1, 0, -1, 0}
	d := 0
	stepsInCurDir := 0
	segmentLength := 1
	segmentRepeat := 0
	for i := 0; i < n; i++ {
		if stepsInCurDir == segmentLength {
			d++
			if d == 4 {
				d = 0
			}
			segmentRepeat++
			stepsInCurDir = 0
		}
		if segmentRepeat == 2 {
			segmentLength++
			segmentRepeat = 0
		}
		x += moveX[d]
		y += moveY[d]
		stepsInCurDir++
	}
	fmt.Println(x, y)
}
