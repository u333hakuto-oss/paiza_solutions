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
	dp[0] = 0
	for i := 1; i <= x; i++ {
		dp[i] = 10001 // おもりの数n<=100
	}
	for i := 1; i <= n; i++ {
		for j := x; j >= weights[i]; j-- {
			if dp[j] > dp[j-weights[i]]+1 {
				dp[j] = dp[j-weights[i]] + 1
			}
		}
	}
	ans := dp[x]
	if ans == 10001 {
		ans = -1
	}
	fmt.Println(ans)
}
