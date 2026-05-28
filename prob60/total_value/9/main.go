package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, sum int
	fmt.Fscan(buf, &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(buf, &a, &b)
		if a == b {
			sum += a * b
		} else {
			sum += a + b
		}
	}
	fmt.Println(sum)
}
