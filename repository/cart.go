package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

type CartRepo interface {
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
}

func (repo *cartRepo) NewCartRepo(db *gorm.DB) CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (repo *cartRepo) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	var err error
	tx := repo.db.Create(&cart)
	if tx.Error != nil {
		log.Println(err)
		return entity.Cart{}, err
	}

	return cart, nil
}
