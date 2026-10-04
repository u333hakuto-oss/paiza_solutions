package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, sum, l int
	maxLen := 0
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for r := 0; r < n; r++ {
		sum += a[r]
		for sum > m {
			sum -= a[l]
			l++
		}
		curLen := r - l + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	fmt.Println(maxLen)
}
