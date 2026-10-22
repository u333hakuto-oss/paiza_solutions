package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Adjacent struct {
	node, weight int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	adjacencyMatrix := make([][]int, n)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]int, n)
	}
	adjacencyList := make([][]Adjacent, n)
	for i := 0; i < m; i++ {
		var a, b, k int
		fmt.Fscan(in, &a, &b, &k)
		a--
		b--
		adjacencyMatrix[a][b] = k
		adjacencyList[a] = append(adjacencyList[a], Adjacent{b, k})
	}
	for _, row := range adjacencyMatrix {
		for _, isAdjacent := range row {
			fmt.Print(isAdjacent)
		}
		fmt.Println()
	}
	for _, neighbors := range adjacencyList {
		sort.Slice(neighbors, func(i, j int) bool {
			return neighbors[i].node < neighbors[j].node
		})
		for _, neighbor := range neighbors {
			fmt.Printf("%d(%d)", neighbor.node, neighbor.weight)
		}
		fmt.Println()
	}
}
