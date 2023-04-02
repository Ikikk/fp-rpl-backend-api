package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
)

type userController struct {
	custSvc   services.CustSvc
	sellerSvc services.SellerSvc
}

type UserController interface {
	Register(ctx *gin.Context)
	GetSellerByName(ctx *gin.Context)
}

func NewUserController(cs services.CustSvc, ss services.SellerSvc) UserController {
	return &userController{
		custSvc:   cs,
		sellerSvc: ss,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var userParam dto.UserCreate
	errParam := ctx.ShouldBindJSON(&userParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process register request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if userParam.Role == "customer" {
		tx, err := c.custSvc.RegisterCust(ctx.Request.Context(), userParam)
		if err != nil {
			response := utils.BuildErrorResponse("Failed to Create New Account", http.StatusBadRequest, utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildResponse("New Customer Account Created", http.StatusCreated, tx)
		ctx.JSON(http.StatusCreated, response)

	} else if userParam.Role == "seller" {
		tx, err := c.sellerSvc.RegisterSeller(ctx.Request.Context(), userParam)
		if err != nil {
			response := utils.BuildErrorResponse("Failed to Create New Account", http.StatusBadRequest, utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		response := utils.BuildResponse("New Seller Account Created", http.StatusCreated, tx)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *userController) GetSellerByName(ctx *gin.Context) {
	firstname := ctx.Param("first_name")
	lastname := ctx.Param("last_name")
	// name, err := strings.
	if firstname == " " && lastname == " " {
		response := utils.BuildErrorResponse("Failed to process get request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	seller, err := c.sellerSvc.FindSellerByName(ctx.Request.Context(), firstname, lastname)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get seller's name", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get Seller Successful", http.StatusOK, seller)
	ctx.JSON(http.StatusCreated, response)
}

// func (c *userController) Logout(ctx *gin.Context) {
// 	tokenString, err := ctx.
// }
