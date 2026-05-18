package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, a, b, c int
	fmt.Fscan(buf, &n, &a, &b, &c)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= a {
			dp[i] += dp[i-a]
		}
		if i >= b {
			dp[i] += dp[i-b]
		}
		if i >= c {
			dp[i] += dp[i-c]
		}
	}
	fmt.Println(dp[n])
}
