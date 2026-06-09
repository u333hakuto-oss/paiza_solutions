package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, k, mostExpensive int
	fmt.Fscan(buf, &n, &k)
	presents := make([]int, n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(buf, &x)
		presents[i] = x
		if x > mostExpensive {
			mostExpensive = x
		}
	}
	ans := "No"
	for i := 0; i < k; i++ {
		var y int
		fmt.Fscan(buf, &y)
		if presents[y-1] == mostExpensive {
			ans = "Yes"
			break
		}
	}
	fmt.Println(ans)
}
