package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var x, d, k int
	fmt.Fscan(buf, &x, &d, &k)
	list := make([]int, k)
	list[0] = x
	for i := 1; i < k; i++ {
		list[i] = list[i-1] + d
	}
	fmt.Println(list[k-1])
}
