package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(buf, &n, &m)
	pos := m % n
	if pos == 0 {
		pos = n
	}
	if (m-1)/n%2 == 0 {
		fmt.Println(m + (n-pos)*2 + 1)
	} else {
		fmt.Println(m - pos*2 + 1)
	}
}
