    # Soal 2 - Tracing: Evaluasi Pernyataan Kondisi

## Nilai Awal

| Variabel | Nilai |
|----------|-------|
| x        | 10    |
| y        | 5     |
| z        | 15    |
| result   | 0     |

## Jawaban

### 1. Nilai akhir `result`

**25**

### 2. Output program

```
Nilai akhir result: 25
```

### 3. Langkah-langkah alur eksekusi

**Kondisi 1: `if x > 5`**

- `x > 5` → `10 > 5` → **benar**.
- Masuk ke blok `if`, lalu dicek `if y < 10` → `5 < 10` → **benar**.
- Jalankan `result = x + y` → `result = 10 + 5 = 15`.
- Blok `else` dilewati.

**Kondisi 2: `if z > 10 && x == 10`**

- `z > 10` → `15 > 10` → benar.
- `x == 10` → `10 == 10` → benar.
- Karena kedua sisi `&&` benar, kondisi ini **benar**.
- Jalankan `result += z` → `result = 15 + 15 = 30`.
- Blok `else` dilewati.
- Pengaruh ke `result`: bertambah 15, dari 15 menjadi 30.

**Kondisi 3: `if x == 10 || y > 10`**

- `x == 10` → `10 == 10` → benar.
- Karena operator `||` hanya butuh salah satu sisi benar, kondisi ini **benar** (sisi kanan `y > 10` yang bernilai salah tidak menjadi masalah, bahkan tidak perlu dievaluasi karena short-circuit).
- Jalankan `result += 5` → `result = 30 + 5 = 35`.
- Blok `else if` dan `else` dilewati.

**Kondisi 4: `if !(x < 15 && y < 10)`**

- `x < 15` → `10 < 15` → benar.
- `y < 10` → `5 < 10` → benar.
- `x < 15 && y < 10` → benar, lalu di-NOT menjadi **salah**.
- Karena kondisi `if` salah, masuk ke blok `else`.
- Jalankan `result -= 10` → `result = 35 - 10 = 25`.

**Cetak hasil**

- `fmt.Println("Nilai akhir result:", result)` mencetak `Nilai akhir result: 25`.

## Tabel Ringkasan Perubahan `result`

| Tahap | Kondisi | Hasil Kondisi | Operasi | `result` |
|-------|---------|---------------|---------|----------|
| Awal  | -       | -             | -       | 0        |
| 1     | `x > 5` lalu `y < 10` | benar, benar | `result = x + y` | 15 |
| 2     | `z > 10 && x == 10` | benar | `result += z` | 30 |
| 3     | `x == 10 \|\| y > 10` | benar | `result += 5` | 35 |
| 4     | `!(x < 15 && y < 10)` | salah (masuk else) | `result -= 10` | 25 |