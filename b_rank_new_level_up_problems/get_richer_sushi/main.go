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
	sushi := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &sushi[i])
	}
	var richer int
	for i := 0; i < k; i++ {
		richer += sushi[i]
	}
	windowSum := richer
	for i := 0; i < n; i++ {
		windowSum = windowSum - sushi[i] + sushi[(i+k)%n]
		if windowSum > richer {
			richer = windowSum
		}
	}
	fmt.Println(richer)
}
