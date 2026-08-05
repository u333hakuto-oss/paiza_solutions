package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(buf, &n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &p[i])
	}
	fmt.Fscan(buf, &q)
	for i := 0; i < q; i++ {
		var c, l, r int
		fmt.Fscan(buf, &c, &l, &r)
		l--
		r--
		switch c {
		case 1:
			res := p[l]
			for i := l + 1; i <= r; i++ {
				res &= p[i]
			}
			fmt.Println(res)
		case 2:
			res := p[l]
			for i := l + 1; i <= r; i++ {
				res |= p[i]
			}
			fmt.Println(res)
		case 3:
			res := p[l]
			for i := l + 1; i <= r; i++ {
				res ^= p[i]
			}
			fmt.Println(res)
		}
	}
}
