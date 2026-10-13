// 条件を満たすAの部分数列が、必ず存在するものとする

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	var k int64
	fmt.Fscan(in, &n, &k)
	if k <= 2 {
		fmt.Println(1)
		return
	}
	b := strconv.FormatInt(k-1, 2)
	requiredTwos := len(b)
	minLen := -1
	left := 0
	count := 0
	a := make([]int, n)
	for right := 0; right < n; right++ {
		fmt.Fscan(in, &a[right])
		switch a[right] {
		case 0:
			left = right + 1
			count = 0
		case 2:
			count++
		}
		for count >= requiredTwos {
			curLen := right - left + 1
			if minLen == -1 || curLen < minLen {
				minLen = curLen
			}
			if a[left] == 2 {
				count--
			}
			left++
		}
	}
	fmt.Println(minLen)
}
