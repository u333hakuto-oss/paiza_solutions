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
		dp[i] = 1
		for j := 1; j < i; j++ {
			if heights[j] > heights[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Println(max)
}
