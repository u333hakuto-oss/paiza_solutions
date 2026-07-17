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
	fmt.Fscan(buf, &n)
	list := make(map[int]string, n)
	positions := make([]int, n)
	for i := 0; i < n; i++ {
		var s string
		var d int
		fmt.Fscan(buf, &s, &d)
		list[d] = s
		positions[i] = d
	}
	sort.Ints(positions)
	for _, position := range positions {
		fmt.Println(list[position])
	}
}
