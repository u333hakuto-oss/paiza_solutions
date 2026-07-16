package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(state int, memo []int, panels []int, bonus [][]int) int {
	if state == (1<<9)-1 {
		return 0
	}
	if memo[state] != -1 {
		return memo[state]
	}
	maxScore := 0
	for panel := 1; panel <= 9; panel++ {
		if (state & (1 << (panel - 1))) == 0 { // まだ倒していないなら
			nextState := state | (1 << (panel - 1))
			gain := calculateGain(state, nextState, panel, panels, bonus)
			score := gain + solve(nextState, memo, panels, bonus)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	memo[state] = maxScore
	return maxScore
}

func calculateGain(state, nextState, panel int, panels []int, bonus [][]int) int {
	gain := panels[panel]
	bingo := []int{
		0b000000111, // 1, 2, 3
		0b000111000, // 4, 5, 6
		0b111000000, // 7, 8, 9
		0b001001001, // 1, 4, 7
		0b010010010, // 2, 5, 8
		0b100100100, // 3, 6, 9
		0b100010001, // 1, 5, 9
		0b001010100, // 3, 5, 7
	}
	var countBefore, countAfter int
	for _, val := range bingo {
		if (state & val) == val {
			countBefore++
		}
		if (nextState & val) == val {
			countAfter++
		}
	}
	newBingo := countAfter - countBefore
	gain += bonus[panel][newBingo]
	return gain
}

func main() {
	const p = 9
	buf := bufio.NewReader(os.Stdin)
	panels := make([]int, p+1)
	for i := 1; i <= p; i++ {
		fmt.Fscan(buf, &panels[i])
	}
	bingoLimits := []int{0, 3, 2, 3, 2, 4, 2, 3, 2, 3}
	bonus := make([][]int, p+1)
	for i := 1; i <= p; i++ {
		size := bingoLimits[i]
		bonus[i] = make([]int, size+1)
		for j := 1; j <= size; j++ {
			fmt.Fscan(buf, &bonus[i][j])
		}
	}
	memo := make([]int, 512)
	for i := range memo {
		memo[i] = -1
	}
	ans := solve(0, memo, panels, bonus)
	fmt.Println(ans)
}
