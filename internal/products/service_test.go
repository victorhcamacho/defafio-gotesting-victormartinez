package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/test/mocks"
)

var testData = []domain.Product{
	{
		ID:          "1",
		SellerID:    "FEX112AB",
		Description: "generic keyboard",
		Price:       99.9,
	},
	{
		ID:          "2",
		SellerID:    "FEX112AC",
		Description: "generic mouse",
		Price:       149.5,
	},
	{
		ID:          "3",
		SellerID:    "FEX112AC",
		Description: "generic screen",
		Price:       15000,
	},
}

func TestServiceGetAllProducts(t *testing.T) {

	//successful
	testMock := mocks.MockRepository{
		MockData: testData,
	}

	service := NewService(&testMock)

	testResult, testErr := service.GetAll()

	assert.Nil(t, testErr)
	assert.Equal(t, testData, testResult)

	//failure
	testMock.ErrOnRead = "not found data"

	service = NewService(&testMock)

	testResult, testErr = service.GetAll()

	assert.NotNil(t, testErr)
	assert.Empty(t, testResult)
	assert.EqualError(t, errors.New("not found data"), testErr.Error())
}

func TestServiceGetAllProductsBySeller(t *testing.T) {

	testMock := mocks.MockRepository{
		MockData: testData,
	}

	service := NewService(&testMock)

	testResult, testErr := service.GetAllBySeller("FEX112AC")

	assert.Nil(t, testErr)
	assert.Equal(t, testData[1:], testResult)

	//failure
	testMock.ErrOnRead = "not found product"

	service = NewService(&testMock)

	testResult, testErr = service.GetAllBySeller("FEX112AC")

	assert.NotNil(t, testErr)
	assert.Empty(t, testResult)
	assert.EqualError(t, errors.New("not found product"), testErr.Error())
}

func TestServiceStoreProduct(t *testing.T) {

	expectedResult := domain.Product{
		ID:          "4",
		SellerID:    "FEX112AB",
		Description: "generic product",
		Price:       789.5,
	}

	testMock := mocks.MockRepository{
		MockData: testData,
	}

	service := NewService(&testMock)

	testResult, testErr := service.Store("FEX112AB", "generic product", 789.5)

	assert.Nil(t, testErr)
	assert.Equal(t, expectedResult, testResult)
	assert.Equal(t, len(testMock.MockData), 4)

	//failure
	testMock.ErrOnWrite = "internal server error"

	service = NewService(&testMock)

	testResult, testErr = service.Store("FEX112AB", "generic product", 789.5)

	assert.NotNil(t, testErr)
	assert.Empty(t, testResult)
	assert.EqualError(t, errors.New("internal server error"), testErr.Error())
}
