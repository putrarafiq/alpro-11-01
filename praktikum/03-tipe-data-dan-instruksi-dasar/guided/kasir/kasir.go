package main 

import "fmt"

func main() {
	var x int

	fmt.Println("Masukkan nominal:")
	fmt.Scan(&x)

	var sepuluhribuan int = x / 10000
	var sisa int = x % 10000

	var limaribuan int = sisa / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println(sepuluhribuan, limaribuan, seribuan)
}