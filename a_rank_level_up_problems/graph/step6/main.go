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
		fmt.Println(0)
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
	apples := make([]int, n)
	for i := range apples {
		fmt.Fscan(in, &apples[i])
	}
	s := 0
	prev, cur, next := -1, -1, 0
	for i := 0; i < n; i++ {
		prev, cur = cur, next
		for _, v := range adjacencyList[cur] {
			if v != prev {
				next = v
				break
			}
		}
		s += apples[cur]
		fmt.Println(s)
	}
}
