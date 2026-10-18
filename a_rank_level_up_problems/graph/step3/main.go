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
	adjacencyMatrix := make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}
	adjacencyList := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		adjacencyMatrix[a][b] = 1
		adjacencyList[a] = append(adjacencyList[a], b)
	}
	for _, row := range adjacencyMatrix {
		for _, isAdjacent := range row {
			fmt.Print(isAdjacent)
		}
		fmt.Println()
	}
	for _, neighbors := range adjacencyList {
		sort.Ints(neighbors)
		for _, neighbor := range neighbors {
			fmt.Print(neighbor)
		}
		fmt.Println()
	}
}
