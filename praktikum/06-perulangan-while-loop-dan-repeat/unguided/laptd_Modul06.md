# <h1 align="center">Tugas Pendahuluan Modul 6 - PERULANGAN WHILE LOOP DAN REPEAT</h1>
<p align="center">PERULANGAN WHILE LOOP DAN REPEAT - 109092600008</p>

### 1. Menentukan Kelulusan

```go
package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("Program selesai.")
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}
}
```

##### Contoh Input

```text
75
45
-1
```

##### Output

```text
Masukkan nilai Anda: Anda lulus ujian.
Masukkan nilai Anda: Anda tidak lulus ujian.
Masukkan nilai Anda: Program selesai.
```

#### Deskripsi
Program terus meminta nilai menggunakan perulangan tanpa kondisi awal (`for {}`). Nilai `-1` berfungsi sebagai penanda untuk menghentikan perulangan dengan `break`. Selain nilai tersebut, program menentukan kelulusan: nilai minimal 60 dinyatakan lulus.

### 2. Menghitung Jumlah Deret

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	i := 1
	for i <= n {
		sum += i
		i++
	}

	fmt.Println(sum)
}
```

##### Contoh Input

```text
5
```

##### Output

```text
15
```

#### Deskripsi
Program menghitung jumlah bilangan bulat dari 1 sampai `n`. Variabel `i` bertambah satu pada setiap iterasi, sementara nilainya ditambahkan ke `sum`. Ketika `i` sudah lebih besar dari `n`, perulangan berhenti dan hasilnya dicetak. Untuk `n = 5`, hasilnya adalah `1 + 2 + 3 + 4 + 5 = 15`.

### 3. Menghitung Biaya Setelah Kenaikan Tahunan

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0
	tahun := 1
	for tahun < n {
		biaya *= 1.06
		tahun++
	}

	fmt.Printf("%.0f\n", biaya)
}
```

##### Contoh Input

```text
3
```

##### Output

```text
11236000
```

#### Deskripsi
Biaya awal sebesar Rp10.000.000 meningkat 6% setiap tahun. Perulangan berjalan selama `tahun < n`; untuk input `3`, kenaikan diterapkan dua kali sehingga hasilnya Rp11.236.000. Format `%.0f` menampilkan hasil tanpa angka di belakang koma.

### 4. Menentukan Dua Nilai Terendah

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	nama1, nama2 := "", ""
	min1, min2 := math.MaxInt, math.MaxInt

	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		if nilai < min1 {
			min2, nama2 = min1, nama1
			min1, nama1 = nilai, nama
		} else if nilai < min2 {
			min2, nama2 = nilai, nama
		}
		i++
	}

	fmt.Println(nama1)
	fmt.Println(nama2)
}
```

##### Contoh Input

```text
4
Ani 80
Budi 65
Citra 90
Dedi 70
```

##### Output

```text
Budi
Dedi
```

#### Deskripsi
Program membaca jumlah peserta, lalu nama dan nilai setiap peserta. Variabel `min1` dan `min2` menyimpan nilai terendah dan kedua terendah yang ditemukan, sedangkan `nama1` dan `nama2` menyimpan nama pemiliknya. Setelah semua data dibaca, program mencetak nama dengan nilai terendah terlebih dahulu, lalu nama dengan nilai kedua terendah.

## Kesimpulan
Pada Modul 6, perulangan `for` digunakan untuk mengulang instruksi selama suatu kondisi terpenuhi maupun terus-menerus hingga dihentikan dengan `break`. Pola perulangan ini dapat dipakai untuk memproses masukan berulang, menghitung deret dan kenaikan bertahap, serta mencari nilai terendah dari sejumlah data.
