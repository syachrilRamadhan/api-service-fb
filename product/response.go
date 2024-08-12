package product

type ProdukResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Deskripsi string `json:deskripsi"`
	Price     int    `json:"price"`
}
