ALGORITMA DigitGanda
KAMUS
    n, puluhan, satuan, hasil : integer

DESKRIPSI
    input(n)
    puluhan ← n div 10
    satuan  ← n mod 10
    hasil   ← puluhan * 1000 + puluhan * 100 + satuan * 10 + satuan
    output(hasil)