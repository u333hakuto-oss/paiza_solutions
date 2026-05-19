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
	dp := make([]int, n+5)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		indexA := i - 2
		if indexA < 0 {
			indexA = 0
		}
		indexB := i - 5
		if indexB < 0 {
			indexB = 0
		}
		plusA := dp[indexA] + a
		plusB := dp[indexB] + b
		if plusA < plusB {
			dp[i] = plusA
		} else {
			dp[i] = plusB
		}
	}
	fmt.Println(dp[n])
}
