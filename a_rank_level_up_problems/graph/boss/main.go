// BFS

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	adjacencyList := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		adjacencyList[a] = append(adjacencyList[a], b)
		adjacencyList[b] = append(adjacencyList[b], a)
	}
	visited := make([]bool, n)
	queue := []int{0}
	visited[0] = true
	for len(queue) != 0 {
		cur := queue[0]
		for _, adjacent := range adjacencyList[cur] {
			if !visited[adjacent] {
				visited[adjacent] = true
				queue = append(queue, adjacent)
			}
		}
		queue = queue[1:]
	}
	ans := "YES"
	for _, val := range visited {
		if !val {
			ans = "NO"
			break
		}
	}
	fmt.Println(ans)
}
