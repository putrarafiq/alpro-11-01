package main

import "fmt"

func main() {
	var penghasilan float64
	fmt.Scan(&penghasilan)

	var pajak float64

	if penghasilan <= 50 {
		// 5% dari seluruh penghasilan
		pajak = penghasilan * 5.0 / 100.0
	} else if penghasilan <= 100 {
		// 5% x 50 + 10% x (penghasilan - 50)
		pajak = 50.0*5.0/100.0 + (penghasilan-50.0)*10.0/100.0
	} else if penghasilan <= 200 {
		// 5% x 50 + 10% x 50 + 15% x (penghasilan - 100)
		pajak = 50.0*5.0/100.0 + 50.0*10.0/100.0 + (penghasilan-100.0)*15.0/100.0
	} else {
		// 5% x 50 + 10% x 50 + 15% x 100 + 20% x (penghasilan - 200)
		pajak = 50.0*5.0/100.0 + 50.0*10.0/100.0 + 100.0*15.0/100.0 + (penghasilan-200.0)*20.0/100.0
	}

	fmt.Println(pajak)
}