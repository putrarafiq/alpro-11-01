package main

import "fmt"

func fpb(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n, m int64
	fmt.Scan(&n, &m)
	fmt.Println(n / fpb(n, m) * m)
}