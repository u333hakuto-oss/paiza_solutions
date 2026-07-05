package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, x, k int
	fmt.Fscan(buf, &n, &x, &k)
	ans := (k/4 - n) * x * 2
	if k%4 == 3 {
		ans += x
	}
	fmt.Println(ans)
}
