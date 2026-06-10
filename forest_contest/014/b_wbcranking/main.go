package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Team struct {
	Country string
	Wins    int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	scoreboard := make([]Team, 5)
	for i := 0; i < 5; i++ {
		var c string
		fmt.Fscan(buf, &c)
		scoreboard[i].Country = c
	}
	for i := 0; i < 10; i++ {
		var s1, winner string
		var s2 int
		fmt.Fscan(buf, &s1, &s2)
		cs := strings.Split(s1, "-")
		if s2 == 1 {
			winner = cs[0]
		} else {
			winner = cs[1]
		}
		for i := range scoreboard {
			if scoreboard[i].Country == winner {
				scoreboard[i].Wins++
				break
			}
		}
	}
	sort.Slice(scoreboard, func(i, j int) bool {
		if scoreboard[i].Wins != scoreboard[j].Wins {
			return scoreboard[i].Wins > scoreboard[j].Wins
		} else {
			return scoreboard[i].Country < scoreboard[j].Country
		}
	})
	fmt.Println(scoreboard[0].Country)
	fmt.Println(scoreboard[1].Country)
}
