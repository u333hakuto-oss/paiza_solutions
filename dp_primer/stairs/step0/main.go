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
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i >= 1 {
			dp[i] += dp[i-1]
		}
		if i >= 2 {
			dp[i] += dp[i-2]
		}
	}
	fmt.Println(dp[n])
}
