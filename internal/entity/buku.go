package entity

type Buku struct {
	ID      int
	Judul   string
	Penulis string
	Stok    int
}

func (b *Buku) Tersedia() bool {
	return b.Stok > 0
}
