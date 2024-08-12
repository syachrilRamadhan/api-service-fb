package product

type Service interface {
	GetProduk() ([]Produk, error)
	GetProdukById(ID int) (Produk, error)
	CreateProduk(product ProdukRequest) (Produk, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProduk() ([]Produk, error) {
	return s.repository.GetProduk()
}

func (s *service) GetProdukById(ID int) (Produk, error) {
	return s.repository.GetProdukById(ID)
}

func (s *service) CreateProduk(product ProdukRequest) (Produk, error) {
	produk := Produk{
		Title:     product.Title,
		Price:     product.Price,
		Deskripsi: product.Deskripsi,
	}
	newProduk, err := s.repository.CreateProduk(produk)
	return newProduk, err
}
