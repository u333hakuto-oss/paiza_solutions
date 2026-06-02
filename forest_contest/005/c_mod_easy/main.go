package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	diff := z - x%y
	if diff < 0 {
		diff += y
	}
	fmt.Println(x + diff)
}
