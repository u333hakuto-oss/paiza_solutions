package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m, c int
	fmt.Fscan(buf, &n, &m)
	hard := make([]bool, n)
	for i := range hard {
		hard[i] = true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(buf, &c)
			if c == 1 {
				hard[j] = false
			}
		}
	}
	ans := "No"
	for _, allIncorrect := range hard {
		if allIncorrect {
			ans = "Yes"
			break
		}
	}
	fmt.Println(ans)
}
