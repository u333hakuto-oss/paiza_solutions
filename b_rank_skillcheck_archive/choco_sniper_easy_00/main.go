package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m, total int
	fmt.Fscan(buf, &n, &m)
	for i := 0; i < m; i++ {
		profit := 0
		for j := 0; j < n; j++ {
			var e int
			fmt.Fscan(buf, &e)
			profit += e
		}
		if profit > 0 {
			total += profit
		}
	}
	fmt.Println(total)
}
