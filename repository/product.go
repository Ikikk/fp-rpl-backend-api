package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type productRepo struct {
	db         *gorm.DB
	sellerRepo SellerRepo
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetAllProduct(ctx context.Context) (product []entity.Product, err error)
}

func NewProductRepo(db *gorm.DB, sr SellerRepo) ProductRepo {
	return &productRepo{
		db:         db,
		sellerRepo: sr,
	}
}

func (repo *productRepo) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	var err error
	tx := repo.db.Create(&product)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	return product, nil
}

func (repo *productRepo) GetAllProduct(ctx context.Context) (product []entity.Product, err error) {
	tx := repo.db.Preload("User").Preload("Like").Preload("Review").Find(&product)
	if tx.Error != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}
