package main

import (
	"bufio"
	"fmt"
	"os"
)

type Record struct {
	b    float64
	r, h int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var (
		n    int
		maxB float64
		maxR int
		maxH int
	)
	fmt.Fscan(buf, &n)
	records := make([]Record, n)
	for i := 0; i < n; i++ {
		var b float64
		var r, h int
		fmt.Fscan(buf, &b, &r, &h)
		records[i] = Record{b, r, h}
		if b > maxB {
			maxB = b
		}
		if r > maxR {
			maxR = r
		}
		if h > maxH {
			maxH = h
		}
	}
	ans := "Nobody"
	for _, val := range records {
		crown := 0
		if val.b == maxB {
			crown++
		}
		if val.r == maxR {
			crown++
		}
		if val.h == maxH {
			crown++
		}
		if crown == 3 {
			ans = "Triple"
			break
		} else if crown == 2 {
			ans = "Double"
		}
	}
	fmt.Println(ans)
}
