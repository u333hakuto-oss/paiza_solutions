package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var s string
	reg := regexp.MustCompile("LLLRB|DDRRA|AAAAA")
	fmt.Fscan(buf, &s)
	for {
		loc := reg.FindStringIndex(s)
		if loc == nil {
			break
		}
		switch s[loc[0]] {
		case 'L':
			fmt.Println("rolling")
		case 'D':
			fmt.Println("upper")
		case 'A':
			fmt.Println("rush")
		}
		s = s[loc[1]:]
	}
}
