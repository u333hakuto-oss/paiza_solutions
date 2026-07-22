package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(buf, &s)
	if strings.Contains(s, "paiza") {
		fmt.Println("paiza")
		return
	}
	for i := 0; i <= len(s)-5; i++ {
		if s[i] != 'p' {
			continue
		}
		if s[i+1] != 'a' && s[i+1] != '4' && s[i+1] != '@' {
			continue
		}
		if s[i+2] != 'i' && s[i+2] != '1' && s[i+2] != '!' {
			continue
		}
		if s[i+3] != 'z' && s[i+3] != '2' {
			continue
		}
		if s[i+4] != 'a' && s[i+4] != '4' && s[i+4] != '@' {
			continue
		}
		fmt.Println("leet")
		return
	}
	fmt.Println("nothing")
}
