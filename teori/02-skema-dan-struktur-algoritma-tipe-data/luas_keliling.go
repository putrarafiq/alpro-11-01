package main

import "fmt"

func main() {
	var p, l, luas, keliling int

	fmt.Scan(&p, &l)
	luas = p * l
	keliling = 2 * (p + l)
	fmt.Println(luas, keliling)
}