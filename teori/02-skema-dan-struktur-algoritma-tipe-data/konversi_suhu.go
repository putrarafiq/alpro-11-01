package main

import "fmt"

func main() {
	var c, f float64

	fmt.Scan(&c)
	f = c*9/5 + 32
	fmt.Println(f)
}