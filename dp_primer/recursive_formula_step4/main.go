package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(buf, &k)
	list := make([]int, k)
	list[0], list[1] = 1, 1
	for i := 2; i < k; i++ {
		list[i] = list[i-2] + list[i-1]
	}
	fmt.Println(list[k-1])
}
