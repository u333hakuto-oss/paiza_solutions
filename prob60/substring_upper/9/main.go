package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	arr := strings.Split(buf.Text(), " ")
	n, _ := strconv.Atoi(arr[0])
	m, _ := strconv.Atoi(arr[1])
	buf.Scan()
	s := buf.Text()
	fmt.Println(s[:n-1] + strings.ToUpper(s[n-1:m]) + s[m:])
}
