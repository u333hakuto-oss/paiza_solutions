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
		var a int
		fmt.Fscan(buf, &a)
		if a >= 5 {
			sum += a
		}
	}
	fmt.Println(sum)
}
