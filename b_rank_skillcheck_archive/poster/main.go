package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, a, b, count int
	var s string
	fmt.Fscan(buf, &n, &a, &s, &b)
	putUp := make([]bool, n)
	cur := a - 1
	putUp[a-1] = true
	ans := []int{a}
	add := 1
	if s == "CCW" {
		add = -1
	}
	for len(ans) < n {
		cur = (cur + add + n) % n
		if !putUp[cur] {
			count++
			if count == b {
				putUp[cur] = true
				ans = append(ans, cur+1)
				count = 0
			}
		}
	}
	for i, val := range ans {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
