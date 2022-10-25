package mocks

import (
	"errors"

	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
)

type MockRepository struct {
	MockData   []domain.Product
	ErrOnRead  string
	ErrOnWrite string
}

func (mr *MockRepository) GetLastID() string {
	return "4"
}

func (mr *MockRepository) GetAll() ([]domain.Product, error) {

	if mr.ErrOnRead != "" {
		return nil, errors.New(mr.ErrOnRead)
	}

	return mr.MockData, nil
}

func (mr *MockRepository) GetAllBySeller(sellerID string) ([]domain.Product, error) {

	var result []domain.Product

	if mr.ErrOnRead != "" {
		return result, errors.New(mr.ErrOnRead)
	}

	for _, product := range mr.MockData {
		if product.SellerID == sellerID {
			result = append(result, product)
		}
	}

	return result, nil
}

func (mr *MockRepository) Store(id string, seller string, desc string, price float64) (domain.Product, error) {

	var result domain.Product

	if mr.ErrOnWrite != "" {
		return result, errors.New(mr.ErrOnWrite)
	}

	result.ID = id
	result.SellerID = seller
	result.Description = desc
	result.Price = price

	mr.MockData = append(mr.MockData, result)

	return result, nil
}
