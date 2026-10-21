package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	nama1, nama2 := "", ""
	min1, min2 := math.MaxInt, math.MaxInt

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < min1 {
			min2, nama2 = min1, nama1
			min1, nama1 = nilai, nama
		} else if nilai < min2 {
			min2, nama2 = nilai, nama
		}
		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}