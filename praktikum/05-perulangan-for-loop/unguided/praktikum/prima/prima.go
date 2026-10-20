package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	faktor := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			faktor = append(faktor, i)
		}
	}

	for i, f := range faktor {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(f)
	}
	fmt.Println()

	if len(faktor) == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}