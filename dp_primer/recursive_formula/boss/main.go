package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	list := make([]int, 40)
	list[0], list[1] = 1, 1
	for i := 2; i < 40; i++ {
		list[i] = list[i-2] + list[i-1]
	}
	var q int
	fmt.Fscan(buf, &q)
	for i := 0; i < q; i++ {
		var k int
		fmt.Fscan(buf, &k)
		fmt.Println(list[k-1])
	}
}
