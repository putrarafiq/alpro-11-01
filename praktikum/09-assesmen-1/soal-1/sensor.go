package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var line string
	info, err := os.Stdin.Stat()
	if err == nil && (info.Mode()&os.ModeCharDevice) != 0 {
		line = "12 12"
	} else {
		reader := bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
	}
	fields := strings.Fields(line)

	var data []int
	for _, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			continue
		}
		data = append(data, n)
		if len(data) >= 2 && data[len(data)-1] == data[len(data)-2] {
			break
		}
	}

	if len(data) == 0 {
		return
	}
	data = data[:len(data)-1]
	if len(data) == 0 {
		return
	}

	tertinggi := data[0]
	terendah := data[0]
	jumlah := 0
	for _, v := range data {
		if v > tertinggi {
			tertinggi = v
		}
		if v < terendah {
			terendah = v
		}
		jumlah += v
	}
	rata := float64(jumlah) / float64(len(data))

	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rata)
}