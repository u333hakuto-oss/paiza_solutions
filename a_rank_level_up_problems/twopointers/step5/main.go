// いもす法(差分配列)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	list := make([]int, n)
	for i := range list {
		fmt.Fscan(in, &list[i])
	}
	d := make([]int, n)
	for i := 0; i < m; i++ {
		var l, u, a int
		fmt.Fscan(in, &l, &u, &a)
		l--
		u--
		d[l] += a
		if u < n-1 {
			d[u+1] -= a
		}
	}
	add := 0
	for i := range list {
		add += d[i]
		fmt.Println(list[i] + add)
	}
}
