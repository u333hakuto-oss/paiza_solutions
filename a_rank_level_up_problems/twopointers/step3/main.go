// 尺取り法
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, sum, l int
	minLen := -1
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	for r := 0; r < n; r++ {
		sum += a[r]
		for sum >= m {
			curLen := r - l + 1
			if minLen == -1 || curLen < minLen {
				minLen = curLen
			}
			sum -= a[l]
			l++
		}
	}
	fmt.Println(minLen)
}
