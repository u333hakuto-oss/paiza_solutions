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
	dp := make([]bool, x+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		var w int
		fmt.Fscan(buf, &w)
		for j := x; j >= w; j-- {
			if dp[j-w] {
				dp[j] = true
			}
		}
	}
	for i := x; i >= 0; i-- {
		if dp[i] {
			fmt.Println(i)
			break
		}
	}
}
