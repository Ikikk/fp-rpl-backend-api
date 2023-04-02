package controller

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/services"
	"FP-RPL-ECommerce/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reviewController struct {
	reviewSvc services.ReviewSvc
}

type ReviewController interface {
	CreateReview(ctx *gin.Context)
}

func NewReviewController(rs services.ReviewSvc) ReviewController {
	return &reviewController{
		reviewSvc: rs,
	}
}

func (c *reviewController) CreateReview(ctx *gin.Context) {
	var reviewParam dto.Review
	errParam := ctx.ShouldBindJSON(&reviewParam)
	if errParam != nil {
		response := utils.BuildErrorResponse("Failed to process create request", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	tx, err := c.reviewSvc.CreateReview(ctx.Request.Context(), reviewParam)
	if err != nil {
		response := utils.BuildErrorResponse("Failed to create new review", http.StatusBadRequest, utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := utils.BuildResponse("Review Created", http.StatusCreated, tx)
	ctx.JSON(http.StatusCreated, response)
}
