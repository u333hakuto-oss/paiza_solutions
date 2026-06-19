// エラストテネスのふるい

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	const maxValue = 6_000_000
	isPrime := make([]bool, maxValue+1)
	for i := 2; i <= maxValue; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= maxValue; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxValue; j += i {
				isPrime[j] = false
			}
		}
	}
	var n int
	fmt.Fscan(buf, &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(buf, &a)
		if isPrime[a] {
			fmt.Fprintln(out, "pass")
		} else {
			fmt.Fprintln(out, "failure")
		}
	}
}
