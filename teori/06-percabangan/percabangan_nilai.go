package main

import "fmt"

func main() {
    nilai := 123

    if nilai >= 0 && nilai < 50 {
        fmt.Println("E")
    } else if nilai >= 50 && nilai < 70 {
        fmt.Println("C")
    } else if nilai >= 70 && nilai < 85 {
        fmt.Println("B")
    } else if nilai >= 85 && nilai <= 100 {
        fmt.Println("A")
    } else {
        fmt.Println("Perintah tidak terdaftar di sini")
    }
}