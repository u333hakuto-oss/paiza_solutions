// Sliding Window

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(buf, &s)
	for i := 0; i <= len(s)-5; {
		switch s[i : i+5] {
		case "LLLRB":
			fmt.Println("rolling")
			i += 5
		case "DDRRA":
			fmt.Println("upper")
			i += 5
		case "AAAAA":
			fmt.Println("rush")
			i += 5
		default:
			i++
		}
	}
}
