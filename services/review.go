package services

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"context"

	"github.com/jinzhu/copier"
)

type reviewSvc struct {
	reviewRepo repository.ReviewRepo
}

type ReviewSvc interface {
	CreateReview(ctx context.Context, reviewParam dto.Review) (entity.Review, error)
}

func NewReviewSvc(reviewRepo repository.ReviewRepo) ReviewSvc {
	return &reviewSvc{
		reviewRepo: reviewRepo,
	}
}

func (svc *reviewSvc) CreateReview(ctx context.Context, reviewParam dto.Review) (entity.Review, error) {
	var review entity.Review
	copier.Copy(&review, &reviewParam)

	createdReview, err := svc.reviewRepo.CreateReview(ctx, review)
	if err != nil {
		return entity.Review{}, err
	}
	return createdReview, nil
}
