package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"

	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sellerController struct {
	sellerSvc services.SellerSvc
	jwtSvc    services.JWTService
}

type SellerController interface {
	LoginCust(ctx *gin.Context)
	ShowSellerByID(ctx *gin.Context)
	GetAllSeller(ctx *gin.Context)
	// GetSellerByName(ctx *gin.Context)
	UpdateProfileSeller(ctx *gin.Context)
	DeleteSeller(ctx *gin.Context)
}

func NewSellerController(cs services.SellerSvc, jwt services.JWTService) SellerController {
	return &sellerController{
		sellerSvc: cs,
		jwtSvc:    jwt,
	}
}

func (c *sellerController) LoginCust(ctx *gin.Context) {
	var sellerParam dto.UserLogin
	errParam := ctx.ShouldBindJSON(&sellerParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process get seller's id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	verify, _ := c.sellerSvc.VerifySeller(ctx.Request.Context(), sellerParam.Email, sellerParam.Password)
	if !verify {
		response := utils.BuildErrorResponse("Failed to seller login, wrong email or password ", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.FindSellerByEmail(ctx.Request.Context(), sellerParam.Email)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get seller's email", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtSvc.GenerateToken(tx.ID, tx.Role)
	sellerResponse := dto.UserResponse{
		Token: token,
		Role:  tx.Role,
	}

	response := utils.BuildResponse("Login Successful", http.StatusCreated, sellerResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) ShowSellerByID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.FindSellerByID(ctx.Request.Context(), custID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get seller's id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get Seller By ID Successfull", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}

// admin yg bisa
func (c *sellerController) GetAllSeller(ctx *gin.Context) {
	seller, err := c.sellerSvc.FindSeller(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get all seller request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get All Seller Successful", http.StatusOK, seller)
	ctx.JSON(http.StatusCreated, response)
}

// func (c *sellerController) GetSellerByName(ctx *gin.Context) {
// 	var sellerParam dto.UserUpdate
// 	errParam := ctx.ShouldBindJSON(&sellerParam)
// 	if errParam != nil {
// 		response := utils.BuildErrorResponse("Failed to process get request", http.StatusBadRequest, utils.EmptyObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	seller, err := c.sellerSvc.FindSellerByName(ctx.Request.Context(), sellerParam.FirstName, sellerParam.LastName)
// 	if err != nil {
// 		response := utils.BuildErrorResponse("Gagal dapatkan seller", http.StatusBadRequest, utils.EmptyObj{})
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := utils.BuildResponse("Berhasil dapatkan seller", http.StatusOK, seller)
// 	ctx.JSON(http.StatusCreated, response)
// }

func (c *sellerController) UpdateProfileSeller(ctx *gin.Context) {
	var sellerParam dto.UserUpdate
	errParam := ctx.ShouldBindJSON(&sellerParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process update profile request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get seller's id by token", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.UpdateSeller(ctx.Request.Context(), sellerParam, id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Seller Profile Updated", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}

func (c *sellerController) DeleteSeller(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get id by token", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.sellerSvc.DeleteSeller(ctx.Request.Context(), id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete profile", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Seller Profile Deleted", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}
