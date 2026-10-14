# Soal 1: Fungsi utama `for` dan tiga komponennya

**Fungsi utama `for`**

Struktur perulangan `for` dipakai untuk mengulang blok kode berkali-kali selama kondisi tertentu masih terpenuhi. Dengan `for`, kita tidak perlu menulis kode yang sama berulang-ulang.

**Bentuk umum**

```go
for inisialisasi; kondisi; post {
    // kode yang diulang
}
```

**Tiga komponen**

1. **Inisialisasi**: dijalankan satu kali di awal perulangan untuk menyiapkan variabel pengulang, contoh `i := 1`.
2. **Kondisi**: dicek sebelum setiap iterasi. Jika bernilai `true`, badan perulangan dijalankan. Jika `false`, perulangan berhenti, contoh `i <= 5`.
3. **Post statement**: dijalankan setelah setiap iterasi selesai, biasanya untuk mengubah variabel pengulang, contoh `i++`.

**Contoh**

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```