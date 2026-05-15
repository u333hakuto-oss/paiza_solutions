package main
import (
    "fmt"
    "bufio"
    "os"
)
func main(){
    buf := bufio.NewReader(os.Stdin)
    var x, d, q int
    fmt.Fscan(buf, &x, &d, &q)
	list := make([]int, 1000)
	list[0] = x
	for i := 1; i < 1000; i++ {
		list[i] = list[i-1] + d
	}
	for i := 0; i < q; i++ {
	    var k int
	    fmt.Fscan(buf, &k)
		fmt.Println(list[k-1])
	}
}