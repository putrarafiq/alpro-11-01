# Jawaban Soal 1 - Percabangan di Go

## Perbedaan `if`, `if-else`, `else if`, dan `switch`

| Struktur | Fungsi | Kapan dipakai |
|----------|--------|---------------|
| `if` | Menjalankan blok kode **hanya jika** kondisi bernilai `true`. Jika `false`, blok dilewati begitu saja. | Ada satu kondisi, dan tidak ada aksi alternatif. |
| `if-else` | Menyediakan **dua jalur**: blok `if` dijalankan jika kondisi `true`, selain itu blok `else` yang dijalankan. | Ada dua kemungkinan yang saling berlawanan (benar/salah). |
| `else if` | Dipakai setelah `if` untuk memeriksa **kondisi tambahan** secara berurutan. Pemeriksaan berhenti pada kondisi pertama yang `true`; `else` di akhir menangani sisanya. | Ada banyak kondisi (rentang nilai, perbandingan, dsb). |
| `switch` | Membandingkan **satu nilai/ekspresi** dengan beberapa `case`. Blok `case` yang cocok dijalankan, `default` dijalankan jika tidak ada yang cocok. | Banyak pilihan berdasarkan satu nilai tertentu, sehingga lebih rapi daripada `else if` panjang. |

### Contoh

```go
// if
if nilai >= 60 {
    fmt.Println("Lulus")
}

// if-else
if nilai >= 60 {
    fmt.Println("Lulus")
} else {
    fmt.Println("Tidak lulus")
}

// else if
if nilai >= 80 {
    fmt.Println("A")
} else if nilai >= 70 {
    fmt.Println("B")
} else if nilai >= 60 {
    fmt.Println("C")
} else {
    fmt.Println("D")
}

// switch
switch hari {
case 1:
    fmt.Println("Senin")
case 2:
    fmt.Println("Selasa")
default:
    fmt.Println("Hari lain")
}
```

## Dua perbedaan `switch` Go dibandingkan `switch` Java

### 1. Tidak ada *fall-through* otomatis

- **Go:** setiap `case` otomatis berhenti setelah blok selesai, jadi **tidak perlu** menulis `break`. Jika ingin lanjut ke `case` berikutnya, harus ditulis eksplisit dengan kata kunci `fallthrough`.
- **Java:** tanpa `break`, eksekusi akan **terus menjalankan** `case` di bawahnya (fall-through), sehingga `break` wajib ditulis di setiap `case`.

```go
// Go: tanpa break pun tidak akan "bocor" ke case berikutnya
switch x {
case 1:
    fmt.Println("satu")
case 2:
    fmt.Println("dua")
}
```

```java
// Java: tanpa break, "satu" dan "dua" sama-sama tercetak jika x = 1
switch (x) {
    case 1:
        System.out.println("satu");
    case 2:
        System.out.println("dua");
}
```

### 2. `case` di Go lebih fleksibel (ekspresi, banyak nilai, dan tanpa tag)

- **Go:** satu `case` boleh berisi **beberapa nilai** dipisah koma, dan `switch` boleh ditulis **tanpa ekspresi** sehingga setiap `case` berisi kondisi boolean (menggantikan `else if` panjang).
- **Java:** `case` pada `switch` klasik hanya berisi **satu nilai konstan** (int, char, String, enum), dan tidak bisa langsung berisi kondisi seperti `x > 10`.

```go
// beberapa nilai dalam satu case
switch hari {
case 6, 7:
    fmt.Println("Akhir pekan")
}

// switch tanpa ekspresi: case berisi kondisi
switch {
case nilai >= 80:
    fmt.Println("A")
case nilai >= 70:
    fmt.Println("B")
default:
    fmt.Println("C atau di bawahnya")
}
```