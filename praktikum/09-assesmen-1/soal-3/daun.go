package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var in io.Reader = os.Stdin
	info, err := os.Stdin.Stat()
	if err == nil && (info.Mode()&os.ModeCharDevice) != 0 {
		in = strings.NewReader("10\n11 32 53 64 85 96 57 38 29 80 \n")
	}
	reader := bufio.NewReader(in)

	var n int
	fmt.Fscan(reader, &n)

	var lebar, tersempit int
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &lebar)
		if i == 0 || lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println(tersempit)
}