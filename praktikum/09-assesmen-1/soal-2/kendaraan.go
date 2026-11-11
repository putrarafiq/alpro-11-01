package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	info, err := os.Stdin.Stat()
	if err == nil && (info.Mode()&os.ModeCharDevice) != 0 {
		line = "T K M ?"
	} else {
		reader := bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
	}

	m, k, t := 0, 0, 0
	for _, f := range strings.Fields(line) {
		if f == "M" {
			m++
		} else if f == "K" {
			k++
		} else if f == "T" {
			t++
		} else {
			break
		}
	}

	fmt.Printf("Tipe M = %d\n", m)
	fmt.Printf("Tipe K = %d\n", k)
	fmt.Printf("Tipe T = %d\n", t)
}