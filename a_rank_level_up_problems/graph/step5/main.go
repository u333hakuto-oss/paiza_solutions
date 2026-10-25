package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	if n == 1 {
		fmt.Println(1)
		return
	}
	adjacencyList := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		adjacencyList[a] = append(adjacencyList[a], b)
		adjacencyList[b] = append(adjacencyList[b], a)
	}
	prev, cur, next := -1, -1, 0
	for i := 0; i < n; i++ {
		fmt.Println(next + 1)
		prev, cur = cur, next
		for _, v := range adjacencyList[cur] {
			if v != prev {
				next = v
				break
			}
		}
	}
}
