package main

import "fmt"

func main() {
	var harga, persen int
	var potongan, hargaAkhir float64

	fmt.Scan(&harga, &persen)
	potongan = float64(harga*persen) / 100
	hargaAkhir = float64(harga) - potongan
	fmt.Println(hargaAkhir)
}