// 0/1ナップサック問題　Dynamic Programming

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(buf, &n, &k)
	ways := make([]int, k+1)
	ways[0] = 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(buf, &a)
		for j := k; j >= a; j-- {
			ways[j] += ways[j-a]
		}
	}
	fmt.Println(ways[k])
}
