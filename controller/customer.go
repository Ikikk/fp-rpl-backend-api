package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type custController struct {
	custSvc services.CustSvc
	jwtSvc  services.JWTService
}

type CustController interface {
	LoginCust(ctx *gin.Context)
	UpdateProfileCust(ctx *gin.Context)
	ShowCustByID(ctx *gin.Context)
	GetAllCust(ctx *gin.Context)
	DeleteCust(ctx *gin.Context)
}

func NewCustController(cs services.CustSvc, jwt services.JWTService) CustController {
	return &custController{
		custSvc: cs,
		jwtSvc:  jwt,
	}
}

func (c *custController) LoginCust(ctx *gin.Context) {
	var custParam dto.UserLogin
	errParam := ctx.ShouldBindJSON(&custParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process customer login request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	verify, _ := c.custSvc.VerifyCust(ctx.Request.Context(), custParam.Email, custParam.Password)
	if !verify {
		response := utils.BuildErrorResponse("Failed to customer login, wrong email or password ", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.FindCustByEmail(ctx.Request.Context(), custParam.Email)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get customer's email", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtSvc.GenerateToken(tx.ID, tx.Role)
	custResponse := dto.UserResponse{
		Token: token,
		Role:  tx.Role,
	}

	response := utils.BuildResponse("Login Successfull", http.StatusOK, custResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) ShowCustByID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process get id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.FindCustByID(ctx.Request.Context(), custID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get customer's id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get Customer Successful", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}

// admin yg bisa
func (c *custController) GetAllCust(ctx *gin.Context) {
	cust, err := c.custSvc.FindCust(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get all customer", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Get All Customer Successful", http.StatusOK, cust)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) UpdateProfileCust(ctx *gin.Context) {
	var custParam dto.UserUpdate
	errParam := ctx.ShouldBindJSON(&custParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process update customer profile request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get customer's id by token", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.UpdateCust(ctx.Request.Context(), custParam, id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update customer profile", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Customer Profile Updated", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) DeleteCust(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to get customer by id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.DeleteCust(ctx.Request.Context(), id)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to delete customer profile", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("Customer Profile Deleted", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}
