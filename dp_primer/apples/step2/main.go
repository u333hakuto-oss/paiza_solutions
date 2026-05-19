package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, x, a, y, b int
	fmt.Fscan(buf, &n, &x, &a, &y, &b)
	size := n + y
	dp := make([]int, size)
	// n ≦ 1,000かつa, b ≦ 10,000より金額は10,000,000を超えない
	const unreachable int = 10_000_001
	for i := range dp {
		dp[i] = unreachable
	}
	dp[0] = 0
	for i := 1; i < size; i++ {
		if i >= x && dp[i-x] != unreachable {
			dp[i] = dp[i-x] + a
		}
		if i-y >= 0 && dp[i-y] != unreachable {
			if dp[i-y]+b < dp[i] {
				dp[i] = dp[i-y] + b
			}
		}
	}
	for i := size - 2; i >= 0; i-- {
		if dp[i+1] < dp[i] {
			dp[i] = dp[i+1]
		}
	}
	fmt.Println(dp[n])
}
