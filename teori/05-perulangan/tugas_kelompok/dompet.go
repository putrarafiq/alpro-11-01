package main

import "fmt"

func main() {
	var (
		transaksi, saldo int
	)

	saldo = 0
	fmt.Scan(&transaksi)
	for transaksi != 0 {
		saldo = saldo + transaksi
		fmt.Scan(&transaksi)
	}

	fmt.Println(saldo)
}