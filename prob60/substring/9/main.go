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
	a, _ := strconv.Atoi(arr[0])
	b, _ := strconv.Atoi(arr[1])
	buf.Scan()
	s := buf.Text()
	fmt.Println(s[a-1 : b])
}
