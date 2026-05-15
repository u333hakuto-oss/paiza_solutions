package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var x, d1, d2, k int
	fmt.Fscan(buf, &x, &d1, &d2, &k)
	list := make([]int, k+1)
	list[0] = x
	for i := 1; i < k; i++ {
		if i%2 == 0 {
			list[i] = list[i-1] + d1
		} else {
			list[i] = list[i-1] + d2
		}
	}
	fmt.Println(list[k-1])
}
