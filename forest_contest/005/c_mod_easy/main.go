package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	for {
		if x%y == z {
			break
		}
		x++
	}
	fmt.Println(x)
}
