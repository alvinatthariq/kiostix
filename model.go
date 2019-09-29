package main

// Book - asdasd
type Book struct {
	ID        string `form:"id" json:"id"`
	JudulBuku string `form:"judul_buku" json:"judul_buku"`
	Penulis   string `form:"penulis" json:"penulis"`
	Kategori  string `form:"kategori" json:"kategori"`
}

// Response - addas
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Book
}
