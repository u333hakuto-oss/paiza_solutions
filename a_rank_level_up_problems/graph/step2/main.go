package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	for _, neighbors := range g {
		sort.Ints(neighbors)
		for _, neighbor := range neighbors {
			fmt.Print(neighbor)
		}
		fmt.Println()
	}
}
