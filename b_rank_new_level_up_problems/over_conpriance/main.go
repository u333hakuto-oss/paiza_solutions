package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(buf, &n, &s)
	frontLen := (len(s) + 1) / 2
	backStart := len(s) / 2
	x := strings.Repeat("x", frontLen)
	for i := 0; i < n; i++ {
		var v string
		fmt.Fscan(buf, &v)
		if len(s) != len(v) {
			fmt.Println(v)
			continue
		} else if v == s {
			fmt.Println("banned")
		} else if s[:frontLen] == v[:frontLen] {
			fmt.Println(x + v[frontLen:])
		} else if s[backStart:] == v[backStart:] {
			fmt.Println(v[:backStart] + x)
		} else {
			fmt.Println(v)
		}
	}
}
