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
	fmt.Fscan(buf, &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(buf, &s)
		l := len(s)
		switch {
		case strings.HasSuffix(s, "s"), strings.HasSuffix(s, "sh"), strings.HasSuffix(s, "ch"), strings.HasSuffix(s, "o"), strings.HasSuffix(s, "x"):
			fmt.Println(s + "es")
		case strings.HasSuffix(s, "f"):
			fmt.Println(s[:l-1] + "ves")
		case strings.HasSuffix(s, "fe"):
			fmt.Println(s[:l-2] + "ves")
		case strings.HasSuffix(s, "y"):
			switch s[l-2] {
			case 'a', 'i', 'u', 'e', 'o':
				fmt.Println(s + "s")
			default:
				fmt.Println(s[:l-1] + "ies")
			}
		default:
			fmt.Println(s + "s")
		}
	}
}
