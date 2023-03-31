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
<<<<<<< HEAD
	UpdateProfileCust(ctx *gin.Context)
	ShowCustByID(ctx *gin.Context)
	GetAllCust(ctx *gin.Context)
=======
	ShowCustByID(ctx *gin.Context)
	GetAllCust(ctx *gin.Context)
	UpdateCust(ctx *gin.Context)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
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
<<<<<<< HEAD
		response := utils.BuildErrorResponse("Failed to process login request", http.StatusBadRequest, utils.EmptyObj{})
=======
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	verify, _ := c.custSvc.VerifyCust(ctx.Request.Context(), custParam.Email, custParam.Password)
	if !verify {
		response := utils.BuildErrorResponse("Gagal, email/password salah", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.FindCustByEmail(ctx.Request.Context(), custParam.Email)
	if err != nil {
<<<<<<< HEAD
		response := utils.BuildErrorResponse("Gagal cari", http.StatusBadRequest, utils.EmptyObj{})
=======
		response := utils.BuildErrorResponse("Gagal cari email", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtSvc.GenerateToken(tx.ID, tx.Role)
	custResponse := dto.UserResponse{
		Token: token,
		Role:  tx.Role,
	}

	response := utils.BuildResponse("Login", http.StatusOK, custResponse)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) ShowCustByID(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process id request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.custSvc.FindCustByID(ctx.Request.Context(), custID)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari id", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Berhasil dapat", http.StatusOK, tx)
	ctx.JSON(http.StatusCreated, response)
}

<<<<<<< HEAD
// admin yg bisa
func (c *custController) GetAllCust(ctx *gin.Context) {
	cust, err := c.custSvc.FindCust(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Gagal dapatkan customer", http.StatusBadRequest, utils.EmptyObj{})
=======
func (c *custController) GetAllCust(ctx *gin.Context) {
	cust, err := c.custSvc.GetAllCust(ctx.Request.Context())
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari semua customer", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

<<<<<<< HEAD
	response := utils.BuildResponse("Berhasil dapatkan customer", http.StatusOK, cust)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) UpdateProfileCust(ctx *gin.Context) {
	var custParam dto.UserUpdate
	errParam := ctx.ShouldBindJSON(&custParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process update request", http.StatusBadRequest, utils.EmptyObj{})
=======
	response := utils.BuildResponse("berhasil cari", http.StatusOK, cust)
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) UpdateCust(ctx *gin.Context) {
	var custParam dto.UserUpdate
	errParam := ctx.ShouldBindJSON(&custParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := ctx.MustGet("token").(string)
<<<<<<< HEAD
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal dapatkan id", http.StatusBadRequest, utils.EmptyObj{})
=======
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process token request", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

<<<<<<< HEAD
	tx, err := c.custSvc.UpdateCust(ctx.Request.Context(), custParam, id)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal Update", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("profile updated", http.StatusCreated, tx)
=======
	tx, err := c.custSvc.UpdateCust(ctx.Request.Context(), custParam, custID)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to update", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Account Updated", http.StatusOK, tx)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	ctx.JSON(http.StatusCreated, response)
}

func (c *custController) DeleteCust(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
<<<<<<< HEAD
	id, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal dapatkan id", http.StatusBadRequest, utils.EmptyObj{})
=======
	custID, err := c.jwtSvc.GetUserIDByToken(token)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to process token request", http.StatusBadRequest, utils.EmptyObj{})
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

<<<<<<< HEAD
	tx, err := c.custSvc.DeleteCust(ctx.Request.Context(), id)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal menghapus", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse("profile deleted", http.StatusCreated, tx)
=======
	tx, err := c.custSvc.DeleteCust(ctx.Request.Context(), custID)
	if err != nil {
		response := utils.BuildErrorResponse("Gagal cari", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Account Deleted", http.StatusOK, tx)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	ctx.JSON(http.StatusCreated, response)
}
