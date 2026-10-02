package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(in, &n)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a)
		sums[i] = a + sums[i-1]
	}
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		var l, u int
		fmt.Fscan(in, &l, &u)
		fmt.Println(sums[u+1] - sums[l])
	}
}
