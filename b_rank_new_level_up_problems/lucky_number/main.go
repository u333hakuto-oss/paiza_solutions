package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(buf, &n)
	list := make([][]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(buf, &x)
		strX := strconv.Itoa(x)
		digits := len(strX)
		place := 1
		for j := 0; j < digits; j++ {
			place *= 10
		}
		list[i] = []int{x, x + 1, x - 1, x + place, x*10 + 1}
	}
	minDiff := 200000000
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for _, num1 := range list[i] {
				for _, num2 := range list[j] {
					diff := num1 - num2
					if diff < 0 {
						diff *= -1
					}
					if diff < minDiff {
						minDiff = diff
					}
				}
			}
		}
	}
	fmt.Println(minDiff)
}
