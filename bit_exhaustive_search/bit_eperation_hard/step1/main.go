package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const Bits = 30
	buf := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(buf, &n)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(buf, &p[i])
	}
	cnt := make([][]int, Bits)
	for b := 0; b < Bits; b++ {
		cnt[b] = make([]int, n+1)
		for i := 1; i <= n; i++ {
			cnt[b][i] = cnt[b][i-1] + (p[i]>>b)&1
		}
	}
	xors := make([]int, n+1)
	for i := 1; i <= n; i++ {
		xors[i] = xors[i-1] ^ p[i]
	}
	fmt.Fscan(buf, &q)
	for i := 0; i < q; i++ {
		var c, l, r int
		fmt.Fscan(buf, &c, &l, &r)
		switch c {
		case 1:
			ans := 0
			for b := 0; b < Bits; b++ {
				if cnt[b][r]-cnt[b][l-1] == r-l+1 {
					ans |= 1 << b
				}
			}
			fmt.Println(ans)
		case 2:
			ans := 0
			for b := 0; b < Bits; b++ {
				if cnt[b][r]-cnt[b][l-1] > 0 {
					ans |= 1 << b
				}
			}
			fmt.Println(ans)
		case 3:
			fmt.Println(xors[r] ^ xors[l-1])
		}
	}
}
