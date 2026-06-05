package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, diff, a, b int
	fmt.Fscan(buf, &n)
	list := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var m int
		fmt.Fscan(buf, &m)
		list[i] = m
	}
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			currentdiff := list[j] - list[i]
			if currentdiff > diff {
				diff = currentdiff
				a, b = i, j
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
