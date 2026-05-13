package main
import (
    "fmt"
    "bufio"
    "os"
)
const B = 100
func main(){
    buf := bufio.NewReader(os.Stdin)
    var n, k int
    fmt.Fscan(buf, &n, &k)
    all := make([]int, n, n)
    mins := make([]int, n/B+1, n/B+1)
    for i := range mins {
        mins[i] = 200000000
    }
    maxs := make([]int, n/B+1, n/B+1)
    for i := range maxs {
        maxs[i] = -200000000
    }
    for i := 0; i < n; i++ {
        var s int
        fmt.Fscan(buf, &s)
        all[i] = s
        if s < mins[i/B] {mins[i/B] = s}
        if s > maxs[i/B] {maxs[i/B] = s}
    }
    for i := 0; i < k; i++ {
        var al, ar, bl, br int
        fmt.Fscan(buf, &al, &ar, &bl, &br)
        al--; ar--; bl--; br--
        minA, minB, maxA, maxB := 200000000, 200000000, -200000000, -200000000
        for j := al; j <= ar; {
            if j % B == 0 && j + B - 1 <= ar {
                if mins[j/B] < minA {minA = mins[j/B]}
                if maxs[j/B] > maxA {maxA = maxs[j/B]}
                j += B
            } else {
                if all[j] < minA {minA = all[j]}
                if all[j] > maxA {maxA = all[j]}
                j++
            }
        }
        for j := bl; j <= br; {
            if j % B == 0 && j + B -1 <= br {
                if mins[j/B] < minB {minB = mins[j/B]}
                if maxs[j/B] > maxB {maxB = maxs[j/B]}
                j += B
            } else {
                if all[j] < minB {minB = all[j]}
                if all[j] > maxB {maxB = all[j]}
                j++
            }
        }
        a, b := maxA - minA, maxB - minB
        if a > b {
            fmt.Println("A")
        } else if a < b {
            fmt.Println("B")
        } else {
            fmt.Println("DRAW")
        }
    }
}
