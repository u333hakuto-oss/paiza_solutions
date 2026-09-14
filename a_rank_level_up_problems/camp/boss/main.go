package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Y, X int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var h, w int
	var n string
	fmt.Fscan(in, &h, &w, &n)
	var firstPlayer, secondPlayer byte
	switch n {
	case "A":
		firstPlayer, secondPlayer = 'A', 'B'
	case "B":
		firstPlayer, secondPlayer = 'B', 'A'
	}
	grid := make([][]byte, h)
	var queueF []Point
	var queueS []Point
	for i := range grid {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
		for j := range grid[i] {
			switch grid[i][j] {
			case firstPlayer:
				queueF = append(queueF, Point{i, j})
			case secondPlayer:
				queueS = append(queueS, Point{i, j})
			}
		}
	}
	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}
	isFirstPlayerturn := true
	startF := 0
	endF := 1
	startS := 0
	endS := 1
	countF := 1
	countS := 1
	for startF < endF || startS < endS {
		var start, end int
		if isFirstPlayerturn {
			start, end = startF, endF
		} else {
			start, end = startS, endS
		}
		for i := start; i < end; i++ {
			var cur Point
			if isFirstPlayerturn {
				cur = queueF[i]
			} else {
				cur = queueS[i]
			}
			for d := 0; d < 4; d++ {
				ny := cur.Y + dy[d]
				nx := cur.X + dx[d]
				if ny < 0 || ny >= h || nx < 0 || nx >= w {
					continue
				}
				if grid[ny][nx] == '.' {
					if isFirstPlayerturn {
						grid[ny][nx] = firstPlayer
						countF++
					} else {
						grid[ny][nx] = secondPlayer
						countS++
					}
					if isFirstPlayerturn {
						queueF = append(queueF, Point{ny, nx})
					} else {
						queueS = append(queueS, Point{ny, nx})
					}
				}
			}
		}
		if isFirstPlayerturn {
			startF = endF
			endF = len(queueF)
		} else {
			startS = endS
			endS = len(queueS)
		}
		isFirstPlayerturn = !isFirstPlayerturn
	}
	winner := firstPlayer
	if countS > countF {
		winner = secondPlayer
	}
	switch firstPlayer {
	case 'A':
		fmt.Println(countF, countS)
	case 'B':
		fmt.Println(countS, countF)
	}
	fmt.Println(string(winner))
}
