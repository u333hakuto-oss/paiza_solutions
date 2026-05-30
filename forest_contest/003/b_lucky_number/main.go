// bit全探索
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
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(buf, &xs[i])
	}
	var answers [][]int
	for i := 0; i < 1<<n; i++ {
		var sum int
		var selected []int
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				sum += xs[j]
				selected = append(selected, xs[j])
			}
		}
		if sum == 777 {
			answers = append(answers, selected)
		}
	}
	if len(answers) == 0 {
		fmt.Println("no answer")
	} else if len(answers) > 1 {
		fmt.Println("multiple answers")
	} else {
		ans := answers[0]
		sort.Ints(ans)
		for i, x := range ans {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(x)
		}
		fmt.Println()
	}
}
