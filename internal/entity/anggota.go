package entity

type Anggota struct {
	ID      int
	BukuId  int
	Nama    string
	Nim     string
	Status  string
	AllBuku []Buku
}

func (a *Anggota) IsAktif() bool {
	return a.Status == "aktif"
}
