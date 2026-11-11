package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func valid(n int) bool {
	jumlah := 0
	prev := 10 
	for n > 0 {
		d := n % 10
		if d > prev {
			return false 
		}
		jumlah += d
		prev = d
		n /= 10
	}
	return isPrima(jumlah)
}

func main() {
	var in io.Reader = os.Stdin
	info, err := os.Stdin.Stat()
	if err == nil && (info.Mode()&os.ModeCharDevice) != 0 {
		in = strings.NewReader("1 1\n")
	}
	reader := bufio.NewReader(in)

	var l, r int
	fmt.Fscan(reader, &l, &r)

	banyak := 0
	for i := l; i <= r; i++ {
		if valid(i) {
			banyak++
		}
	}

	fmt.Println(banyak)
}