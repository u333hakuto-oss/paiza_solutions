package main

import (
	"bufio"
	"fmt"
	"os"
)

const Size = 5

func Winner(b [][]rune) rune {
	var p rune
	// 縦
	for j := 0; j < Size; j++ {
		p = b[0][j]
		if p == '.' {
			continue
		}
		for i := 1; i < Size; i++ {
			if b[i][j] != p {
				break
			}
			if i == Size-1 {
				return p
			}
		}
	}
	// 横
	for i := 0; i < Size; i++ {
		p = b[i][0]
		if p == '.' {
			continue
		}
		for j := 1; j < Size; j++ {
			if b[i][j] != p {
				break
			}
			if j == Size-1 {
				return p
			}
		}
	}
	// 斜め
	p = b[0][0]
	if p != '.' {
		for i := 1; i < Size; i++ {
			if b[i][i] != p {
				break
			}
			if i == Size-1 {
				return p
			}
		}
	}
	p = b[4][0]
	if p != '.' {
		for i := 1; i < Size; i++ {
			if b[4-i][i] != p {
				break
			}
			if i == Size-1 {
				return p
			}
		}
	}
	return 'D'
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	board := make([][]rune, Size)
	for i := 0; i < 5; i++ {
		var s string
		fmt.Fscan(buf, &s)
		board[i] = []rune(s)
	}
	fmt.Println(string(Winner(board)))
}
