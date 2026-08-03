package product

type ProductCreateResponse struct {
	IdProduct int     `json:"id_product"`
	Name      string  `json:"name"`
	Stok      int     `json:"stok"`
	Price     float64 `json:"price"`
}