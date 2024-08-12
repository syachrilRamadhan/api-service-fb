package product

import "time"

type Produk struct {
	ID        int
	Title     string
	Deskripsi string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
