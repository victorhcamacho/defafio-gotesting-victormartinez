package products

import "github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"

type Repository interface {
	GetLastID() string
	GetAll() ([]domain.Product, error)
	GetAllBySeller(sellerID string) ([]domain.Product, error)
	Store(id string, seller string, desc string, price float64) (domain.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetLastID() string {
	return ""
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *repository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (r *repository) Store(id string, seller string, desc string, price float64) (domain.Product, error) {
	return domain.Product{ID: id, SellerID: seller, Description: desc, Price: price}, nil
}
