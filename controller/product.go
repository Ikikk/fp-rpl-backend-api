package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productSvc services.ProductSvc
}

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetAllProduct(ctx *gin.Context)
}

func NewProductController(ps services.ProductSvc) ProductController {
	return &productController{
		productSvc: ps,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var productParam dto.Product
	errParam := ctx.ShouldBindJSON(&productParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.productSvc.CreateProduct(ctx.Request.Context(), productParam)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to Create New Product", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Product Created", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)

}

func (c *productController) GetAllProduct(ctx *gin.Context) {
	product, err := c.productSvc.GetAllProduct(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Gagal dapatkan produk", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Berhasil dapatkan produk", http.StatusOK, product)
	ctx.JSON(http.StatusCreated, response)
}
