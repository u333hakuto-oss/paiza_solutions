package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var s, t string
	fmt.Fscan(buf, &s, &t)
	count, pos := 0, 0
	for {
		idx := strings.Index(t[pos:], s)
		if idx == -1 {
			break
		}
		count++
		pos += idx + 1
	}
	fmt.Println(count)
}
