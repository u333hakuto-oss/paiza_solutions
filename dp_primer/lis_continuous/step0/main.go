package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(buf, &n)
	size := n + 1
	heights := make([]int, size)
	for i := 1; i < size; i++ {
		var a int
		fmt.Fscan(buf, &a)
		heights[i] = a
	}
	dp := make([]int, size)
	dp[1] = 1
	max := 1
	for i := 2; i <= n; i++ {
		if heights[i-1] <= heights[i] {
			dp[i] = dp[i-1] + 1
			if dp[i] > max {
				max = dp[i]
			}
		} else {
			dp[i] = 1
		}
	}
	fmt.Println(max)
}
