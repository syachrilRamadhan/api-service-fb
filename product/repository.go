package product

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetProduk() ([]Produk, error)
	GetProdukById(ID int) (Produk, error)
	CreateProduk(product Produk) (Produk, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProduk() ([]Produk, error) {
	var produks []Produk

	err := r.db.Find(&produks).Error

	return produks, err
}

func (r *repository) GetProdukById(ID int) (Produk, error) {
	var produk Produk

	err := r.db.Find(&produk, ID).Error

	return produk, err
}

func (r *repository) CreateProduk(product Produk) (Produk, error) {
	err := r.db.Create(&product).Error
	return product, err
}
