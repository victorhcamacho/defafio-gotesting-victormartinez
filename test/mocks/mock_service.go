package mocks

import (
	"errors"

	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
)

type MockService struct {
	MockData   []domain.Product
	ErrOnRead  string
	ErrOnWrite string
}

func (ms *MockService) GetAll() ([]domain.Product, error) {
	if ms.ErrOnRead != "" {
		return nil, errors.New(ms.ErrOnRead)
	}
	return ms.MockData, nil
}

func (ms *MockService) GetAllBySeller(sellerID string) ([]domain.Product, error) {

	var result []domain.Product

	if ms.ErrOnRead != "" {
		return result, errors.New(ms.ErrOnRead)
	}

	for _, product := range ms.MockData {
		if product.SellerID == sellerID {
			result = append(result, product)
		}
	}

	return result, nil
}

func (ms *MockService) Store(seller string, desc string, price float64) (domain.Product, error) {

	var result domain.Product

	if ms.ErrOnWrite != "" {
		return result, errors.New(ms.ErrOnWrite)
	}

	result.ID = "4"
	result.SellerID = seller
	result.Description = desc
	result.Price = price

	ms.MockData = append(ms.MockData, result)

	return result, nil
}
