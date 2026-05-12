// 平方分割
package main
import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()
    var n, k int
    fmt.Fscan(reader, &n, &k)
    // 生徒数の平方根をバケットサイズとして設定
    bucket := 1
    if bucket * bucket < n {
        bucket++
    }
    
    // 全生徒の得点、バケット内の最高点・最低点のスライスを作成
    all := make([]int, n)
    sliceSize := n / bucket + 1
    maxs := make([]int, sliceSize)
    for i := range maxs {
        maxs[i] = -200000000
    }
    mins := make([]int, sliceSize)
    for i := range mins {
        mins[i] = 200000000
    }
    for i := 0; i < n; i++ {
        var s int
        fmt.Fscan(reader, &s)
        all[i] = s
        bucketIdx := i / bucket
        if s > maxs[bucketIdx] {maxs[bucketIdx] = s}
        if s < mins[bucketIdx] {mins[bucketIdx] = s}
    }
    
    // ジャッジ
    for i := 0; i < k; i++ {
        var al, ar, bl, br int
        fmt.Fscan(reader, &al, &ar, &bl, &br)
        al--; ar--; bl--; br--
        // プレイヤーAの最高点と最低点を求める
        maxA, minA := -200000000, 200000000
        for j := al; j <= ar; {
            if j % bucket == 0 && j + bucket - 1 <= ar {
                bucketIdx := j / bucket
                if maxs[bucketIdx] > maxA {maxA = maxs[bucketIdx]}
                if mins[bucketIdx] < minA {minA = mins[bucketIdx]}
                j += bucket
            } else {
                if all[j] > maxA {maxA = all[j]}
                if all[j] < minA {minA = all[j]}
                j++
            }
        }
        // プレイヤーBの最高点と最低点を求める
        maxB, minB := -200000000, 200000000
        for j := bl; j <= br; {
            if j % bucket == 0 && j + bucket - 1 <= br {
                bucketIdx := j / bucket
                if maxs[bucketIdx] > maxB {maxB = maxs[bucketIdx]}
                if mins[bucketIdx] < minB {minB = mins[bucketIdx]}
                j += bucket
            } else {
                if all[j] > maxB {maxB = all[j]}
                if all[j] < minB {minB = all[j]}
                j++
            }
        }
        
        // プレイヤーA・Bの得点の幅を比較
        diffA, diffB := maxA - minA, maxB - minB
        if diffA > diffB {
            fmt.Fprintln(writer, "A")
        } else if diffA < diffB {
            fmt.Fprintln(writer, "B")
        } else {
            fmt.Fprintln(writer, "DRAW")
        }
    }
}