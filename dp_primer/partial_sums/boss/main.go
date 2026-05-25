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
	dp := make([]bool, x+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := weights[i]; j <= x; j++ {
			if dp[j-weights[i]] {
				dp[j] = true
			}
		}
	}
	if dp[x] {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
