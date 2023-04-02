package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productSvc services.ProductSvc
	jwtSvc     services.JWTService
}

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetAllProduct(ctx *gin.Context)
	GetProductByID(ctx *gin.Context)
	GetProductByName(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
}

func NewProductController(ps services.ProductSvc, js services.JWTService) ProductController {
	return &productController{
		productSvc: ps,
		jwtSvc:     js,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	var productParam dto.Product
	errParam := ctx.ShouldBindJSON(&productParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process create request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.productSvc.CreateProduct(ctx.Request.Context(), productParam)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to create new product", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Product Created", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)

}

func (c *productController) GetAllProduct(ctx *gin.Context) {
	product, err := c.productSvc.GetAllProduct(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get all product", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get All Product Successfull", http.StatusOK, product)
	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) GetProductByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	product_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	product, err := c.productSvc.GetProductByID(ctx.Request.Context(), product_id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get product's id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get Product By ID Successfull", http.StatusOK, product)
	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) GetProductByName(ctx *gin.Context) {
	name := ctx.Params.ByName("product_name")
	if name == " " {
		response := utils.BuildErrorResponse("Failed to process get name request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	product, err := c.productSvc.GetProductByName(ctx.Request.Context(), name)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get product's name", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get Product By Name Successfull", http.StatusOK, product)
	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	product_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var productParam dto.Product
	errParam := ctx.ShouldBindJSON(&productParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process update product request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.productSvc.UpdateProduct(ctx.Request.Context(), productParam, product_id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update product", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Product Updated", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}

func (c *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	product_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.productSvc.DeleteProduct(ctx.Request.Context(), product_id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete product", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Product Deleted", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}
