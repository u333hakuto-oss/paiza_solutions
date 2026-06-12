package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(buf, &n)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &prices[i])
	}
	fmt.Fscan(buf, &q)
	for i := 0; i < q; i++ {
		var l, r, count int
		fmt.Fscan(buf, &l, &r)
		for _, price := range prices {
			if price >= l && price <= r {
				count++
			}
		}
		fmt.Println(count)
	}
}
