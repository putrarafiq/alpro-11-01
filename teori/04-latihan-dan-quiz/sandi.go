package main

import "fmt"

func main() {
    var n, x, total int
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&x)
        total += x/1000 + x%10
    }
    fmt.Println(total)
}