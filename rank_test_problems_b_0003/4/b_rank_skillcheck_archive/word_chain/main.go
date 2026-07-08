package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, k, m int
	fmt.Fscan(buf, &n, &k, &m)
	players := make([]bool, n)
	for i := range players {
		players[i] = true
	}
	usableWords := make(map[string]bool, k)
	for i := 0; i < k; i++ {
		var d string
		fmt.Fscan(buf, &d)
		usableWords[d] = true
	}
	crr := 0
	var before byte
	before = 'z'
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(buf, &s)
		isUsable, ok := usableWords[s]
		if ok && isUsable && (before == s[0] || before == 'z') && s[len(s)-1] != 'z' {
			usableWords[s] = false
			before = s[len(s)-1]
		} else {
			players[crr] = false
			before = 'z'
		}
		for {
			crr++
			crr %= n
			if players[crr] {
				break
			}
		}
	}
	count := 0
	for _, val := range players {
		if val {
			count++
		}
	}
	fmt.Println(count)
	for i, val := range players {
		if val {
			fmt.Println(i + 1)
		}
	}
}
