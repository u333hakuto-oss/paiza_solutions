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
	size := n + 5
	dp := make([]int, size)
	// n ≦ 1,000かつa, b ≦ 10,000より金額は10,000,000を超えない
	const unreachable int = 10_000_001
	for i := range dp {
		dp[i] = unreachable
	}
	dp[0] = 0
	for i := 1; i < size; i++ {
		if i >= 2 && dp[i-2] != unreachable {
			dp[i] = dp[i-2] + a
		}
		if i-5 >= 0 && dp[i-5] != unreachable {
			if dp[i-5]+b < dp[i] {
				dp[i] = dp[i-5] + b
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
