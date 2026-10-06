package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	puluhan := n / 10
	satuan := n % 10

	hasil := puluhan*1000 + puluhan*100 + satuan*10 + satuan
	fmt.Println(hasil)
}