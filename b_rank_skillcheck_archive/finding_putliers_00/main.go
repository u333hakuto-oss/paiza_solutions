package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	X, Y float64
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(buf, &n)
	points := make([]Coordinate, n)
	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Fscan(buf, &x, &y)
		points[i] = Coordinate{x, y}
	}
	bestOutlierCount := 101
	var outliers []int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			outlierCount := 0
			var curOutliers []int
			a := points[j].Y - points[i].Y
			b := points[i].X - points[j].X
			c := (points[j].X-points[i].X)*points[i].Y - (points[j].Y-points[i].Y)*points[i].X
			for idx, val := range points {
				equation := a*val.X + b*val.Y + c
				if equation*equation >= 4*(a*a+b*b) {
					outlierCount++
					curOutliers = append(curOutliers, idx+1)
				}
			}
			if outlierCount < bestOutlierCount {
				bestOutlierCount = outlierCount
				outliers = curOutliers
			}
		}
	}
	if bestOutlierCount == 0 {
		fmt.Println("none")
	} else {
		for _, o := range outliers {
			fmt.Println(o)
		}
	}
}
