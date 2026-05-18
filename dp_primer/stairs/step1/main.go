package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, a, b int
	fmt.Fscan(buf, &n, &a, &b)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= a {
			dp[i] += dp[i-a]
		}
		if i >= b {
			dp[i] += dp[i-b]
		}
	}
	fmt.Println(dp[n])
}
