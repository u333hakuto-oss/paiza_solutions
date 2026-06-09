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
	easy := make([]bool, n)
	for i := range easy {
		easy[i] = true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(buf, &c)
			if c == 0 {
				easy[j] = false
			}
		}
	}
	ans := "No"
	for _, allCorrect := range easy {
		if allCorrect {
			ans = "Yes"
			break
		}
	}
	fmt.Println(ans)
}
