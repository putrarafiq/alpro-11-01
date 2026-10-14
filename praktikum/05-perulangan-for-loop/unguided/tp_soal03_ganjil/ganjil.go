package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		jumlah += i
	}
	fmt.Println()
	fmt.Println(jumlah)
}