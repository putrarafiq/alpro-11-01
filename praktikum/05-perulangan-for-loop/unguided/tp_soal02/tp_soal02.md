# Soal 2: Perbedaan `for` tiga komponen dan `for` hanya kondisi

Go tidak memiliki perulangan `while` dan `do-while`. Semua kebutuhan perulangan ditangani oleh `for`.

## a. `for` dengan tiga komponen

Dipakai ketika jumlah pengulangan sudah jelas. Variabel pengulang dibuat, dicek, dan diubah langsung di baris `for`.

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

## b. `for` hanya dengan kondisi

Berfungsi seperti `while` di bahasa lain. Tidak ada inisialisasi dan post statement di baris `for`, sehingga variabel dibuat sebelum perulangan dan diubah sendiri di dalam badan perulangan. Cocok ketika jumlah pengulangan tidak pasti.

```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

## Perbedaan utama

| Aspek | `for` tiga komponen | `for` hanya kondisi |
|---|---|---|
| Inisialisasi | Di baris `for` | Sebelum `for` |
| Kondisi | Di baris `for` | Di baris `for` |
| Post statement | Di baris `for` | Di dalam badan perulangan |
| Cocok untuk | Jumlah pengulangan pasti | Jumlah pengulangan tidak pasti |