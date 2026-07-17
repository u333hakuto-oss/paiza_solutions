package main
import (
    "fmt"
    "bufio"
    "os"
)
func main(){
    buf := bufio.NewReader(os.Stdin)
    var s string
    fmt.Fscan(buf, &s)
	var positions []int
	for i, r := range s {
		if '0' <= r && r <= '9' {
			positions = append(positions, i)
		}
	}
	for i := 0; i < len(positions); i++ {
		for j := i; j < len(positions); j++ {
			fmt.Println(s[positions[i]:positions[j]+1])
		} 
	}
}