package main

import "fmt"

func main() {
	var l1, r1, l2, r2 int
	fmt.Scan(&l1, &r1, &l2, &r2)
	if r1 < l2 || r2 < l1 {
		// 日程が重ならない
		fmt.Println(0)
	} else {
		var start, end int
		// 開始日
		if l1 > l2 {
			start = l1
		} else {
			start = l2
		}
		// 終了日
		if r1 < r2 {
			end = r1
		} else {
			end = r2
		}
		fmt.Println(end - start + 1)
	}
}
