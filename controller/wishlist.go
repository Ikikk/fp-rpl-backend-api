package controller

import (
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type wishlistController struct {
	wishlistSvc services.WishlistSvc
	jwtSvc      services.JWTService
}

type WishlistController interface {
	AddWishlist(ctx *gin.Context)
}

func NewWishlistController(ws services.WishlistSvc, js services.JWTService) WishlistController {
	return &wishlistController{
		wishlistSvc: ws,
		jwtSvc:      js,
	}
}

func (c *wishlistController) AddWishlist(ctx *gin.Context) {

	token := ctx.MustGet("token").(string)
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get user_id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	product_id := ctx.Params.ByName("product_id")
	productID, err := strconv.ParseUint(product_id, 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get product_id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.wishlistSvc.AddWishlist(ctx, custID, productID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to add wishlist / product undiscovered", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Add Wishlist Successfuly", http.StatusCreated, result)
	ctx.JSON(http.StatusCreated, response)
}
