package main

import "fmt"

func main() {
	const passwordBenar = "1234"
	const maksKesempatan = 3

	saldo := 10000000
	kesempatan := 0
	login := false

	for !login && kesempatan < maksKesempatan {
		var password string
		fmt.Print("Masukkan password ATM: ")
		if _, err := fmt.Scan(&password); err != nil {
			fmt.Println("Input tidak valid. Program selesai.")
			return
		}

		if password == passwordBenar {
			login = true
		} else {
			kesempatan++
			fmt.Println("Password salah. Silakan coba lagi.")
		}
	}

	if !login {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Println("Login berhasil!")

	pilihan := 0
	for pilihan != 3 {
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		if _, err := fmt.Scan(&pilihan); err != nil {
			fmt.Println("Input tidak valid. Program selesai.")
			return
		}

		if pilihan == 1 {
			fmt.Println("Saldo Anda: Rp", saldo)
		} else if pilihan == 2 {
			var jumlah int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			if _, err := fmt.Scan(&jumlah); err != nil {
				fmt.Println("Input tidak valid. Program selesai.")
				return
			}

			if jumlah <= 0 {
				fmt.Println("Jumlah penarikan harus lebih dari 0.")
			} else if jumlah <= saldo {
				saldo -= jumlah
				fmt.Println("Anda berhasil menarik uang sebesar Rp", jumlah)
				fmt.Println("Saldo tersisa: Rp", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
