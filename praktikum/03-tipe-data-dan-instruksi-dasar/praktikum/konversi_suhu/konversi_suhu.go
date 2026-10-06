package main

import "fmt"

func main() {
	var celcius float64
	fmt.Scan(&celcius)
	reamur := 4.0 / 5.0 * celcius
	fmt.Println(reamur)
}
