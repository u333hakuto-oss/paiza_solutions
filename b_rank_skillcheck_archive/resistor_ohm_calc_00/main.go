package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(buf, &n)
	resistors := make(map[byte]float64, n)
	for i := 0; i < n; i++ {
		var s string
		var w float64
		fmt.Fscan(buf, &s, &w)
		resistors[s[0]] = w
	}
	fmt.Fscan(buf, &m)
	var ohm float64
	for i := 0; i < m; i++ {
		var t string
		fmt.Fscan(buf, &t)
		if len(t) == 1 {
			ohm += resistors[t[0]]
		} else {
			var bottom float64
			for j := 0; j < len(t); j++ {
				bottom += 1 / resistors[t[j]]
			}
			ohm += 1 / bottom
		}
	}
	fmt.Println(int(ohm))
}
