package product

type Product struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}
