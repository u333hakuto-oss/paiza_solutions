package main

import (
	"bufio"
	"fmt"
	"os"
)

type Friendship struct {
	VillagerA, VillagerB, FriendLevel int
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	var n, m, q int
	fmt.Fscan(buf, &n, &m, &q)
	friendships := make([]Friendship, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(buf, &friendships[i].VillagerA, &friendships[i].VillagerB, &friendships[i].FriendLevel)
	}
	members := make([]bool, n+1)
	for i := 0; i < q; i++ {
		var op string
		var m int
		fmt.Fscan(buf, &op, &m)
		members[m] = !members[m]
	}
	popularity := 0
	for _, f := range friendships {
		if members[f.VillagerA] != members[f.VillagerB] {
			popularity = max(popularity, f.FriendLevel)
		}
	}
	fmt.Println(popularity)
}
