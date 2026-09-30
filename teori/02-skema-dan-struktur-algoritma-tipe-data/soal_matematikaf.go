package main

import "fmt"

func main() {
	var x, hasil float64

	fmt.Scan(&x)
	hasil = (x*x + 2*x + 1) / (x - 3)
	fmt.Println(hasil)
}