package main

import "fmt"

func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0
}

func hariDalamBulan(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if kabisat(tahun) {
			return 29
		}
		return 28
	}
	return 0
}

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	// Validasi
	if tahun <= 0 || bulan < 1 || bulan > 12 || tanggal < 1 || tanggal > hariDalamBulan(bulan, tahun) {
		fmt.Println("Tidak valid")
		return
	}

	// Hitung tanggal besok
	tanggal++
	if tanggal > hariDalamBulan(bulan, tahun) {
		tanggal = 1
		bulan++
		if bulan > 12 {
			bulan = 1
			tahun++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}