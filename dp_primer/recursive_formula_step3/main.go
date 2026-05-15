package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var x, d1, d2, q int
	fmt.Fscan(buf, &x, &d1, &d2, &q)
	list := make([]int, 1000)
	list[0] = x
	for i := 1; i < 1000; i++ {
		if i%2 == 0 {
			list[i] = list[i-1] + d1
		} else {
			list[i] = list[i-1] + d2
		}
	}
	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(buf, &k)
		fmt.Println(list[k-1])
	}
}
