package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)

	tahun := hari / 360
	hari %= 360

	bulan := hari / 30
	hari %= 30

	minggu := hari / 7
	sisaHari := hari % 7

	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisaHari)
}