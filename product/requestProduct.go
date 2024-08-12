package product

type ProdukRequest struct {
	Title     string `json:"title" binding:"required"`
	Price     int    `json:"price" binding:"required,number"`
	Deskripsi string `json:"deskripsi" binding:"required"`
}
