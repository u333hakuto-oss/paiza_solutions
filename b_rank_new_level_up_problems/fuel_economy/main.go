package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var x, f1, f2, l, n, fuel int
	fmt.Fscan(buf, &x, &f1, &f2, &l, &n)
	lastStop := 0
	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(buf, &s)
		diff := s - lastStop
		if diff > x {
			fuel += x*f1 + (diff-x)*f2
		} else {
			fuel += diff * f1
		}
		lastStop = s
	}
	diff := l - lastStop
	if diff > x {
		fuel += x*f1 + (diff-x)*f2
	} else {
		fuel += diff * f1
	}
	fmt.Println(fuel)
}
