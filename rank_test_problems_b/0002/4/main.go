package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m, q int
	fmt.Fscan(buf, &n, &m, &q)
	board := make([][]string, n)
	for i := range board {
		board[i] = make([]string, m)
	}
	for i := 0; i < q; i++ {
		var p, notP string
		var x, y int
		fmt.Fscan(buf, &p, &x, &y)
		x--
		y--
		if p == "A" {
			notP = "B"
		} else {
			notP = "A"
		}
		switch board[x][y] {
		case "":
			board[x][y] = p
		case p:
			board[x][y] += "L"
		case notP:
			board[x][y] = ""
		}
	}
	var countA, countB int
	for _, row := range board {
		for _, whose := range row {
			switch whose {
			case "A", "AL":
				countA++
			case "B", "BL":
				countB++
			}
		}
	}
	if countA > countB {
		fmt.Println("A")
	} else if countB > countA {
		fmt.Println("B")
	} else {
		fmt.Println("Draw")
	}
}
