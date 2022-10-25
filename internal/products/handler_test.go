package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/test/mocks"
)

func createServer(mockService *mocks.MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewHandler(mockService)

	server := gin.Default()

	// version := server.Group("/api/v1")
	routes := server.Group("/products")

	routes.POST("/", handler.StoreNewProduct)
	routes.GET("/", handler.GetProducts)

	return server
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	httpRequest := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	httpRequest.Header.Add("Content-Type", "application/json")

	return httpRequest, httptest.NewRecorder()
}

func TestHandlerGetAllProducts(t *testing.T) {

	// successful
	testMock := mocks.MockService{
		MockData: testData,
	}

	server := createServer(&testMock)
	req, rec := createRequestTest(http.MethodGet, "/products/", "")

	server.ServeHTTP(rec, req)

	var testResult []domain.Product
	err := json.Unmarshal(rec.Body.Bytes(), &testResult)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, testMock.MockData, testResult)

	// failure
	testMock.ErrOnRead = "not found data"

	expectedErr := map[string]string{"error": "not found data"}

	server = createServer(&testMock)
	req, rec = createRequestTest(http.MethodGet, "/products/", "")

	server.ServeHTTP(rec, req)

	result := map[string]string{}
	err = json.Unmarshal(rec.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.Equal(t, 500, rec.Code)
	assert.Contains(t, result, "error")
	assert.Equal(t, expectedErr, result)
}

func TestHandlerGetAllProductsBySeller(t *testing.T) {

	// successful
	testMock := mocks.MockService{
		MockData: testData,
	}

	server := createServer(&testMock)
	req, rec := createRequestTest(http.MethodGet, "/products/?seller_id=FEX112AC", "")

	server.ServeHTTP(rec, req)

	var testResult []domain.Product
	err := json.Unmarshal(rec.Body.Bytes(), &testResult)

	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, testMock.MockData[1:], testResult)

	// failure
	testMock.ErrOnRead = "not found products"

	expectedErr := map[string]string{"error": "not found products"}

	server = createServer(&testMock)
	req, rec = createRequestTest(http.MethodGet, "/products/?seller_id=FEX112AC", "")

	server.ServeHTTP(rec, req)

	result := map[string]string{}
	err = json.Unmarshal(rec.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.Equal(t, 500, rec.Code)
	assert.Contains(t, result, "error")
	assert.Equal(t, expectedErr, result)
}

func TestHandlerStoreProduct(t *testing.T) {

	// successful
	expectedResult := domain.Product{
		ID:          "4",
		SellerID:    "FEX112AB",
		Description: "generic product",
		Price:       789.5,
	}

	testMock := mocks.MockService{
		MockData: testData,
	}

	server := createServer(&testMock)
	req, rec := createRequestTest(http.MethodPost, "/products/", `{"SellerID":"FEX112AB","Description":"generic product","Price":789.5}`)

	server.ServeHTTP(rec, req)

	var testResult domain.Product
	err := json.Unmarshal(rec.Body.Bytes(), &testResult)

	assert.Nil(t, err)
	assert.Equal(t, 201, rec.Code)
	assert.Equal(t, expectedResult, testResult)
	assert.Equal(t, len(testMock.MockData), 4)

	// failure
	server = createServer(&testMock)
	req, rec = createRequestTest(http.MethodPost, "/products/", "")

	server.ServeHTTP(rec, req)

	result := map[string]string{}
	err = json.Unmarshal(rec.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.Equal(t, 400, rec.Code)
	assert.Contains(t, result, "error")

	server = createServer(&testMock)
	req, rec = createRequestTest(http.MethodPost, "/products/", `{"Description":"generic product","Price":789.5}`)

	server.ServeHTTP(rec, req)

	result = map[string]string{}
	err = json.Unmarshal(rec.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.Equal(t, 400, rec.Code)
	assert.Contains(t, result, "error")

	// failure at mock
	testMock.ErrOnWrite = "internal server error"

	expectedErr := map[string]string{"error": "internal server error"}

	server = createServer(&testMock)
	req, rec = createRequestTest(http.MethodPost, "/products/", `{"SellerID":"FEX112AB","Description":"generic product","Price":789.5}`)

	server.ServeHTTP(rec, req)

	result = map[string]string{}
	err = json.Unmarshal(rec.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.Equal(t, 409, rec.Code)
	assert.Contains(t, result, "error")
	assert.Equal(t, expectedErr, result)
}
