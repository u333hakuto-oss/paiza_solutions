package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var c, s string
	fmt.Fscan(buf, &c, &s)
	fmt.Println(strings.Count(s, c))
}
