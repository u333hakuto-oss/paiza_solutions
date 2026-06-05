package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, minidx, diff, a, b int
	min := 130 // m ≦ 130
	fmt.Fscan(buf, &n)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Fscan(buf, &m)
		if m < min {
			min = m
			minidx = i

		} else {
			currentdiff := m - min
			if currentdiff > diff {
				diff = currentdiff
				a, b = minidx, i
			}
		}
	}
	if a == 0 && b == 0 {
		fmt.Println("No")
	} else {
		fmt.Println(a)
		fmt.Println(b)
	}
}
