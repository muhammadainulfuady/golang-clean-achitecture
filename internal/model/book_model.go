package model

// 1. DTO untuk Request Tambah Buku (Client -- Server)
type CreateBookRequest struct {
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publication_year" validate:"required"`
	Description     string `json:"description"`
}

// 2. DTO untuk Request Update Buku
type UpdateBookRequest struct {
	ID              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publication_year" validate:"required"`
	Description     string `json:"description"`
}

// 3. DTO untuk Response Buku (Server -- Client)
type BookResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
}
