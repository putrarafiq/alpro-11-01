package main

import "fmt"

func main() {
	var tahun int
	var bulan string
	fmt.Scan(&tahun, &bulan)

	hari := 0
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		hari = 31
	case "Apr", "Jun", "Sep", "Nov":
		hari = 30
	case "Feb":
		if (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0 {
			hari = 29
		} else {
			hari = 28
		}
	default:
		fmt.Println("Nama bulan tidak valid")
		return
	}

	fmt.Println(hari)
}