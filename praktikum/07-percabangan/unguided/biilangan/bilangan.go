package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0 adalah bilangan nol")
	} else {
		var jenis, tanda string

		if n%2 == 0 {
			jenis = "Genap"
		} else {
			jenis = "Ganjil"
		}

		if n > 0 {
			tanda = "Positif"
		} else {
			tanda = "Negatif"
		}

		fmt.Printf("%s, %s\n", jenis, tanda)
	}
}