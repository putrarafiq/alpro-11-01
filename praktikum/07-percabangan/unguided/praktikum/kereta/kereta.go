package main

import (
	"fmt"
	"strings"
)

func main() {
	var kelas, hari string
	var usia int
	fmt.Scan(&kelas, &usia, &hari)

	kelas = strings.ToLower(kelas)
	hari = strings.ToLower(hari)

	var harga int
	switch kelas {
	case "ekonomi":
		harga = 100000
	case "bisnis":
		harga = 200000
	case "eksekutif":
		harga = 350000
	default:
		fmt.Println("Kelas tidak valid")
		return
	}

	if hari == "sabtu" || hari == "minggu" {
		harga = harga * 120 / 100
	}

	if usia < 3 {
		harga = 0
	} else if usia <= 12 {
		harga = harga * 50 / 100
	} else if usia >= 60 {
		harga = harga * 70 / 100
	}

	fmt.Println(harga)
}