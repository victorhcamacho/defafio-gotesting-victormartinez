package products

import (
	"log"

	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetAllBySeller(sellerID string) ([]domain.Product, error)
	Store(sellerID string, desc string, price float64) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]domain.Product, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		log.Println("error in repository ", err.Error())
		return nil, err
	}
	return data, err
}

func (s *service) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)
	if err != nil {
		log.Println("error in repository ", err.Error(), " sellerId: ", sellerID)
		return nil, err
	}
	return data, nil
}

func (s *service) Store(sellerID string, desc string, price float64) (domain.Product, error) {
	id := s.repo.GetLastID()
	data, err := s.repo.Store(id, sellerID, desc, price)
	if err != nil {
		log.Println("error in repository ", err.Error())
		return domain.Product{}, err
	}
	return data, nil
}
