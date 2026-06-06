// 区間スケジューリング問題　貪欲法
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Task struct {
	Start, End int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(buf, &n)
	tasks := make([]Task, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &tasks[i].Start, &tasks[i].End)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].End < tasks[j].End
	})
	lastSelected := 0
	ans := 1
	for i := 1; i < n; i++ {
		if tasks[i].Start > tasks[lastSelected].End {
			ans++
			lastSelected = i
		}
	}
	fmt.Println(ans)
}
