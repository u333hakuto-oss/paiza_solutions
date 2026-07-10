package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	t, y, x int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(buf, &n)
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		var t, y, x int
		fmt.Fscan(buf, &t, &y, &x)
		points[i] = Point{t, y, x}
	}
	seg := 0
	for t := 0; t <= 100; t++ {
		if t > points[seg+1].t {
			seg++
		}
		cur := points[seg]
		next := points[seg+1]
		ansY := cur.y + (next.y-cur.y)*(t-cur.t)/(next.t-cur.t)
		ansX := cur.x + (next.x-cur.x)*(t-cur.t)/(next.t-cur.t)
		fmt.Println(ansY, ansX)
	}
}
