package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorhcamacho/defafio-gotesting-victormartinez/internal/products"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}

}

func buildProductsRoutes(r *gin.RouterGroup) {
	repo := products.NewRepository()
	service := products.NewService(repo)
	handler := products.NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.POST("/", handler.StoreNewProduct)
		prodRoute.GET("/", handler.GetProducts)
	}
}
