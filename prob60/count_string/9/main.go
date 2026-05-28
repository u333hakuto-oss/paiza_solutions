package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var s, t string
	fmt.Fscan(buf, &s, &t)
	ls, lt := len(s), len(t)
	if lt < ls {
		fmt.Println(0)
	} else {
		count := 0
		for i := 0; i <= lt-ls; i++ {
			if t[i:i+ls] == s {
				count++
			}
		}
		fmt.Println(count)
	}
}
