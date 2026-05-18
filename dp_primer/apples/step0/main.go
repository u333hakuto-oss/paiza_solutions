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
	dp[0], dp[1] = 0, a
	for i := 2; i <= n; i++ {
		plusA := dp[i-1] + a
		plusB := dp[i-2] + b
		if plusA < plusB {
			dp[i] = plusA
		} else {
			dp[i] = plusB
		}
	}
	fmt.Println(dp[n])
}
