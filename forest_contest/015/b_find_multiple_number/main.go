package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if (b <= a && a <= c) || c-b >= a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
