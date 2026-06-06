// Sliding Window
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, k, windowSum int
	fmt.Fscan(buf, &n, &k)
	a := make([]int, n)
	for i := 0; i < k; i++ {
		windowSum += a[i]
	}

	maxSum := windowSum

	for i := k; i < n; i++ {
		windowSum = windowSum + a[i] - a[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	fmt.Println(maxSum)
}
