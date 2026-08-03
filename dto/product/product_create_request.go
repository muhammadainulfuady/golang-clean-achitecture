package product

type ProductCreateRequest struct {
	Name  string  `json:"name" validate:"required, max=255, min=5"`
	Stok  int     `json:"stok" validate:"required, min=1"`
	Price float64 `json:"price" validate:"required, min=1"`
}
