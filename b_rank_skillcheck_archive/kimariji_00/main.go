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
	list := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &list[i])
	}
	kimariji := 1
	target := list[m-1]
	for i, val := range list {
		if i == m-1 {
			continue
		}
		diff := -1
		for j := 0; j < len(target); j++ {
			if target[j] != val[j] {
				diff = j + 1
				break
			}
		}
		if diff > kimariji {
			kimariji = diff
		}
	}
	fmt.Println(target[:kimariji])
}
