package main

import (
	"bufio"
	"fmt"
	"os"
)

func divide(square int, primes []int, isPrime []bool) bool {
	// 奇数を2つの素数の和で表す場合、一方が必ず2になる
	if square%2 == 1 {
		return isPrime[square-2]
	}
	for _, p := range primes {
		if p > square/2 {
			break
		}
		if isPrime[square-p] {
			return true
		}
	}
	return false
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	const limit = 100_000_000
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i <= limit/2; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	conjectureCorrect := true
	for i := 2; i*i <= limit; i++ {
		if !divide(i*i, primes, isPrime) {
			fmt.Fprintln(out, i*i)
			conjectureCorrect = false
		}
	}
	if conjectureCorrect {
		fmt.Fprintln(out, "paiza's conjecture is correct.")
	}
}
