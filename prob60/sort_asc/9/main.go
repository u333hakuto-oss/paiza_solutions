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
	list := make(map[int]string, n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var s string
		var d int
		fmt.Fscan(buf, &s, &d)
		list[d] = s
		nums[i] = d
	}
	sort.Ints(nums)
	for _, num := range nums {
		fmt.Println(list[num])
	}
}
