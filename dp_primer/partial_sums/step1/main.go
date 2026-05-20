package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, x int
	fmt.Fscan(buf, &n, &x)
	weights := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(buf, &weights[i])
	}
	dp := make([]int, x+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := x; j >= weights[i]; j-- {
			dp[j] = (dp[j] + dp[j-weights[i]]) % 1_000_000_007
		}
	}
	fmt.Println(dp[x])
}
