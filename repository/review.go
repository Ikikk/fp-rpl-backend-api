package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type reviewRepo struct {
	db *gorm.DB
}

type ReviewRepo interface {
	CreateReview(ctx context.Context, review entity.Review) (entity.Review, error)
}

func NewReviewRepo(db *gorm.DB) ReviewRepo {
	return &reviewRepo{
		db: db,
	}
}

func (repo *reviewRepo) CreateReview(ctx context.Context, review entity.Review) (entity.Review, error) {
	var err error
	tx := repo.db.Create(&review)
	if tx.Error != nil {
		log.Println(err)
		return entity.Review{}, err
	}
	return review, nil
}
