package products

import (
	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/domain"
)

type requestDTO struct {
	SellerID    string
	Description string
	Price       float64
}

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		svc: s,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {

	var err error
	var products []domain.Product

	sellerID := ctx.Query("seller_id")

	if sellerID != "" {
		products, err = h.svc.GetAllBySeller(sellerID)
	} else {
		products, err = h.svc.GetAll()
	}

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, products)
}

func (h *Handler) StoreNewProduct(ctx *gin.Context) {

	var request requestDTO

	if errParse := ctx.ShouldBindJSON(&request); errParse != nil {
		ctx.JSON(400, gin.H{"error": errParse.Error()})
		return
	}

	if request.SellerID == "" {
		ctx.JSON(400, gin.H{"error": "seller id param is required"})
		return
	}

	if request.Description == "" {
		ctx.JSON(400, gin.H{"error": "description param is required"})
		return
	}

	if request.Price == 0 {
		ctx.JSON(400, gin.H{"error": "price param is required"})
		return
	}

	result, err := h.svc.Store(request.SellerID, request.Description, request.Price)

	if err != nil {
		ctx.JSON(409, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, result)
}
