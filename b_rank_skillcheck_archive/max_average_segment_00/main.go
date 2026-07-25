package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(buf, &n, &k)
	accesses := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &accesses[i])
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += accesses[i]
	}
	maxSum := sum
	optionCount := 1
	firstOption := 1
	for i := k; i < n; i++ {
		sum = sum + accesses[i] - accesses[i-k]
		if sum > maxSum {
			maxSum = sum
			optionCount = 1
			firstOption = i - k + 2
		} else if sum == maxSum {
			optionCount++
		}
	}
	fmt.Println(optionCount, firstOption)
}
