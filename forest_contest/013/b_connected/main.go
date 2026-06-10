// 連結判定　Depth First Search

package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(v int, edges [][]int, visited []bool) {
	visited[v] = true
	for _, next := range edges[v] {
		if !visited[next] {
			dfs(next, edges, visited)
		}
	}
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(buf, &n, &m)
	edges := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(buf, &a, &b)
		if a != b {
			edges[a] = append(edges[a], b)
			edges[b] = append(edges[b], a)
		}
	}
	visited := make([]bool, n+1)
	dfs(1, edges, visited)
	ans := "Yes"
	for i := 1; i <= n; i++ {
		if !visited[i] {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}
