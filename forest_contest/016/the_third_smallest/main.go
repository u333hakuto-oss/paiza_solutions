package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(buf, &n, &s)
	chars := make([]string, n)
	for i, char := range s {
		chars[i] = string(char)
	}
	sort.Strings(chars)
	// 昇順文字列から辞書順3番目を構成
	chars[n-3], chars[n-2] = chars[n-2], chars[n-3]
	for _, char := range chars {
		fmt.Print(char)
	}
	fmt.Println()
}
